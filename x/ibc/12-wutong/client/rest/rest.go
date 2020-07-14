package rest

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/12-wutong/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, queryRoute string) {
	registerTxRoutes(cliCtx, r)
}

// CreateClientReq defines the properties of a create client request's body.
type CreateClientReq struct {
	BaseReq  rest.BaseReq `json:"base_req" yaml:"base_req"`
	ClientID string       `json:"client_id" yaml:"client_id"`
	Header   types.Header `json:"header" yaml:"header"`
}
