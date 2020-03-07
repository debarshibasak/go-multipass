package multipass

import (
	"fmt"
	"os/exec"
)

type DeleteRequest struct {
	Name string
}

func DeleteAll() error {
	cmd := fmt.Sprintf("multipass delete --all -p")

	cmdExec := exec.Command("sh", "-c", cmd)

	out, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	fmt.Println(string(out))

	return nil
}

func Delete(req *DeleteRequest) error {

	cmd := fmt.Sprintf("multipass delete -p "+req.Name)

	cmdExec := exec.Command("sh", "-c", cmd)

	out, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	fmt.Println(string(out))

	return nil
}
