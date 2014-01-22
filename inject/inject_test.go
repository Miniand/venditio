package inject

import (
	"testing"
)

func TestNew(t *testing.T) {
	New()
}

func TestBindFactory(t *testing.T) {
	i := New()
	i.BindFactory("testDep", func() interface{} {
		return "blah"
	})
	d, ok := i.Get("testDep")
	if !ok {
		t.Fatal("Could not find testDep")
	}
	if d != "blah" {
		t.Fatalf("Expected blah but got %s", d)
	}
}

func TestBindValue(t *testing.T) {
	i := New()
	i.BindValue("testDep", "blah")
	d, ok := i.Get("testDep")
	if !ok {
		t.Fatal("Could not find testDep")
	}
	if d != "blah" {
		t.Fatalf("Expected blah but got %s", d)
	}
}

func TestHas(t *testing.T) {
	i := New()
	if i.Has("testDep") {
		t.Error("Expected not to have testDep")
	}
	i.BindValue("testDepValue", "blah")
	if !i.Has("testDepValue") {
		t.Error("Expected to have testDepValue")
	}
	i.BindFactory("testDepFactory", func() interface{} {
		return "blah"
	})
	if !i.Has("testDepFactory") {
		t.Error("Expected to have testDepFactory")
	}
}

func TestWith(t *testing.T) {
	i := New()
	i.BindFactory("testDep", func() interface{} {
		return "Magooba!"
	})
	hasRun := false
	_, err := i.With("testDep", func(testDep string) {
		if testDep != "Magooba!" {
			t.Errorf("Expected Magooba! but got %s", testDep)
		}
		hasRun = true
	})
	if err != nil {
		t.Error(err)
	}
	if !hasRun {
		t.Error("Callback did not run")
	}
}

func TestParent(t *testing.T) {
	parent := New()
	i := New()
	i.SetParent(parent)
	parent.BindValue("parentVal1", 1)
	parent.BindValue("parentVal2", 2)
	i.BindValue("parentVal2", 3) // Override
	i.BindValue("childVal", 4)
	_, err := i.With("parentVal1", "parentVal2", "childVal",
		func(parentVal1, parentVal2, childVal int) {
			if parentVal1 != 1 {
				t.Errorf("Expected parentVal1 to be 1 but got %d", parentVal1)
			}
			if parentVal2 != 3 {
				t.Errorf("Expected parentVal2 to be 3 but got %d", parentVal2)
			}
			if childVal != 4 {
				t.Errorf("Expected childVal to be 4 but got %d", childVal)
			}
		})
	if err != nil {
		t.Error(err)
	}
	_, err = parent.With("parentVal1", "parentVal2",
		func(parentVal1, parentVal2 int) {
			if parentVal1 != 1 {
				t.Errorf("Expected parentVal1 to be 1 but got %d", parentVal1)
			}
			if parentVal2 != 2 {
				t.Errorf("Expected parentVal2 to be 2 but got %d", parentVal2)
			}
		})
	if err != nil {
		t.Error(err)
	}
}
