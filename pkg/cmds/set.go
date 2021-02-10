package cmds

import (
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Set defines the `set` command
func Set(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 3 {
		return "", ArityErr(i, argv[0], argv)
	}
	i.SetVar(argv[1], argv[2])
	return argv[2], nil
}

// Unset defines the `unset` command
func Unset(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", ArityErr(i, argv[0], argv)
	}
	i.UnsetVar(argv[1])
	return "", nil
}