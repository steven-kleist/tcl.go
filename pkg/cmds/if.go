package cmds

import (
	"strconv"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)


// If defines the `if` command
func If(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 3 && len(argv) != 5 {
		return "", ArityErr(i, argv[0], argv)
	}

	result, err := i.Eval(argv[1])
	if err != nil {
		return "", err
	}

	if r, _ := strconv.Atoi(result); r != 0 {
		return i.Eval(argv[2])
	} else if len(argv) == 5 {
		return i.Eval(argv[4])
	}

	return result, nil
}
