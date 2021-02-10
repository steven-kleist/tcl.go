package cmds

import (
	"fmt"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// ArityErr can be used to tell if the used command has the wrong number of arguments.
func ArityErr(i *tcl.Interp, name string, argv []string) error {
	return fmt.Errorf("Wrong number of args for %s %s", name, argv)
}