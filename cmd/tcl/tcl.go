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
	Exec string `arg:"-e" help:"Script to evaluate before file processing."`
}

func (args) Version() string {
	return fmt.Sprintf("%s v.%s\n", AppName, AppVersion)
}

func main() {
	var args args
	arg.MustParse(&args)

	tcl := tcl.InitInterp()
	registerCoreCommands(tcl)
	tcl.RegisterCommand("puts", cmds.Puts, nil)

	if args.Input == "" && args.Exec == "" {
		fmt.Println("Please give a path...")
		os.Exit(1)
	}

	if args.Exec != "" {
		_, errr := tcl.Eval(args.Exec)
		if errr != nil {
			fmt.Println("Error:", errr)
			os.Exit(1)
		}
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


func registerCoreCommands(i *tcl.Interp) {
	name := [...]string{"+", "-", "*", "/", ">", ">=", "<", "<=", "==", "!="}
	for _, n := range name {
		i.RegisterCommand(n, cmds.Math, nil)
	}
	i.RegisterCommand("set", cmds.Set, nil)
	i.RegisterCommand("unset", cmds.Unset, nil)
	i.RegisterCommand("if", cmds.If, nil)
	i.RegisterCommand("while", cmds.While, nil)
	i.RegisterCommand("break", cmds.RetCodes, nil)
	i.RegisterCommand("continue", cmds.RetCodes, nil)
	i.RegisterCommand("proc", cmds.Proc, nil)
	i.RegisterCommand("return", cmds.Return, nil)
	i.RegisterCommand("error", cmds.Error, nil)
}