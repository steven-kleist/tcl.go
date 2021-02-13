package cmds

import (
	"testing"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func TestCd (t *testing.T) {
	i := tcl.InitInterp()
	i.RegisterCommand("cd", Cd, nil)

	_, err := i.Eval("cd ..")
	if err != nil {
		t.Error(err)
	}
}
