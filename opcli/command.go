package opcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Client struct{}

// Takes the primary op command (example: "vault") and arguments (example:
// "list", etc...) to run the op CLI command.
func (c *Client) RunCommand(opCmd string, args []string) ([]byte, error) {
	cArgs := []string{opCmd}
	cArgs = append(cArgs, args...)
	cArgs = append(cArgs, "--format", "json")

	cmd := exec.Command("op", cArgs...)
	errBuf := bytes.NewBuffer(nil)
	cmd.Stderr = errBuf

	out, err := cmd.Output()
	if err != nil {
		if errBuf.String() != "" {
			return nil, fmt.Errorf(errBuf.String())
		}
		return nil, err
	}
	return out, nil
}

// Unmarshals the output JSON of c.RunCommand.
func (c *Client) RunOpUnmarshal(opCmd string, args []string, target any) error {
	out, err := c.RunCommand(opCmd, args)
	if err != nil {
		return err
	}

	err = json.Unmarshal(out, target)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %s", err)
	}

	return nil
}
