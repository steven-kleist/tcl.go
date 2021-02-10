package cmds

import (
	"fmt"
	"strings"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Proc defines the `proc` command for tcl
func Proc(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 4 {
		return "", ArityErr(i, argv[0], argv)
	}
	return "", i.RegisterCommand(argv[1], CallProc, []string{argv[2], argv[3]})
}

// CallProc is the handler user defined tcl `proc`edures.
func CallProc(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	var x []string

	if pd, ok := pd.([]string); ok {
		x = pd
	} else {
		return "", nil
	}

	i.CallFrame = &tcl.CallFrame{Vars: make(map[string]tcl.Var), Parent: i.CallFrame}
	defer func() { i.CallFrame = i.CallFrame.Parent }() // remove the called proc callframe

	arity := 0
	for _, arg := range strings.Split(x[0], " ") {
		if len(arg) == 0 {
			continue
		}
		arity++
		i.SetVar(arg, argv[arity])
	}

	if arity != len(argv)-1 {
		return "", fmt.Errorf("Proc '%s' called with wrong arg num", argv[0])
	}

	body := x[1]
	result, err := i.Eval(body)
	if err == tcl.ErrTclReturn {
		err = nil
	}
	return result, err
}

// RetCodes defines the `break` and `continue` key words.
func RetCodes(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 1 {
		return "", ArityErr(i, argv[0], argv)
	}
	switch argv[0] {
	case "break":
		return "", tcl.ErrTclBreak
	case "continue":
		return "", tcl.ErrTclContinue
	}
	return "", nil
}

// Return defines the `return` key word
func Return(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 1 && len(argv) != 2 {
		return "", ArityErr(i, argv[0], argv)
	}
	var r string
	if len(argv) == 2 {
		r = argv[1]
	}
	return r, tcl.ErrTclReturn
}

// Error defines the `error` key word
func Error(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 1 && len(argv) != 2 {
		return "", ArityErr(i, argv[0], argv)
	}
	return "", fmt.Errorf(argv[1])
}
