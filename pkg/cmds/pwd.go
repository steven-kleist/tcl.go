package cmds

import (
	"os"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Pwd defines the `pwd` command
func Pwd(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) > 1 {
		return "", ArityErr(i, "pwd", argv)
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}
