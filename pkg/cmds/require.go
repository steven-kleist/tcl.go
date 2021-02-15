package cmds

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Require implements the `include` command
func Require(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", ArityErr(i, "require", argv)
	}

	file, err := filepath.Abs(argv[1])
	if err != nil {
		return "", err
	}

	registry := pd.(map[string]bool)

	if registry[file] != false {
		return "", nil
	}

	if _, stat := os.Stat(file); os.IsNotExist(stat) {
		return "", fmt.Errorf("file  '%s' not found", file)
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	var res string
	ret, err2 := i.Eval(string(buf))
	if err2 != nil {
		if err2 == tcl.ErrTclReturn {
			res = ret
		} else {
			return "", err2
		}
		
	}
	registry[file] = true

	return res, nil
}
