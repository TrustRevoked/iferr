package main

import (
	"go/ast"
	"testing"
)

func assertContainsError(t *testing.T, fieldList *ast.FieldList, expectedOutput bool) {
	actualOutput := containsError(fieldList)
	if actualOutput != expectedOutput {
		t.Fatalf(`Expected identifier "%#v" to yield output: %v (got %v instead)`, fieldList, expectedOutput, actualOutput)
	}
}

func TestContainsErrorNil(t *testing.T) {
	var fieldList *ast.FieldList = nil
	assertContainsError(t, fieldList, false)
}

func TestContainsErrorEmpty(t *testing.T) {
	fieldList := &ast.FieldList{}
	assertContainsError(t, fieldList, false)
}

func TestContainsErrorString(t *testing.T) {
	fieldList := &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Type: &ast.Ident{
					Name: "string",
				},
			},
		},
	}
	assertContainsError(t, fieldList, false)
}

func TestContainsErrorError(t *testing.T) {
	fieldList := &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Type: &ast.Ident{
					Name: "error",
				},
			},
		},
	}
	assertContainsError(t, fieldList, true)
}
