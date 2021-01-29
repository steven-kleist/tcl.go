package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/stevenkl/tcl.go/pkg/tcl"
	"github.com/stevenkl/tcl.go/pkg/cmds"
)

const (
	AppName    = "tcl.exe"
	AppVersion = "0.0.1"
)

type args struct {
	Input string `arg:"positional" help:"File to parse/execute"`
}

func (args) Version() string {
	return fmt.Sprintf("%s v.%s", AppName, AppVersion)
}

func main() {
	var args args
	arg.MustParse(&args)

	tcl := tcl.InitInterp()
	tcl.RegisterCoreCommands()
	tcl.RegisterCommand("puts", cmds.CommandPuts, nil)

	if args.Input == "" {
		fmt.Println("Please give a path...")
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(args.Input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err := tcl.Eval(string(buf))
		if err != nil {
			fmt.Println("Error:", err, result)
		} else {
			fmt.Println(result)
		}
	}
	
}
