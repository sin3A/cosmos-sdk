package keys

import (
	"bufio"
	"crypto/sha512"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

const (
	printableASCIIRegexString = "^[\x20-\x7E]*$"
	stringMsgMaxLength        = 256
	objectMsgMaxLength        = 128
)

var printableASCIIRegex = regexp.MustCompile(printableASCIIRegexString)

func signCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign <name> <chain-id> <filename>",
		Short: "Sign a plain text payload with a private key and print the signed document to STDOUT",
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
		RunE: runSignCmd,
	}
	cmd.SetOut(os.Stdout)
	cmd.Flags().StringP("type", "t", "string", "Message type; can be string or object")
	return cmd
}

func runSignCmd(cmd *cobra.Command, args []string) error {
	name := args[0]
	chainID := args[1]
	filename := args[2]

	buf := bufio.NewReader(cmd.InOrStdin())
	kb, err := NewKeyringFromHomeFlag(buf)
	if err != nil {
		return err
	}

	payload, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	msgType, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}
	if msgType == "object" {
		hash := sha512.New512_256()
		payload = hash.Sum(payload)
	}

	msg := newSignedMsg(chainID, msgType, payload, nil)
	sig, _, err := kb.Sign(name, "", msg.Bytes())
	if err != nil {
		return err
	}

	msg.Sig = sig
	out, err := cdc.MarshalJSON(msg)
	if err != nil {
		return err
	}

	cmd.Println(string(out))
	return nil
}
