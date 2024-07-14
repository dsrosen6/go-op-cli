package opcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Client struct{}

func (op *Client) commandWithUnmarshal(mainCommand string, args []string, target any) error {
	cmd := op.createCommand(mainCommand, args)
	out, err := op.runCommand(cmd)
	if err != nil {
		return err
	}

	if err := op.unmarshalOutput(out, target); err != nil {
		return err
	}

	return nil
}

// Create the command slice that will be passed to op - i.e. "op read", and then args.
func (op *Client) createCommand(mainCommand string, args []string) []string {
	c := []string{"op", mainCommand}
	c = append(c, args...)
	c = append(c, "--format", "json")
	return c
}

func (op *Client) runCommand(command []string) ([]byte, error) {
	cmd := exec.Command(command[0], command[1:]...)
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

func (op *Client) unmarshalOutput(output []byte, target any) error {
	err := json.Unmarshal(output, target)
	if err != nil {
		return fmt.Errorf("error unmarshaling output: %w", err)
	}

	return nil
}
