package cli

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	printableASCIIRegexString = "^[\x20-\x7E]*$"
	stringMsgMaxLength        = 256
	objectMsgMaxLength        = 128
)

var printableASCIIRegex = regexp.MustCompile(printableASCIIRegexString)

func GetSignMessageCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-message <name> <chain-id> <filename>",
		Short: "Sign an arbitrary message with a private key and print the signed document to STDOUT",
		Long: `Sign an arbitrary text file with a private key and produce an amino-encoded JSON output.
The signed JSON document could eventually be verified through the 'keys verify' command and will
have the following structure:
{
  "text": original text file contents,
  "pub": public key,
  "sig": signature
}
`,
		Args: cobra.ExactArgs(3),
		RunE: makeRunSignCmd(cdc),
	}
	cmd.SetOut(os.Stdout)
	cmd.Flags().StringP("type", "t", "string", "Message type; can be string or object")
	return cmd
}

func makeRunSignCmd(cdc *codec.Codec) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		name := args[0]
		chainID := args[1]
		filename := args[2]

		payload, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		msgType := viper.GetString("type")
		if msgType == "object" {
			payload = []byte(fmt.Sprintf("%x", sha256.Sum256(payload)))
		}

		msg := sdk.NewArbitraryMsg(msgType, payload)
		if err := msg.ValidateBasic(); err != nil {
			return err
		}

		buf := bufio.NewReader(cmd.InOrStdin())
		kb, err := keys.NewKeyringFromHomeFlag(buf)
		if err != nil {
			return err
		}

		cliCtx := context.NewCLIContext().WithCodec(cdc)
		stdSignature, err := types.MakeSignature(kb, name, "", types.StdSignMsg{
			ChainID: chainID,
			Msgs:    []sdk.Msg{msg},
			Memo:    "",
		})
		if err != nil {
			return err
		}

		cliCtx.PrintOutput(stdSignature)
		return nil
	}
}
