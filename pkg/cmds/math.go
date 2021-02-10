package cmds

import (
	"fmt"
	"strconv"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Math provides mathematical function for the tcl interpreter
func Math(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 3 {
		return "", ArityErr(i, argv[0], argv)
	}
	a, _ := strconv.Atoi(argv[1])
	b, _ := strconv.Atoi(argv[2])
	var c int
	switch {
	case argv[0] == "+":
		c = a + b
	case argv[0] == "-":
		c = a - b
	case argv[0] == "*":
		c = a * b
	case argv[0] == "/":
		c = a / b
	case argv[0] == ">":
		if a > b {
			c = 1
		}
	case argv[0] == ">=":
		if a >= b {
			c = 1
		}
	case argv[0] == "<":
		if a < b {
			c = 1
		}
	case argv[0] == "<=":
		if a <= b {
			c = 1
		}
	case argv[0] == "==":
		if a == b {
			c = 1
		}
	case argv[0] == "!=":
		if a != b {
			c = 1
		}
	default: // FIXME I hate warnings
		c = 0
	}
	return fmt.Sprintf("%d", c), nil
}