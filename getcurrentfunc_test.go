package main

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestGetCurrentFunc(t *testing.T) {
	desiredFunc := &ast.FuncDecl{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				Opening: token.Pos(99),
			},
			Results: &ast.FieldList{
				Closing: token.Pos(101),
			},
		},
	}
	file := &ast.File{
		Decls: []ast.Decl{
			desiredFunc,
		},
	}
	position := token.Pos(100)
	funcDecl := getCurrentFunc(file, position)
	if funcDecl != desiredFunc {
		t.Fatalf("Expectation failed")
	}
}
