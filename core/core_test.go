package core

import (
	"testing"
)

type testRegisterer struct {
	SomeValue string
}

func (tr *testRegisterer) Register(v *Venditio) {
	v.BindValue("testRegisterer", tr)
}

func TestRegister(t *testing.T) {
	v := New()
	tr := &testRegisterer{SomeValue: "blah"}
	if v.Has("testRegisterer") {
		t.Fatal("There is already a testRegisterer available")
	}
	tr.Register(v)
	if !v.Has("testRegisterer") {
		t.Fatal("testRegisterer did not register")
	}
	_, err := v.With("testRegisterer", func(tr *testRegisterer) {
		if tr.SomeValue != "blah" {
			t.Fatal("SomeValue wasn't 'blah'")
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}
