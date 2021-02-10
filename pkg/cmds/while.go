package cmds

import (
	"strconv"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// While defines the `while` command
func While(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 3 {
		return "", ArityErr(i, argv[0], argv)
	}

	for {
		result, err := i.Eval(argv[1])
		if err != nil {
			return "", err
		}
		if r, _ := strconv.Atoi(result); r != 0 {
			result, err := i.Eval(argv[2])
			switch err {
			case tcl.ErrTclContinue, nil:
				//pass
			case tcl.ErrTclBreak:
				return result, nil
			default:
				return result, err
			}
		} else {
			return result, nil
		}
	}
}
