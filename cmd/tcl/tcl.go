package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

const (
	AppName    = "tcl.exe"
	AppVersion = "0.0.1"
)

type args struct {
	Input string `arg:"positional" help:"File to parse/execute"`
	Exec  string `arg:"-e" help:"Script to evaluate before file processing."`
}

func (args) Version() string {
	return fmt.Sprintf("%s v.%s\n", AppName, AppVersion)
}

func main() {
	var args args
	arg.MustParse(&args)

	interpreter := tcl.InitInterp()
	registerCoreCommands(interpreter)
	interpreter.RegisterCommands([]tcl.CmdDef{
		{Name: "puts", Fn: cmds.Puts, Privdata: nil},
		{Name: "append", Fn: cmds.Append, Privdata: nil},
		{Name: "pwd", Fn: cmds.Pwd, Privdata: nil},
	})

	if args.Input == "" && args.Exec == "" {
		fmt.Println("Please give a path...")
		os.Exit(1)
	}

	if args.Exec != "" {
		_, errr := interpreter.Eval(args.Exec)
		if errr != nil {
			fmt.Print("Error:", errr)
			os.Exit(1)
		}
	}

	buf, err := ioutil.ReadFile(args.Input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err := interpreter.Eval(string(buf))
		if err != nil {
			fmt.Println("Error:", err, result)
		} else {
			fmt.Print(result)
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
