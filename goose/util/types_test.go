package util

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"slices"
	"testing"
)

// TestTypeGetDependenciesThroughAlias checks that a field declared with a type
// alias depends on the named type the alias stands for, so that goose emits
// that type before the struct.
func TestTypeGetDependenciesThroughAlias(t *testing.T) {
	const src = `package p

type NewTypeName [2]byte

type AliasName = NewTypeName

type Outer struct {
	x AliasName
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("p", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}
	outer := pkg.Scope().Lookup("Outer").Type().Underlying()
	deps := slices.Collect(TypeGetDependencies("p", outer))
	if !slices.Contains(deps, "NewTypeName") {
		t.Errorf("dependencies of Outer = %v, want NewTypeName", deps)
	}
}
