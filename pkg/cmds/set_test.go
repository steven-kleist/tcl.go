package cmds

import (
	"testing"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func TestSet(t *testing.T) {
	i := tcl.InitInterp()
	i.RegisterCommand("set", Set, nil)
	res, err := i.Eval("set name Golang-Test")
	if err != nil {
		t.Error(err)
	}
	if res != "Golang-Test" {
		t.Errorf("Expected: Golang-Test, got: %s\n", res)
	}
}

func TestUnset(t *testing.T) {
	i := tcl.InitInterp()
	i.RegisterCommand("set", Set, nil)
	i.RegisterCommand("unset", Unset, nil)

	res, err := i.Eval("set name Golang-Test")
	if err != nil {
		t.Fail()
	}
	if res != "Golang-Test" {
		t.Fail()
	}

	res2, err2 := i.Eval("unset name")
	if err2 != nil || res2 != "" {
		t.Error(res2, err2)
	}
}
