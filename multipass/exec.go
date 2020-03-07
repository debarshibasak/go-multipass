package multipass

import (
	"github.com/pkg/errors"
	"log"
	"os/exec"
)

type ExecRequest struct {
	Name string
	Command string
}

func Exec(req *ExecRequest) error {

	if req.Command == "" {
		return errors.New("command cannot be empty")
	}

	cmdString := "multipass exec "+req.Name+" -- "+req.Command

	cmd := exec.Command("sh", "-c", cmdString)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return err
	}

	return nil
}
