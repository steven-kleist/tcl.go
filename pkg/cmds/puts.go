package cmds

import (
	"fmt"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// CommandPuts implements the `puts` command for tcl.
func CommandPuts(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", fmt.Errorf("Wrong number of arguments for %s %s", argv[0], argv)
	}

	fmt.Println(argv[1])
	return "", nil
}
