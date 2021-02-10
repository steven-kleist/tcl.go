package cmds

import (
	"fmt"
	"strings"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Puts implements the `puts` command for tcl.
func Puts(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) > 3 {
		return "", fmt.Errorf("Wrong number of arguments (%d) for %s %s", len(argv), argv[0], argv)
	}
	
	if strings.ToLower(argv[1]) == "-nonewline" {
		fmt.Print(argv[2])
		return "", nil
	}

	fmt.Println(argv[1])
	return "", nil
}
