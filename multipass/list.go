package multipass

import (
	"os/exec"
	"strings"
)

func List() ([]*Instance, error) {

	cmdFormat := "multipass list"

	cmdExec := exec.Command("sh", "-c", cmdFormat)

	out, err := cmdExec.Output()
	if err != nil {
		return nil, err
	}

	return parseInstances(string(out))
}

func parseInstances(out string) ([]*Instance, error) {

	var instances []*Instance
	allLines := strings.Split(strings.TrimSpace(out), "\n")

	for _, line := range allLines {
		space := strings.TrimSpace(strings.Split(line, " ")[0])
		if space != "Name" {
			instance, err := Info(&InfoRequest{Name: space})
			if err != nil {
				return nil, err
			}
			instances = append(instances, instance)
		}
	}

	return instances, nil
}
