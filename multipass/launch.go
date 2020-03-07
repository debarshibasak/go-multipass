package multipass

import (
	"fmt"
	"os/exec"
	"strings"
)

type LaunchReq struct {
	CPU int
	Disk string
	Name string
	Memory string
	CloudInitFile string
}

func Launch(launchReq *LaunchReq) (*Instance, error) {

	var args = ""

	if launchReq.CPU != 0 {
		args = args + fmt.Sprintf(" --cpus %v", launchReq.CPU)
	}

	if launchReq.Name != "" {
		args = args + fmt.Sprintf(" --name %v", launchReq.Name)
	}

	if launchReq.Disk != "" {
		args = args + fmt.Sprintf(" --disk %v", launchReq.Disk)
	}

	if launchReq.Memory != "" {
		args = args + fmt.Sprintf(" --mem %v", launchReq.Memory)
	}

	if launchReq.CloudInitFile != "" {
		args = args + fmt.Sprintf(" --cloud-init %v", launchReq.CloudInitFile)
	}

	cmd := fmt.Sprintf("multipass launch "+args)

	cmdExec := exec.Command("sh", "-c", cmd)
	out, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return nil, err
	}

	var b []byte
	b = append(b, out...)

	out2 := string(b)

	o := strings.Split(strings.TrimSpace(out2), "\n")[0]

	name := strings.TrimSpace(strings.Split(o, "Launched: ")[1])

	instance, err := Info(&InfoRequest{Name:name})
	if err != nil {
		return nil, err
	}

	return instance, nil
}