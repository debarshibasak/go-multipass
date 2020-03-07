package multipass

import (
	"log"
	"os/exec"
)

type ExecRequest struct {
	Name string
	Command string
}

func Exec(req *ExecRequest) error {

	cmdString := "multipass exec "+req.Name+" -- "+req.Command

	cmd := exec.Command("sh", "-c", cmdString)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return err
	}

	return nil
}
