package multipass

import (
	"fmt"
	"testing"
)

func TestCreateInfoAndDelete(t *testing.T) {
	instance, err := Launch(&LaunchReq{
		CPUS: "2",
		Name: "test2",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("created instance " + instance.IP)

	instanceInfo, err := Info(&InfoRequest{Name: instance.Name})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("instance information " + instanceInfo.Name)
	t.Log("instance information " + instanceInfo.IP)

	err = Exec(&ExecRequest{Name: instanceInfo.Name, Command: "ls"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("now deleting " + instance.Name)

	if err := Delete(&DeleteRequest{Name: instance.Name}); err != nil {
		t.Fatal(err)
	}

}

func TestParsing(t *testing.T) {

	instance := parseInstance(`Name:           perennial-trogon
State:          Running
IPv4:           192.168.64.55
Release:        Ubuntu 18.04.4 LTS
Image hash:     3c3a67a14257 (Ubuntu 18.04 LTS)
Load:           0.06 0.07 0.02
Disk usage:     988.2M out of 4.7G
Memory usage:   83.6M out of 985.6M`)

	fmt.Printf(instance.IP)

}
