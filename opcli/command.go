package opcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Client struct{}

func runWithUnmarshal(mainCommand string, args []string, target any) error {
	cmd := createCommand(mainCommand, args)
	out, err := runNoUnmarshal(cmd)
	if err != nil {
		return err
	}

	if err := unmarshalOutput(out, target); err != nil {
		return err
	}

	return nil
}

func runNoUnmarshal(command []string) ([]byte, error) {
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

// Create the command slice that will be passed to op - i.e. "op read", and then args.
func createCommand(mainCommand string, args []string) []string {
	c := []string{"op", mainCommand}
	c = append(c, args...)
	c = append(c, "--format", "json")
	return c
}

func unmarshalOutput(output []byte, target any) error {
	err := json.Unmarshal(output, target)
	if err != nil {
		return fmt.Errorf("error unmarshaling output: %w", err)
	}

	return nil
}
