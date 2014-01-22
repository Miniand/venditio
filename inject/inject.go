package inject

import (
	"errors"
	"fmt"
	"reflect"
)

type Injector interface {
	BindFactory(dep string, factory func() interface{})
	BindValue(dep string, value interface{})
	Get(dep string) (interface{}, bool)
	Has(dep string) bool
	With(args ...interface{}) ([]reflect.Value, error)
	Parent() Injector
	SetParent(parent Injector)
}

func New() Injector {
	return &basic{
		factories: map[string]func() interface{}{},
		deps:      map[string]interface{}{},
	}
}

type basic struct {
	factories map[string]func() interface{}
	deps      map[string]interface{}
	parent    Injector
}

func (b *basic) BindFactory(dep string, factory func() interface{}) {
	b.factories[dep] = factory
}

func (b *basic) BindValue(dep string, value interface{}) {
	b.deps[dep] = value
}

func (b *basic) Get(dep string) (interface{}, bool) {
	var (
		ok      bool
		factory func() interface{}
	)
	if _, ok = b.deps[dep]; !ok {
		if factory, ok = b.factories[dep]; !ok {
			if parent := b.Parent(); parent != nil {
				return parent.Get(dep)
			}
			return nil, false
		}
		b.deps[dep] = factory()
	}
	return b.deps[dep], true
}

func (b *basic) Has(dep string) bool {
	if _, ok := b.deps[dep]; ok {
		return true
	}
	if _, ok := b.factories[dep]; ok {
		return true
	}
	if parent := b.Parent(); parent != nil {
		return parent.Has(dep)
	}
	return false
}

func (b *basic) With(args ...interface{}) ([]reflect.Value, error) {
	if len(args) == 0 {
		return nil, errors.New(
			"Must provide dependency names followed by callback to run")
	}
	depCount := len(args) - 1
	deps := make([]string, depCount)
	for i, dName := range args[:depCount] {
		strDep, ok := dName.(string)
		if !ok {
			return nil, fmt.Errorf("Argument %d is not a string", i)
		}
		deps[i] = strDep
	}
	cb := args[len(args)-1]
	t := reflect.TypeOf(cb)
	numIn := t.NumIn()
	if numIn > len(deps) {
		return nil, errors.New(
			"More arguments in callback than in dependency list")
	}
	var in = make([]reflect.Value, numIn)
	for i := 0; i < numIn; i++ {
		v, ok := b.Get(deps[i])
		if !ok {
			return nil, fmt.Errorf("Could not find dependency %s", deps[i])
		}
		in[i] = reflect.ValueOf(v)
	}
	return reflect.ValueOf(cb).Call(in), nil
}

func (b *basic) Parent() Injector {
	return b.parent
}

func (b *basic) SetParent(parent Injector) {
	b.parent = parent
}
