package cmds

import (
	"strings"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Append defines the `append` command
func Append(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) >= 3 {
		newVal := append([]string{}, argv[2:]...)

		val, ok := i.Var(argv[1])
		if !ok {
			i.SetVar(argv[1], strings.Join(newVal, ""))
			return strings.Join(newVal, ""), nil
		}
		newVal = append([]string{string(val)}, newVal...)
		i.SetVar(argv[1], strings.Join(newVal, ""))
		return strings.Join(newVal, ""), nil
	}
	return "", ArityErr(i, "append", argv)
}
