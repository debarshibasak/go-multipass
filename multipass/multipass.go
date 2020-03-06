package multipass

import (
	"fmt"
	"os/exec"
)

type LaunchReq struct {
	CPU int
	Disk string
	Name string
	Memory string
	CloudInitFile string
}

func Launch(launchReq *LaunchReq) error {

	var args = ""

	if launchReq.CPU != 0 {
		args = args + fmt.Sprintf("--cpus %v", launchReq.CPU)
	}

	if launchReq.Name != "" {
		args = args + fmt.Sprintf("--name %v", launchReq.Name)
	}

	if launchReq.Disk != "" {
		args = args + fmt.Sprintf("--disk %v", launchReq.Disk)
	}

	if launchReq.Memory != "" {
		args = args + fmt.Sprintf("--mem %v", launchReq.Memory)
	}

	if launchReq.CloudInitFile != "" {
		args = args + fmt.Sprintf("--cloud-init %v", launchReq.CloudInitFile)
	}

	cmd := fmt.Sprintf("multipass launch "+args)

	cmdExec := exec.Command("sh", "-c", cmd)

	out, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	fmt.Println(string(out))

	return nil
}