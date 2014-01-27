package cmd

import (
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	c := New()
	res := ""
	c.Register("blah", func(args []string) error {
		res = strings.Join(args, ",")
		return nil
	})
	err := c.Handle([]string{"blah", "smeg", "glorb"})
	if err != nil {
		t.Fatal(err)
	}
	if res != "smeg,glorb" {
		t.Errorf("Got %s", res)
	}
}

func TestSubHandler(t *testing.T) {
	subC := New()
	gotHere := false
	subC.Register("plah", func(args []string) error {
		gotHere = true
		return nil
	})
	c := New()
	c.Register("blah", subC.Handle)
	err := c.Handle([]string{"blah", "plah"})
	if err != nil {
		t.Fatal(err)
	}
	if !gotHere {
		t.Error("Did not get into handler")
	}
}
