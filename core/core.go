package core

import (
	"github.com/codegangsta/martini"
	"reflect"
)

type Venditio struct {
	*martini.ClassicMartini
}

func New() *Venditio {
	return &Venditio{martini.Classic()}
}

func (v *Venditio) Has(i interface{}) bool {
	return v.Injector.Get(reflect.TypeOf(i)).IsValid()
}
