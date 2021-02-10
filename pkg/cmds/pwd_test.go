package cmds

import (
	"os"
	"testing"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func TestPwd(t *testing.T) {
	i := tcl.InitInterp()
	i.RegisterCommand("pwd", Pwd, nil)

	pwd, err := os.Getwd()
	if err != nil {
		t.Fail()
	}

	res, err := i.Eval("pwd")
	if err != nil {
		t.Fail()
	}
	if res != pwd {
		t.Errorf("Expected: %s, got %s", pwd, res)
	}
}
