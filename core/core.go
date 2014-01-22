package core

import (
	"github.com/Miniand/venditio/inject"
)

type Venditio struct {
	inject.Injector
}

func New() *Venditio {
	return &Venditio{inject.New()}
}
