package multipass

import (
	"os/exec"
)

type InfoRequest struct {
	Name string
}

func Info(req *InfoRequest) (*Instance, error) {

	cmdFormat := "multipass info " + req.Name

	cmdExec := exec.Command("sh", "-c", cmdFormat)

	out, err := cmdExec.Output()
	if err != nil {
		return nil, err
	}

	return parseInstance(string(out)), nil
}
