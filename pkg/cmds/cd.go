package cmds

import (
	"os"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func Cd(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) == 1 {
		var home string
		home = os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
			if home == "" {
				home = "~"
			}
		}
		err := os.Chdir(home)
		if err != nil {
			return "", err
		}
		return "", nil
	}
	if len(argv) == 2 {
		err := os.Chdir(argv[1])
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", ArityErr(i, "cd", argv)
}
