package asset

import (
	"testing"
)

func TestResolve(t *testing.T) {
	// Try to resolve this file
	a := New()
	a.AddPackagePath("github.com/Miniand/venditio/asset")
	r := a.Resolve("asset_test.go")
	if len(r) != 1 {
		t.Error("Could not find one result for asset_test.go")
	}
}
