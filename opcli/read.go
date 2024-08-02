package opcli

import "errors"

type ReadOptions struct {
	Reference string
	FileMode  string
	Force     bool
	NoNewLine bool
	OutFile   string
}

func Read(o ReadOptions) (string, error) {
	var args []string
	if o.Reference == "" {
		return "", errors.New("no secret reference provided")
	}

	if o.FileMode != "" { // TODO: FileMode validation, int?
		args = append(args, "--file-mode", o.FileMode)
	}

	if o.Force {
		args = append(args, "-f") // TODO: Find out how this even works
	}

	if o.NoNewLine {
		args = append(args, "-n") // TODO: Can this and FileMode/OutFile be used together?
	}

	if o.OutFile != "" {
		args = append(args, "-o", o.OutFile) // TODO: Directory presence check?
	}

	args = append(args, o.Reference)

	c := createCommand("read", args)
	out, err := runCommand(c)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
