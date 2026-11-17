package proofgen

import (
	"slices"
	"testing"

	"golang.org/x/tools/go/packages"

	"github.com/mit-pdos/perennial/goose/declfilter"
	"github.com/mit-pdos/perennial/goose/util"
)

// TestDepsThroughNamedType checks that a type whose field has a named type
// with no generated declarations (`type NewTypeName UnderlyingType`) depends
// on the type behind it, so that leanChunks imports the chunk with
// UnderlyingType's instances.
func TestDepsThroughNamedType(t *testing.T) {
	pkgs, err := packages.Load(util.NewPackageConfig(".", false, nil), "./testdata/aliasdeps")
	if err != nil {
		t.Fatal(err)
	}
	if len(pkgs) != 1 || len(pkgs[0].Errors) > 0 {
		t.Fatalf("could not load testdata/aliasdeps: %v", pkgs[0].Errors)
	}
	deps := make(map[string][]string)
	for _, ty := range translateTypesDeps(pkgs[0], declfilter.New(declfilter.FilterConfig{})) {
		deps[ty.name] = ty.deps
	}
	if _, ok := deps["NewTypeName"]; ok {
		t.Fatalf("NewTypeName has generated declarations; the test no longer exercises a type without them")
	}
	if !slices.Contains(deps["Outer"], "UnderlyingType") {
		t.Errorf("deps of Outer = %v, want UnderlyingType", deps["Outer"])
	}
}
