package multipass

import (
	"errors"
	"os/exec"
	"strings"
)

type LaunchReq struct {
	Image         string
	CPUS          string
	Disk          string
	Name          string
	Memory        string
	CloudInitFile string
}

func Launch(launchReq *LaunchReq) (*Instance, error) {
	var args = []string{"launch"}

	if launchReq.Image != "" {
		args = append(args, launchReq.Image)
	}

	if launchReq.CPUS != "" {
		args = append(args, "--cpus", launchReq.CPUS)
	}

	if launchReq.Name != "" {
		args = append(args, "--name", launchReq.Name)
	}

	if launchReq.Disk != "" {
		args = append(args, "--disk", launchReq.Disk)
	}

	if launchReq.Memory != "" {
		args = append(args, "--mem", launchReq.Memory)
	}

	if launchReq.CloudInitFile != "" {
		args = append(args, "--cloud-init", launchReq.CloudInitFile)
	}

	result := exec.Command("multipass", args...)
	out, err := result.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out) + " " + err.Error())
	}

	var b []byte
	b = append(b, out...)

	out2 := string(b)

	o := strings.Split(strings.TrimSpace(out2), "\n")[0]

	name := strings.TrimSpace(strings.Split(o, "Launched: ")[1])

	instance, err := Info(&InfoRequest{Name: name})
	if err != nil {
		return nil, err
	}

	return instance, nil
}
