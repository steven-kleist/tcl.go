package tcl

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	TCL_RETURN   = errors.New("RETURN")
	TCL_BREAK    = errors.New("BREAK")
	TCL_CONTINUE = errors.New("CONTINUE")
)

type Var string
type CmdFunc func(i *Interp, argv []string, privdata interface{}) (string, error)
type Cmd struct {
	fn       CmdFunc
	privdata interface{}
}
type CallFrame struct {
	vars   map[string]Var
	parent *CallFrame
}
type Interp struct {
	level     int
	callframe *CallFrame
	commands  map[string]Cmd
}

func InitInterp() *Interp {
	return &Interp{
		level:     0,
		callframe: &CallFrame{vars: make(map[string]Var)},
		commands:  make(map[string]Cmd),
	}
}

func (i *Interp) Var(name string) (Var, bool) {
	for frame := i.callframe; frame != nil; frame = frame.parent {
		v, ok := frame.vars[name]
		if ok {
			return v, ok
		}
	}
	return "", false
}
func (i *Interp) SetVar(name, val string) {
	log.Printf("[(*Interp).SetVar] name = %s, val = %s", name, val)
	i.callframe.vars[name] = Var(val)
}

func (i *Interp) UnsetVar(name string) {
	log.Printf("[(*Interp).UnsetVar] name = %s", name)
	delete(i.callframe.vars, name)
}

func (i *Interp) Command(name string) *Cmd {
	v, ok := i.commands[name]
	if !ok {
		return nil
	}
	return &v
}

func (i *Interp) RegisterCommand(name string, fn CmdFunc, privdata interface{}) error {
	c := i.Command(name)
	if c != nil {
		return fmt.Errorf("Command '%s' already defined", name)
	}

	i.commands[name] = Cmd{fn, privdata}
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
			log.Printf("[(*Interp).Eval] p.Type == PT_VAR: t = %s", t)
			v, ok := i.Var(t)
			if !ok {
				return "", fmt.Errorf("No such variable '%s'", t)
			}
			t = string(v)
		case PT_CMD:
			log.Printf("[(*Interp).Eval] p.Type == PT_CMD: t = %s", t)
			result, err = i.Eval(t)
			if err != nil {
				return result, err
			}
			t = result

		case PT_ESC:
			log.Printf("[(*Interp).Eval] p.Type == PT_ESC: t = %s", t)
			// XXX: escape handling missing!
		case PT_SEP:
			log.Printf("[(*Interp).Eval] p.Type == PT_SEP: t = %s", t)
			prevtype = p.Type
			continue
		}

		// We have a complete command + args. Call it!
		if p.Type == PT_EOL {
			prevtype = p.Type
			if len(argv) != 0 {
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
