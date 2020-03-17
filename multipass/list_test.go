package multipass

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {

	instances, err := List()
	if err != nil {
		t.Fatal(err)
	}

	for _, i := range instances {
		fmt.Println(i.IP)
	}
}
