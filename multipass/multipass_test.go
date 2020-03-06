package multipass_test

import (
	"github.com/debarshibasak/go-multipass/multipass"
	"testing"
)

func TestCreate(t *testing.T) {
	if err := multipass.Launch(&multipass.LaunchReq{
		CPU:           2,
	}); err != nil {
		t.Fatal(err)
	}
}