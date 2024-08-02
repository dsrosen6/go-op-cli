package opcli

import (
	"bytes"
	"fmt"
	"os/exec"
)

type OpClient struct{}

func runCommand(command []string) ([]byte, error) { // TODO: any better way?
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
	return c
}
