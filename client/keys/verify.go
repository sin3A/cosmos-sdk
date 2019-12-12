package keys

import (
	"bufio"
	"errors"
	"io/ioutil"

	"github.com/spf13/cobra"
)

func verifyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify <name> <filename>",
		Short: "Verify an arbitrary message signature",
		Long: `Read a document generated with the 'key sign' command and verify the signature.
It exits with 0 if the signature verification succeed; it returns a value different than 0
if the signature verification fails.
`,
		Args: cobra.ExactArgs(1),
		RunE: runVerifyCmd,
	}
	return cmd
}

func runVerifyCmd(cmd *cobra.Command, args []string) error {
	name := args[0]
	filename := args[1]

	kb, err := NewKeyringFromHomeFlag(bufio.NewReader(cmd.InOrStdin()))
	if err != nil {
		return err
	}

	info, err := kb.Get(name)
	if err != nil {
		return err
	}

	signedDoc, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var doc signedMsg
	if err := cdc.UnmarshalJSON(signedDoc, &doc); err != nil {
		return err
	}

	if !info.GetPubKey().VerifyBytes(doc.Bytes(), doc.Sig) {
		return errors.New("bad signature")
	}

	return nil
}
