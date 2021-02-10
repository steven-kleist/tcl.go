package tcl

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrTclReturn symbols a default tcl return
	ErrTclReturn   = errors.New("RETURN")
	// ErrTclBreak symbols a tcl break
	ErrTclBreak    = errors.New("BREAK")
	// ErrTclContinue symbols a tcl continue
	ErrTclContinue = errors.New("CONTINUE")
)

// Var represents a tcl variable
type Var string

// CmdFunc defines the signature of a tcl command used in tcl.go
type CmdFunc func(i *Interp, argv []string, privdata interface{}) (string, error)

// Cmd defines the space used for a command
type Cmd struct {
	fn       CmdFunc
	privdata interface{}
}
// CallFrame defines the current tcl environment
type CallFrame struct {
	Vars   map[string]Var
	Parent *CallFrame
}
// Interp represents the tcl interpreter
type Interp struct {
	Level     int
	CallFrame *CallFrame
	Commands  map[string]Cmd
}

// InitInterp creates a new interpreter
func InitInterp() *Interp {
	return &Interp{
		Level:     0,
		CallFrame: &CallFrame{Vars: make(map[string]Var)},
		Commands:  make(map[string]Cmd),
	}
}

// Var returns the value of a variable in the current interpreter
func (i *Interp) Var(name string) (Var, bool) {
	for frame := i.CallFrame; frame != nil; frame = frame.Parent {
		v, ok := frame.Vars[name]
		if ok {
			return v, ok
		}
	}
	return "", false
}

// SetVar sets the value of a variable in the current interpreter
func (i *Interp) SetVar(name, val string) {
	// log.Printf("[(*Interp).SetVar] name = %s, val = %s", name, val)
	i.CallFrame.Vars[name] = Var(val)
}

// UnsetVar removes a variable from the current interpreter
func (i *Interp) UnsetVar(name string) {
	// log.Printf("[(*Interp).UnsetVar] name = %s", name)
	delete(i.CallFrame.Vars, name)
}

// Command returns the `tcl.Cmd` with the given name
func (i *Interp) Command(name string) *Cmd {
	v, ok := i.Commands[name]
	if !ok {
		return nil
	}
	return &v
}

// RegisterCommand adds a new `tcl.Cmd` to the current interpreter
func (i *Interp) RegisterCommand(name string, fn CmdFunc, privdata interface{}) error {
	c := i.Command(name)
	if c != nil {
		return fmt.Errorf("Command '%s' already defined", name)
	}

	i.Commands[name] = Cmd{fn, privdata}
	return nil
}

// CmdDef represents a tcl command for registration in tcl.
type CmdDef struct {
	Name string
	Fn CmdFunc
	Privdata interface{}
}

// RegisterCommands registers a list of commands
func (i *Interp) RegisterCommands(cmds []CmdDef) error {
	var err error
	for _, cmd := range cmds {
		err = i.RegisterCommand(cmd.Name, cmd.Fn, cmd.Privdata)
		if err != nil {
			return err
		}
	}
	return nil
}

// Eval evaluates the given string
func (i *Interp) Eval(t string) (string, error) {
	p := InitParser(t)
	var result string
	var err error

	argv := []string{}

	for {
		prevtype := p.Type
		// XXX
		t = p.GetToken()
		if p.Type == PT_EOF {
			break
		}

		switch p.Type {
		case PT_VAR:
			// log.Printf("[(*Interp).Eval] p.Type == PT_VAR: t = %s", t)
			v, ok := i.Var(t)
			if !ok {
				return "", fmt.Errorf("No such variable '%s'", t)
			}
			t = string(v)
		case PT_CMD:
			// log.Printf("[(*Interp).Eval] p.Type == PT_CMD: t = %s", t)
			result, err = i.Eval(t)
			if err != nil {
				return result, err
			}
			t = result

		case PT_ESC:
			// log.Printf("[(*Interp).Eval] p.Type == PT_ESC: t = %s", t)
			// XXX: escape handling missing!
		case PT_SEP:
			// log.Printf("[(*Interp).Eval] p.Type == PT_SEP: t = %s", t)
			prevtype = p.Type
			continue
		}

		// We have a complete command + args. Call it!
		if p.Type == PT_EOL {
			prevtype = p.Type
			if len(argv) != 0 {
				// log.Printf("[(*Interp).Eval] command: '%s'", argv[0])
				c := i.Command(argv[0])
				if c == nil {
					return "", fmt.Errorf("No such command '%s'", argv[0])
				}
				result, err = c.fn(i, argv, c.privdata)
				if err != nil {
					return result, err
				}
			}
			// Prepare for the next command
			argv = []string{}
			continue
		}

		// We have a new token, append to the previous or as new arg?
		if prevtype == PT_SEP || prevtype == PT_EOL {
			argv = append(argv, t)
		} else { // Interpolation
			argv[len(argv)-1] = strings.Join([]string{argv[len(argv)-1], t}, "")
		}
		prevtype = p.Type
	}
	return result, nil
}
