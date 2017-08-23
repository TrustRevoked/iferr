package main

import (
	"go/ast"
	"testing"
)

func assertZeroValue(t *testing.T, expr ast.Expr, expectedOutput string) {
	actualOutput := zeroValue(expr)
	if actualOutput != expectedOutput {
		t.Fatalf(`Expected identifier "%#v" to yield output: %s (got %s instead)`, expr, expectedOutput, actualOutput)
	}
}

func assertIdentZeroValue(t *testing.T, ident, expectedOutput string) {
	_ident := ast.NewIdent(ident)
	assertZeroValue(t, _ident, expectedOutput)
}

func TestZeroValueString(t *testing.T) {
	assertIdentZeroValue(t, "string", `""`)
}

func TestZeroValueIntTypes(t *testing.T) {
	assertIdentZeroValue(t, "int", "0")
	assertIdentZeroValue(t, "int64", "0")
	assertIdentZeroValue(t, "int32", "0")
}

func TestZeroValueFloatTypes(t *testing.T) {
	assertIdentZeroValue(t, "float", "0")
	assertIdentZeroValue(t, "float64", "0")
	assertIdentZeroValue(t, "float32", "0")
}

func TestZeroValueBool(t *testing.T) {
	assertIdentZeroValue(t, "bool", "false")
}

func TestZeroValueStructName(t *testing.T) {
	assertIdentZeroValue(t, "Ident", "Ident{}")
}

func TestZeroValueError(t *testing.T) {
	assertIdentZeroValue(t, "error", "err")
}

func TestZeroValueStar(t *testing.T) {
	ident := ast.NewIdent("Ident")
	star := &ast.StarExpr{
		X: ident,
	}

	assertZeroValue(t, star, "nil")
}

func TestZeroValueSelector(t *testing.T) {
	_ast := ast.NewIdent("ast")
	_ident := ast.NewIdent("Ident")
	sel := &ast.SelectorExpr{
		X:   _ast,
		Sel: _ident,
	}

	assertZeroValue(t, sel, "ast.Ident{}")
}

func TestZeroValueStarSelector(t *testing.T) {
	_ast := ast.NewIdent("ast")
	_ident := ast.NewIdent("Ident")
	sel := &ast.SelectorExpr{
		X:   _ast,
		Sel: _ident,
	}
	star := &ast.StarExpr{
		X: sel,
	}

	assertZeroValue(t, star, "nil")
}
