package baseapp

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const OptimisticProcessingTimeoutInSeconds = 30

// InitChain implements the ABCI interface. It runs the initialization logic
// directly on the CommitMultiStore.
func (app *BaseApp) InitChain(req abci.RequestInitChain) (res abci.ResponseInitChain) {
	// On a new chain, we consider the init chain block height as 0, even though
	// req.InitialHeight is 1 by default.
	initHeader := tmproto.Header{ChainID: req.ChainId, Time: req.Time}

	// If req.InitialHeight is > 1, then we set the initial version in the
	// stores.
	if req.InitialHeight > 1 {
		app.initialHeight = req.InitialHeight
		initHeader = tmproto.Header{ChainID: req.ChainId, Height: req.InitialHeight, Time: req.Time}
		err := app.cms.SetInitialVersion(req.InitialHeight)
		if err != nil {
			panic(err)
		}
	}

	// initialize the deliver state and check state with a correct header
	app.setDeliverState(initHeader)
	app.setCheckState(initHeader)
	app.setProcessProposalState(initHeader)

	// Store the consensus params in the BaseApp's paramstore. Note, this must be
	// done after the deliver state and context have been set as it's persisted
	// to state.
	if req.ConsensusParams != nil {
		app.StoreConsensusParams(app.deliverState.ctx, req.ConsensusParams)
		app.StoreConsensusParams(app.processProposalState.ctx, req.ConsensusParams)
	}

	if app.initChainer == nil {
		return
	}

	// add block gas meter for any genesis transactions (allow infinite gas)
	app.deliverState.ctx = app.deliverState.ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())
	app.processProposalState.ctx = app.processProposalState.ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())

	app.initChainer(app.deliverState.ctx, req)
	res = app.initChainer(app.processProposalState.ctx, req)
	// sanity check
	if len(req.Validators) > 0 {
		if len(req.Validators) != len(res.Validators) {
			panic(
				fmt.Errorf(
					"len(RequestInitChain.Validators) != len(GenesisValidators) (%d != %d)",
					len(req.Validators), len(res.Validators),
				),
			)
		}

		sort.Sort(abci.ValidatorUpdates(req.Validators))
		sort.Sort(abci.ValidatorUpdates(res.Validators))

		for i := range res.Validators {
			if !proto.Equal(&res.Validators[i], &req.Validators[i]) {
				panic(fmt.Errorf("genesisValidators[%d] != req.Validators[%d] ", i, i))
			}
		}
	}

	// In the case of a new chain, AppHash will be the hash of an empty string.
	// During an upgrade, it'll be the hash of the last committed block.
	var appHash []byte
	if !app.LastCommitID().IsZero() {
		appHash = app.LastCommitID().Hash
	} else {
		// $ echo -n '' | sha256sum
		// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
		emptyHash := sha256.Sum256([]byte{})
		appHash = emptyHash[:]
	}

	// NOTE: We don't commit, but BeginBlock for block `initial_height` starts from this
	// deliverState.
	return abci.ResponseInitChain{
		ConsensusParams: res.ConsensusParams,
		Validators:      res.Validators,
		AppHash:         appHash,
	}
}

// Info implements the ABCI interface.
func (app *BaseApp) Info(req abci.RequestInfo) abci.ResponseInfo {
	lastCommitID := app.cms.LastCommitID()

	return abci.ResponseInfo{
		Data:             app.name,
		Version:          app.version,
		AppVersion:       app.appVersion,
		LastBlockHeight:  lastCommitID.Version,
		LastBlockAppHash: lastCommitID.Hash,
	}
}

// SetOption implements the ABCI interface.
func (app *BaseApp) SetOption(req abci.RequestSetOption) (res abci.ResponseSetOption) {
	// TODO: Implement!
	return
}

// FilterPeerByAddrPort filters peers by address/port.
func (app *BaseApp) FilterPeerByAddrPort(ctx sdk.Context, info string) abci.ResponseQuery {
	if app.addrPeerFilter != nil {
		return app.addrPeerFilter(ctx, info)
	}

	return abci.ResponseQuery{}
}

// FilterPeerByID filters peers by node ID.
func (app *BaseApp) FilterPeerByID(ctx sdk.Context, info string) abci.ResponseQuery {
	if app.idPeerFilter != nil {
		return app.idPeerFilter(ctx, info)
	}

	return abci.ResponseQuery{}
}

func (app *BaseApp) ProcessProposal(req abci.RequestProcessProposal) (res abci.ResponseProcessProposal) {
	app.Logger().Info("cosmos.app.ProcessProposal")
	if app.optimisticProcessingInfo != nil {
		app.deliverState = nil
		app.processProposalState = nil
		app.stateToCommit = nil
		app.optimisticProcessingInfo = nil
	}
	app.Logger().Info("cosmos.app.ProcessProposal.create OptimisticProcessingInfo")
	optimisticProcessingInfo := &OptimisticProcessingInfo{
		Height:                     req.Height,
		Hash:                       req.Hash,
		BeginBlockResultCompletion: make(chan struct{}, 1),
		Completion:                 make(chan struct{}, 1),
		EndBlockResultCompletion:   make(chan struct{}, 1),
	}
	app.optimisticProcessingInfo = optimisticProcessingInfo

	if app.cms.TracingEnabled() {
		app.cms.SetTracingContext(map[string]interface{}{"blockHeight": req.GetHeight()})
	}

	header := tmproto.Header{
		ChainID:         req.GetChainId(),
		Height:          req.GetHeight(),
		Time:            req.GetTime(),
		ProposerAddress: req.ProposerAddress,
	}
	if app.processProposalState == nil {
		app.setProcessProposalState(header)
	} else {
		app.processProposalState.ctx = app.processProposalState.ctx.
			WithBlockHeader(header).
			WithBlockHeight(req.GetHeight())
	}

	// add block gas meter
	var gasMeter sdk.GasMeter
	if maxGas := app.getMaximumBlockGas(app.processProposalState.ctx); maxGas > 0 {
		gasMeter = sdk.NewGasMeter(maxGas)
	} else {
		gasMeter = sdk.NewInfiniteGasMeter()
	}

	app.processProposalState.ctx = app.processProposalState.ctx.
		WithBlockGasMeter(gasMeter).
		WithHeaderHash(req.Hash).
		WithConsensusParams(app.GetConsensusParams(app.processProposalState.ctx))

	// we also set block gas meter to checkState in case the application needs to
	// verify gas consumption during (Re)CheckTx
	if app.checkState != nil {
		app.checkState.ctx = app.checkState.ctx.
			WithBlockGasMeter(gasMeter).
			WithHeaderHash(req.Hash)
	}

	go app.doProcessProposal(req, header)
	return abci.ResponseProcessProposal{
		Status: abci.ResponseProcessProposal_ACCEPT,
	}
}

func (app *BaseApp) doProcessProposal(req abci.RequestProcessProposal, header tmproto.Header) {
	app.Logger().Info("cosmos.app.ProcessProposal.create doProcessProposal")
	if app.optimisticProcessingInfo != nil {
		optimisticBeginBlockResult := app.optimisticBeginBlock(
			abci.RequestBeginBlock{
				Hash:   req.Hash,
				Header: header,
				LastCommitInfo: abci.LastCommitInfo{
					Round: req.ProposedLastCommit.Round,
					Votes: req.ProposedLastCommit.Votes,
				},
				ByzantineValidators: req.Misbehavior,
			},
		)
		app.optimisticProcessingInfo.BeginBlockResult = &optimisticBeginBlockResult
		app.optimisticProcessingInfo.BeginBlockResultCompletion <- struct{}{}
	}
	app.logger.Info("start tx")
	app.logger.Info(fmt.Sprintf("current txs size=%d", req.Size()))
	app.logger.Info(fmt.Sprintf("current txs lenth=%d", len(req.Txs)))
	start := time.Now().UnixMilli()
	if app.optimisticProcessingInfo != nil {
		results, _ := app.buildDependenciesAndRunTxs(app.processProposalState.ctx, req.Txs)
		app.optimisticProcessingInfo.ResponseDeliverTxs = results
		app.optimisticProcessingInfo.Completion <- struct{}{}
	}

	//app.processProposalState.ctx = ctxCurrent
	/*for _, tx := range req.Txs {
		app.optimisticProcessingInfo.DeliverTxResult <- app.OptimisticDeliverTx(
			abci.RequestDeliverTx{
				Tx: tx,
			}, app.processProposalState.ctx,
		)
	}*/
	end := time.Now().UnixMilli()
	app.logger.Info(fmt.Sprintf("txs executr time=%d", end-start))
	app.logger.Info("end tx")
	if app.optimisticProcessingInfo != nil {
		optimisticEndBlockResult := app.optimisticEndBlock(
			abci.RequestEndBlock{
				Height: req.Height,
			},
		)
		app.optimisticProcessingInfo.EndBlockResult = &optimisticEndBlockResult
		app.optimisticProcessingInfo.EndBlockResultCompletion <- struct{}{}
	}
}

// BeginBlock implements the ABCI application interface.
func (app *BaseApp) BeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	/*if app.optimisticProcessingInfo != nil && !app.optimisticProcessingInfo.Aborted && bytes.Equal(app.optimisticProcessingInfo.Hash, req.Hash) {
		return <-app.optimisticProcessingInfo.BeginBlockResult
	}*/

	defer telemetry.MeasureSince(time.Now(), "abci", "begin_block")

	if app.cms.TracingEnabled() {
		app.cms.SetTracingContext(sdk.TraceContext(
			map[string]interface{}{"blockHeight": req.Header.Height},
		))
	}

	if err := app.validateHeight(req); err != nil {
		panic(err)
	}

	// Initialize the DeliverTx state. If this is the first block, it should
	// already be initialized in InitChain. Otherwise app.deliverState will be
	// nil, since it is reset on Commit.
	header := tmproto.Header{
		ChainID:         req.Header.ChainID,
		Height:          req.Header.Height,
		Time:            req.Header.Time,
		ProposerAddress: req.Header.ProposerAddress,
	}
	if app.deliverState == nil {
		app.setDeliverState(header)
	} else {
		// In the first block, app.deliverState.ctx will already be initialized
		// by InitChain. Context is now updated with Header information.
		app.deliverState.ctx = app.deliverState.ctx.
			WithBlockHeader(header).
			WithBlockHeight(header.Height)
	}
	// add block gas meter
	var gasMeter sdk.GasMeter
	if maxGas := app.getMaximumBlockGas(app.deliverState.ctx); maxGas > 0 {
		gasMeter = sdk.NewGasMeter(maxGas)
	} else {
		gasMeter = sdk.NewInfiniteGasMeter()
	}

	// NOTE: header hash is not set in NewContext, so we manually set it here

	app.deliverState.ctx = app.deliverState.ctx.
		WithBlockGasMeter(gasMeter).
		WithHeaderHash(req.Hash).
		WithConsensusParams(app.GetConsensusParams(app.deliverState.ctx))

	// we also set block gas meter to checkState in case the application needs to
	// verify gas consumption during (Re)CheckTx
	if app.checkState != nil {
		app.checkState.ctx = app.checkState.ctx.
			WithBlockGasMeter(gasMeter).
			WithHeaderHash(req.Hash)
	}

	if app.beginBlocker != nil {
		res = app.beginBlocker(app.deliverState.ctx, req)
		res.Events = sdk.MarkEventsToIndex(res.Events, app.indexEvents)
	}
	// set the signed validators for addition to context in deliverTx
	app.voteInfos = req.LastCommitInfo.GetVotes()
	return res
}

// EndBlock implements the ABCI interface.
func (app *BaseApp) EndBlock(req abci.RequestEndBlock) (res abci.ResponseEndBlock) {
	/*if app.optimisticProcessingInfo != nil && !app.optimisticProcessingInfo.Aborted {
		return <-app.optimisticProcessingInfo.EndBlockResult
	}*/

	defer telemetry.MeasureSince(time.Now(), "abci", "end_block")

	if app.deliverState.ms.TracingEnabled() {
		app.deliverState.ms = app.deliverState.ms.SetTracingContext(nil).(sdk.CacheMultiStore)
	}

	if app.endBlocker != nil {
		res = app.endBlocker(app.deliverState.ctx, req)
		res.Events = sdk.MarkEventsToIndex(res.Events, app.indexEvents)
	}

	if cp := app.GetConsensusParams(app.deliverState.ctx); cp != nil {
		res.ConsensusParamUpdates = cp
	}
	app.stateToCommit = app.deliverState

	return res
}

// CheckTx implements the ABCI interface and executes a tx in CheckTx mode. In
// CheckTx mode, messages are not executed. This means messages are only validated
// and only the AnteHandler is executed. State is persisted to the BaseApp's
// internal CheckTx state if the AnteHandler passes. Otherwise, the ResponseCheckTx
// will contain releveant error information. Regardless of tx execution outcome,
// the ResponseCheckTx will contain relevant gas execution context.
func (app *BaseApp) CheckTx(req abci.RequestCheckTx) abci.ResponseCheckTx {
	defer telemetry.MeasureSince(time.Now(), "abci", "check_tx")

	var mode runTxMode

	switch {
	case req.Type == abci.CheckTxType_New:
		mode = runTxModeCheck

	case req.Type == abci.CheckTxType_Recheck:
		mode = runTxModeReCheck

	default:
		panic(fmt.Sprintf("unknown RequestCheckTx type: %s", req.Type))
	}

	gInfo, result, anteEvents, err := app.runTx(mode, req.Tx, app.checkState.ctx)
	if err != nil {
		return sdkerrors.ResponseCheckTxWithEvents(err, gInfo.GasWanted, gInfo.GasUsed, anteEvents, app.trace)
	}

	return abci.ResponseCheckTx{
		GasWanted: int64(gInfo.GasWanted), // TODO: Should type accept unsigned ints?
		GasUsed:   int64(gInfo.GasUsed),   // TODO: Should type accept unsigned ints?
		Log:       result.Log,
		Data:      result.Data,
		Events:    sdk.MarkEventsToIndex(result.Events, app.indexEvents),
	}
}

// DeliverTx implements the ABCI interface and executes a tx in DeliverTx mode.
// State only gets persisted if all messages are valid and get executed successfully.
// Otherwise, the ResponseDeliverTx will contain releveant error information.
// Regardless of tx execution outcome, the ResponseDeliverTx will contain relevant
// gas execution context.
func (app *BaseApp) DeliverTx(req abci.RequestDeliverTx) abci.ResponseDeliverTx {
	defer telemetry.MeasureSince(time.Now(), "abci", "deliver_tx")

	/*if app.optimisticProcessingInfo != nil && !app.optimisticProcessingInfo.Aborted {
		return <-app.optimisticProcessingInfo.DeliverTxResult
	}*/

	gInfo := sdk.GasInfo{}
	resultStr := "successful"

	defer func() {
		telemetry.IncrCounter(1, "tx", "count")
		telemetry.IncrCounter(1, "tx", resultStr)
		telemetry.SetGauge(float32(gInfo.GasUsed), "tx", "gas", "used")
		telemetry.SetGauge(float32(gInfo.GasWanted), "tx", "gas", "wanted")
	}()

	gInfo, result, anteEvents, err := app.runTx(runTxModeDeliver, req.Tx, app.deliverState.ctx)
	if err != nil {
		resultStr = "failed"
		return sdkerrors.ResponseDeliverTxWithEvents(err, gInfo.GasWanted, gInfo.GasUsed, anteEvents, app.trace)
	}

	return abci.ResponseDeliverTx{
		GasWanted: int64(gInfo.GasWanted), // TODO: Should type accept unsigned ints?
		GasUsed:   int64(gInfo.GasUsed),   // TODO: Should type accept unsigned ints?
		Log:       result.Log,
		Data:      result.Data,
		Events:    sdk.MarkEventsToIndex(result.Events, app.indexEvents),
	}
}

func (app *BaseApp) DeliverTxGenUtil(req abci.RequestDeliverTx, ctx sdk.Context) abci.ResponseDeliverTx {
	defer telemetry.MeasureSince(time.Now(), "abci", "deliver_tx")

	/*if app.optimisticProcessingInfo != nil && !app.optimisticProcessingInfo.Aborted {
		return <-app.optimisticProcessingInfo.DeliverTxResult
	}*/

	gInfo := sdk.GasInfo{}
	resultStr := "successful"
	cache := ctx.MultiStore().CacheMultiStore()
	ctx = ctx.WithCacheMultiStore(cache)
	defer func() {
		telemetry.IncrCounter(1, "tx", "count")
		telemetry.IncrCounter(1, "tx", resultStr)
		telemetry.SetGauge(float32(gInfo.GasUsed), "tx", "gas", "used")
		telemetry.SetGauge(float32(gInfo.GasWanted), "tx", "gas", "wanted")
	}()

	gInfo, result, anteEvents, err := app.runTx(runTxModeDeliver, req.Tx, ctx)
	if err != nil {
		resultStr = "failed"
		return sdkerrors.ResponseDeliverTxWithEvents(err, gInfo.GasWanted, gInfo.GasUsed, anteEvents, app.trace)
	}
	cache.Write()
	return abci.ResponseDeliverTx{
		GasWanted: int64(gInfo.GasWanted), // TODO: Should type accept unsigned ints?
		GasUsed:   int64(gInfo.GasUsed),   // TODO: Should type accept unsigned ints?
		Log:       result.Log,
		Data:      result.Data,
		Events:    sdk.MarkEventsToIndex(result.Events, app.indexEvents),
	}
}

func (app *BaseApp) FinalizeBlocker(blocker abci.RequestFinalizeBlocker) abci.ResponseFinalizeBlocker {
	defer telemetry.MeasureSince(time.Now(), "abci", "Finalize_Blocker")
	result := abci.ResponseFinalizeBlocker{}
	if app.optimisticProcessingInfo != nil {
		app.Logger().Info("optimistic processing FinalizeBlocker")
		if bytes.Equal(app.optimisticProcessingInfo.Hash, blocker.Hash) {
			select {
			case <-app.optimisticProcessingInfo.BeginBlockResultCompletion:
				app.Logger().Info("optimistic processing recive BeginBlock")
				result.ResponseBeginBlock = app.optimisticProcessingInfo.BeginBlockResult
			case <-time.After(OptimisticProcessingTimeoutInSeconds * time.Second):
				app.Logger().Info("optimistic processing timed out")
				break
			}
		}
		select {
		case <-app.optimisticProcessingInfo.Completion:
			app.Logger().Info("optimistic processing recive Completion")
			if len(app.optimisticProcessingInfo.ResponseDeliverTxs) > 0 {
				app.Logger().Info("optimistic processing recive ResponseDeliverTxs" + strconv.Itoa(len(app.optimisticProcessingInfo.ResponseDeliverTxs)))
				app.Logger().Info(string(len(app.optimisticProcessingInfo.ResponseDeliverTxs)))
				result.ResponseDeliverTx = app.optimisticProcessingInfo.ResponseDeliverTxs
				/*result := abci.ResponseFinalizeBlocker{
					Height:            blocker.Height,
					ResponseDeliverTx: app.optimisticProcessingInfo.ResponseDeliverTxs,
				}
				return result*/
			}
		case <-time.After(OptimisticProcessingTimeoutInSeconds * time.Second):
			app.Logger().Info("optimistic processing timed out")
			break
		}

		select {
		case <-app.optimisticProcessingInfo.EndBlockResultCompletion:
			app.Logger().Info("optimistic processing recive EndBlock")
			result.ResponseEndBlock = app.optimisticProcessingInfo.EndBlockResult
		case <-time.After(OptimisticProcessingTimeoutInSeconds * time.Second):
			app.Logger().Info("optimistic processing timed out")
			break
		}

	}
	return result
}

// Commit implements the ABCI interface. It will commit all state that exists in
// the deliver state's multi-store and includes the resulting commit ID in the
// returned abci.ResponseCommit. Commit will set the check state based on the
// latest header and reset the deliver state. Also, if a non-zero halt height is
// defined in config, Commit will execute a deferred function call to check
// against that height and gracefully halt if it matches the latest committed
// height.
func (app *BaseApp) Commit() (res abci.ResponseCommit) {
	defer telemetry.MeasureSince(time.Now(), "abci", "commit")

	header := app.stateToCommit.ctx.BlockHeader()
	retainHeight := app.GetBlockRetentionHeight(header.Height)

	// Write the DeliverTx state into branched storage and commit the MultiStore.
	// The write to the DeliverTx state writes all state transitions to the root
	// MultiStore (app.cms) so when Commit() is called is persists those values.
	app.stateToCommit.ms.Write()

	commitID := app.cms.Commit()
	app.logger.Info("commit synced", "commit", fmt.Sprintf("%X", commitID))

	// Reset the Check state to the latest committed.
	//
	// NOTE: This is safe because Tendermint holds a lock on the mempool for
	// Commit. Use the header from this latest block.
	app.setCheckState(header)

	// empty/reset the deliver state
	app.deliverState = nil
	app.processProposalState = nil
	app.stateToCommit = nil
	app.optimisticProcessingInfo = nil

	var halt bool

	switch {
	case app.haltHeight > 0 && uint64(header.Height) >= app.haltHeight:
		halt = true

	case app.haltTime > 0 && header.Time.Unix() >= int64(app.haltTime):
		halt = true
	}

	if halt {
		// Halt the binary and allow Tendermint to receive the ResponseCommit
		// response with the commit ID hash. This will allow the node to successfully
		// restart and process blocks assuming the halt configuration has been
		// reset or moved to a more distant value.
		app.halt()
	}

	if app.snapshotInterval > 0 && uint64(header.Height)%app.snapshotInterval == 0 {
		go app.snapshot(header.Height)
	}

	return abci.ResponseCommit{
		Data:         commitID.Hash,
		RetainHeight: retainHeight,
	}
}

// halt attempts to gracefully shutdown the node via SIGINT and SIGTERM falling
// back on os.Exit if both fail.
func (app *BaseApp) halt() {
	app.logger.Info("halting node per configuration", "height", app.haltHeight, "time", app.haltTime)

	p, err := os.FindProcess(os.Getpid())
	if err == nil {
		// attempt cascading signals in case SIGINT fails (os dependent)
		sigIntErr := p.Signal(syscall.SIGINT)
		sigTermErr := p.Signal(syscall.SIGTERM)

		if sigIntErr == nil || sigTermErr == nil {
			return
		}
	}

	// Resort to exiting immediately if the process could not be found or killed
	// via SIGINT/SIGTERM signals.
	app.logger.Info("failed to send SIGINT/SIGTERM; exiting...")
	os.Exit(0)
}

// snapshot takes a snapshot of the current state and prunes any old snapshottypes.
func (app *BaseApp) snapshot(height int64) {
	if app.snapshotManager == nil {
		app.logger.Info("snapshot manager not configured")
		return
	}

	app.logger.Info("creating state snapshot", "height", height)

	snapshot, err := app.snapshotManager.Create(uint64(height))
	if err != nil {
		app.logger.Error("failed to create state snapshot", "height", height, "err", err)
		return
	}

	app.logger.Info("completed state snapshot", "height", height, "format", snapshot.Format)

	if app.snapshotKeepRecent > 0 {
		app.logger.Debug("pruning state snapshots")

		pruned, err := app.snapshotManager.Prune(app.snapshotKeepRecent)
		if err != nil {
			app.logger.Error("Failed to prune state snapshots", "err", err)
			return
		}

		app.logger.Debug("pruned state snapshots", "pruned", pruned)
	}
}

// Query implements the ABCI interface. It delegates to CommitMultiStore if it
// implements Queryable.
func (app *BaseApp) Query(req abci.RequestQuery) (res abci.ResponseQuery) {
	defer telemetry.MeasureSince(time.Now(), "abci", "query")

	// Add panic recovery for all queries.
	// ref: https://github.com/cosmos/cosmos-sdk/pull/8039
	defer func() {
		if r := recover(); r != nil {
			res = sdkerrors.QueryResult(sdkerrors.Wrapf(sdkerrors.ErrPanic, "%v", r))
		}
	}()

	// when a client did not provide a query height, manually inject the latest
	if req.Height == 0 {
		req.Height = app.LastBlockHeight()
	}

	// handle gRPC routes first rather than calling splitPath because '/' characters
	// are used as part of gRPC paths
	if grpcHandler := app.grpcQueryRouter.Route(req.Path); grpcHandler != nil {
		return app.handleQueryGRPC(grpcHandler, req)
	}

	path := splitPath(req.Path)
	if len(path) == 0 {
		sdkerrors.QueryResult(sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no query path provided"))
	}

	switch path[0] {
	// "/app" prefix for special application queries
	case "app":
		return handleQueryApp(app, path, req)

	case "store":
		return handleQueryStore(app, path, req)

	case "p2p":
		return handleQueryP2P(app, path)

	case "custom":
		return handleQueryCustom(app, path, req)
	}

	return sdkerrors.QueryResult(sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown query path"))
}

// ListSnapshots implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) ListSnapshots(req abci.RequestListSnapshots) abci.ResponseListSnapshots {
	resp := abci.ResponseListSnapshots{Snapshots: []*abci.Snapshot{}}
	if app.snapshotManager == nil {
		return resp
	}

	snapshots, err := app.snapshotManager.List()
	if err != nil {
		app.logger.Error("failed to list snapshots", "err", err)
		return resp
	}

	for _, snapshot := range snapshots {
		abciSnapshot, err := snapshot.ToABCI()
		if err != nil {
			app.logger.Error("failed to list snapshots", "err", err)
			return resp
		}
		resp.Snapshots = append(resp.Snapshots, &abciSnapshot)
	}

	return resp
}

// LoadSnapshotChunk implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) LoadSnapshotChunk(req abci.RequestLoadSnapshotChunk) abci.ResponseLoadSnapshotChunk {
	if app.snapshotManager == nil {
		return abci.ResponseLoadSnapshotChunk{}
	}
	chunk, err := app.snapshotManager.LoadChunk(req.Height, req.Format, req.Chunk)
	if err != nil {
		app.logger.Error(
			"failed to load snapshot chunk",
			"height", req.Height,
			"format", req.Format,
			"chunk", req.Chunk,
			"err", err,
		)
		return abci.ResponseLoadSnapshotChunk{}
	}
	return abci.ResponseLoadSnapshotChunk{Chunk: chunk}
}

// OfferSnapshot implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) OfferSnapshot(req abci.RequestOfferSnapshot) abci.ResponseOfferSnapshot {
	if app.snapshotManager == nil {
		app.logger.Error("snapshot manager not configured")
		return abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_ABORT}
	}

	if req.Snapshot == nil {
		app.logger.Error("received nil snapshot")
		return abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_REJECT}
	}

	snapshot, err := snapshottypes.SnapshotFromABCI(req.Snapshot)
	if err != nil {
		app.logger.Error("failed to decode snapshot metadata", "err", err)
		return abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_REJECT}
	}

	err = app.snapshotManager.Restore(snapshot)
	switch {
	case err == nil:
		return abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_ACCEPT}

	case errors.Is(err, snapshottypes.ErrUnknownFormat):
		return abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_REJECT_FORMAT}

	case errors.Is(err, snapshottypes.ErrInvalidMetadata):
		app.logger.Error(
			"rejecting invalid snapshot",
			"height", req.Snapshot.Height,
			"format", req.Snapshot.Format,
			"err", err,
		)
		return abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_REJECT}

	default:
		app.logger.Error(
			"failed to restore snapshot",
			"height", req.Snapshot.Height,
			"format", req.Snapshot.Format,
			"err", err,
		)

		// We currently don't support resetting the IAVL stores and retrying a different snapshot,
		// so we ask Tendermint to abort all snapshot restoration.
		return abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_ABORT}
	}
}

// ApplySnapshotChunk implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) ApplySnapshotChunk(req abci.RequestApplySnapshotChunk) abci.ResponseApplySnapshotChunk {
	if app.snapshotManager == nil {
		app.logger.Error("snapshot manager not configured")
		return abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_ABORT}
	}

	_, err := app.snapshotManager.RestoreChunk(req.Chunk)
	switch {
	case err == nil:
		return abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_ACCEPT}

	case errors.Is(err, snapshottypes.ErrChunkHashMismatch):
		app.logger.Error(
			"chunk checksum mismatch; rejecting sender and requesting refetch",
			"chunk", req.Index,
			"sender", req.Sender,
			"err", err,
		)
		return abci.ResponseApplySnapshotChunk{
			Result:        abci.ResponseApplySnapshotChunk_RETRY,
			RefetchChunks: []uint32{req.Index},
			RejectSenders: []string{req.Sender},
		}

	default:
		app.logger.Error("failed to restore snapshot", "err", err)
		return abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_ABORT}
	}
}

func (app *BaseApp) handleQueryGRPC(handler GRPCQueryHandler, req abci.RequestQuery) abci.ResponseQuery {
	ctx, err := app.createQueryContext(req.Height, req.Prove)
	if err != nil {
		return sdkerrors.QueryResult(err)
	}

	res, err := handler(ctx, req)
	if err != nil {
		res = sdkerrors.QueryResult(gRPCErrorToSDKError(err))
		res.Height = req.Height
		return res
	}

	return res
}

func gRPCErrorToSDKError(err error) error {
	status, ok := grpcstatus.FromError(err)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	switch status.Code() {
	case codes.NotFound:
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, err.Error())
	case codes.InvalidArgument:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	case codes.FailedPrecondition:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	case codes.Unauthenticated:
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	default:
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}
}

func checkNegativeHeight(height int64) error {
	if height < 0 {
		// Reject invalid heights.
		return sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			"cannot query with height < 0; please provide a valid height",
		)
	}
	return nil
}

// createQueryContext creates a new sdk.Context for a query, taking as args
// the block height and whether the query needs a proof or not.
func (app *BaseApp) createQueryContext(height int64, prove bool) (sdk.Context, error) {
	if err := checkNegativeHeight(height); err != nil {
		return sdk.Context{}, err
	}

	// when a client did not provide a query height, manually inject the latest
	if height == 0 {
		height = app.LastBlockHeight()
	}

	if height <= 1 && prove {
		return sdk.Context{},
			sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"cannot query with proof when height <= 1; please provide a valid height",
			)
	}

	cacheMS, err := app.cms.CacheMultiStoreWithVersion(height)
	if err != nil {
		return sdk.Context{},
			sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"failed to load state at height %d; %s (latest height: %d)", height, err, app.LastBlockHeight(),
			)
	}

	// branch the commit-multistore for safety
	ctx := sdk.NewContext(
		cacheMS, app.checkState.ctx.BlockHeader(), true, app.logger,
	).WithMinGasPrices(app.minGasPrices).WithBlockHeight(height)

	return ctx, nil
}

// GetBlockRetentionHeight returns the height for which all blocks below this height
// are pruned from Tendermint. Given a commitment height and a non-zero local
// minRetainBlocks configuration, the retentionHeight is the smallest height that
// satisfies:
//
// - Unbonding (safety threshold) time: The block interval in which validators
// can be economically punished for misbehavior. Blocks in this interval must be
// auditable e.g. by the light client.
//
// - Logical store snapshot interval: The block interval at which the underlying
// logical store database is persisted to disk, e.g. every 10000 heights. Blocks
// since the last IAVL snapshot must be available for replay on application restart.
//
// - State sync snapshots: Blocks since the oldest available snapshot must be
// available for state sync nodes to catch up (oldest because a node may be
// restoring an old snapshot while a new snapshot was taken).
//
// - Local (minRetainBlocks) config: Archive nodes may want to retain more or
// all blocks, e.g. via a local config option min-retain-blocks. There may also
// be a need to vary retention for other nodes, e.g. sentry nodes which do not
// need historical blocks.
func (app *BaseApp) GetBlockRetentionHeight(commitHeight int64) int64 {
	// pruning is disabled if minRetainBlocks is zero
	if app.minRetainBlocks == 0 {
		return 0
	}

	minNonZero := func(x, y int64) int64 {
		switch {
		case x == 0:
			return y
		case y == 0:
			return x
		case x < y:
			return x
		default:
			return y
		}
	}

	// Define retentionHeight as the minimum value that satisfies all non-zero
	// constraints. All blocks below (commitHeight-retentionHeight) are pruned
	// from Tendermint.
	var retentionHeight int64

	// Define the number of blocks needed to protect against misbehaving validators
	// which allows light clients to operate safely. Note, we piggy back of the
	// evidence parameters instead of computing an estimated nubmer of blocks based
	// on the unbonding period and block commitment time as the two should be
	// equivalent.
	cp := app.GetConsensusParams(app.deliverState.ctx)
	if cp != nil && cp.Evidence != nil && cp.Evidence.MaxAgeNumBlocks > 0 {
		retentionHeight = commitHeight - cp.Evidence.MaxAgeNumBlocks
	}

	// Define the state pruning offset, i.e. the block offset at which the
	// underlying logical database is persisted to disk.
	statePruningOffset := int64(app.cms.GetPruning().KeepEvery)
	if statePruningOffset > 0 {
		if commitHeight > statePruningOffset {
			v := commitHeight - (commitHeight % statePruningOffset)
			retentionHeight = minNonZero(retentionHeight, v)
		} else {
			// Hitting this case means we have persisting enabled but have yet to reach
			// a height in which we persist state, so we return zero regardless of other
			// conditions. Otherwise, we could end up pruning blocks without having
			// any state committed to disk.
			return 0
		}
	}

	if app.snapshotInterval > 0 && app.snapshotKeepRecent > 0 {
		v := commitHeight - int64(app.snapshotInterval*uint64(app.snapshotKeepRecent))
		retentionHeight = minNonZero(retentionHeight, v)
	}

	v := commitHeight - int64(app.minRetainBlocks)
	retentionHeight = minNonZero(retentionHeight, v)

	if retentionHeight <= 0 {
		// prune nothing in the case of a non-positive height
		return 0
	}

	return retentionHeight
}

func handleQueryApp(app *BaseApp, path []string, req abci.RequestQuery) abci.ResponseQuery {
	if len(path) >= 2 {
		switch path[1] {
		case "simulate":
			txBytes := req.Data

			gInfo, res, err := app.Simulate(txBytes)
			if err != nil {
				return sdkerrors.QueryResult(sdkerrors.Wrap(err, "failed to simulate tx"))
			}

			simRes := &sdk.SimulationResponse{
				GasInfo: gInfo,
				Result:  res,
			}

			bz, err := codec.ProtoMarshalJSON(simRes, app.interfaceRegistry)
			if err != nil {
				return sdkerrors.QueryResult(sdkerrors.Wrap(err, "failed to JSON encode simulation response"))
			}

			return abci.ResponseQuery{
				Codespace: sdkerrors.RootCodespace,
				Height:    req.Height,
				Value:     bz,
			}

		case "version":
			return abci.ResponseQuery{
				Codespace: sdkerrors.RootCodespace,
				Height:    req.Height,
				Value:     []byte(app.version),
			}

		default:
			return sdkerrors.QueryResult(sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query: %s", path))
		}
	}

	return sdkerrors.QueryResult(
		sdkerrors.Wrap(
			sdkerrors.ErrUnknownRequest,
			"expected second parameter to be either 'simulate' or 'version', neither was present",
		),
	)
}

func handleQueryStore(app *BaseApp, path []string, req abci.RequestQuery) abci.ResponseQuery {
	// "/store" prefix for store queries
	queryable, ok := app.cms.(sdk.Queryable)
	if !ok {
		return sdkerrors.QueryResult(sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "multistore doesn't support queries"))
	}

	req.Path = "/" + strings.Join(path[1:], "/")

	if req.Height <= 1 && req.Prove {
		return sdkerrors.QueryResult(
			sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"cannot query with proof when height <= 1; please provide a valid height",
			),
		)
	}

	resp := queryable.Query(req)
	resp.Height = req.Height

	return resp
}

func handleQueryP2P(app *BaseApp, path []string) abci.ResponseQuery {
	// "/p2p" prefix for p2p queries
	if len(path) < 4 {
		return sdkerrors.QueryResult(sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "path should be p2p filter <addr|id> <parameter>"))
	}

	ctx, err := app.createQueryContext(0, false)
	if err != nil {
		return sdkerrors.QueryResult(err)
	}

	var resp abci.ResponseQuery
	cmd, typ, arg := path[1], path[2], path[3]
	switch cmd {
	case "filter":
		switch typ {
		case "addr":
			resp = app.FilterPeerByAddrPort(ctx, arg)

		case "id":
			resp = app.FilterPeerByID(ctx, arg)
		}

	default:
		resp = sdkerrors.QueryResult(sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "expected second parameter to be 'filter'"))
	}

	return resp
}

func handleQueryCustom(app *BaseApp, path []string, req abci.RequestQuery) abci.ResponseQuery {
	// path[0] should be "custom" because "/custom" prefix is required for keeper
	// queries.
	//
	// The QueryRouter routes using path[1]. For example, in the path
	// "custom/gov/proposal", QueryRouter routes using "gov".
	if len(path) < 2 || path[1] == "" {
		return sdkerrors.QueryResult(sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no route for custom query specified"))
	}

	querier := app.queryRouter.Route(path[1])
	if querier == nil {
		return sdkerrors.QueryResult(sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "no custom querier found for route %s", path[1]))
	}

	ctx, err := app.createQueryContext(req.Height, req.Prove)
	if err != nil {
		return sdkerrors.QueryResult(err)
	}

	// Passes the rest of the path as an argument to the querier.
	//
	// For example, in the path "custom/gov/proposal/test", the gov querier gets
	// []string{"proposal", "test"} as the path.
	resBytes, err := querier(ctx, path[2:], req)
	if err != nil {
		res := sdkerrors.QueryResult(err)
		res.Height = req.Height
		return res
	}

	return abci.ResponseQuery{
		Height: req.Height,
		Value:  resBytes,
	}
}

// splitPath splits a string path using the delimiter '/'.
//
// e.g. "this/is/funny" becomes []string{"this", "is", "funny"}
func splitPath(requestPath string) (path []string) {
	path = strings.Split(requestPath, "/")

	// first element is empty string
	if len(path) > 0 && path[0] == "" {
		path = path[1:]
	}

	return path
}

func (app *BaseApp) optimisticBeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	if app.beginBlocker != nil {
		res = app.beginBlocker(app.processProposalState.ctx, req)
		res.Events = sdk.MarkEventsToIndex(res.Events, app.indexEvents)
	}
	// set the signed validators for addition to context in deliverTx
	app.voteInfos = req.LastCommitInfo.GetVotes()
	return res
}

func (app *BaseApp) OptimisticDeliverTx(req abci.RequestDeliverTx, ctx sdk.Context) abci.ResponseDeliverTx {
	gInfo := sdk.GasInfo{}
	resultStr := "successful"

	defer func() {
		telemetry.IncrCounter(1, "tx", "count")
		telemetry.IncrCounter(1, "tx", resultStr)
		telemetry.SetGauge(float32(gInfo.GasUsed), "tx", "gas", "used")
		telemetry.SetGauge(float32(gInfo.GasWanted), "tx", "gas", "wanted")
	}()

	gInfo, result, anteEvents, err := app.runTx(runTxModeDeliver, req.Tx, ctx)
	if err != nil {
		resultStr = "failed"
		return sdkerrors.ResponseDeliverTxWithEvents(err, gInfo.GasWanted, gInfo.GasUsed, anteEvents, app.trace)
	}

	return abci.ResponseDeliverTx{
		GasWanted: int64(gInfo.GasWanted), // TODO: Should type accept unsigned ints?
		GasUsed:   int64(gInfo.GasUsed),   // TODO: Should type accept unsigned ints?
		Log:       result.Log,
		Data:      result.Data,
		Events:    sdk.MarkEventsToIndex(result.Events, app.indexEvents),
	}
}

func (app *BaseApp) optimisticEndBlock(req abci.RequestEndBlock) (res abci.ResponseEndBlock) {

	if app.processProposalState.ms.TracingEnabled() {
		app.processProposalState.ms = app.processProposalState.ms.SetTracingContext(nil).(sdk.CacheMultiStore)
	}

	if app.endBlocker != nil {
		res = app.endBlocker(app.processProposalState.ctx, req)
		res.Events = sdk.MarkEventsToIndex(res.Events, app.indexEvents)
	}

	if cp := app.GetConsensusParams(app.processProposalState.ctx); cp != nil {
		res.ConsensusParamUpdates = cp
	}
	app.stateToCommit = app.processProposalState
	return res
}
