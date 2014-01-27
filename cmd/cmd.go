package cmd

import (
	"errors"
	"fmt"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/inject"
	"os"
)

const (
	DEP_COMMANDER = "cmdCommander"
)

type Handler func(args []string) error

type Commander interface {
	Register(command string, handler Handler)
	Handle(args []string) error
}

func Register(v *core.Venditio) {
	v.BindFactory(DEP_COMMANDER, func(i inject.Injector) interface{} {
		return New()
	})
}

func New() *Cmdr {
	return &Cmdr{
		Handlers: map[string]Handler{},
	}
}

func Run(c Commander) {
	err := c.Handle(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

type Cmdr struct {
	Handlers map[string]Handler
}

func (c *Cmdr) Register(command string, handler Handler) {
	c.Handlers[command] = handler
}

func (c *Cmdr) Handle(args []string) error {
	if len(args) == 0 {
		return errors.New("A command must be provided")
	}
	handler := c.Handlers[args[0]]
	if handler == nil {
		return fmt.Errorf("Command not found: %s", args[0])
	}
	return handler(args[1:])
}
