package main

import (
	"go/ast"
	"testing"
)

func assertCodeGen(t *testing.T, fieldList *ast.FieldList, expectedOutput string) {
	actualOutput := codeGen(fieldList)
	if actualOutput != expectedOutput {
		t.Fatalf(`Expected identifier "%#v" to yield output: %v (got %v instead)`, fieldList, expectedOutput, actualOutput)
	}
}

func TestCodeGenNil(t *testing.T) {
	var fieldList *ast.FieldList = nil
	assertCodeGen(t, fieldList, "log.Fatal(err)")
}

func TestCodeGenErrorEmpty(t *testing.T) {
	fieldList := &ast.FieldList{}
	assertCodeGen(t, fieldList, "log.Fatal(err)")
}

func TestCodeGenString(t *testing.T) {
	fieldList := &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Type: &ast.Ident{
					Name: "string",
				},
			},
		},
	}
	assertCodeGen(t, fieldList, "log.Fatal(err)")
}

func TestCodeGenError(t *testing.T) {
	fieldList := &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Type: &ast.Ident{
					Name: "error",
				},
			},
		},
	}
	assertCodeGen(t, fieldList, "return err")
}

func TestCodeGenStringError(t *testing.T) {
	fieldList := &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Type: &ast.Ident{
					Name: "string",
				},
			},
			&ast.Field{
				Type: &ast.Ident{
					Name: "error",
				},
			},
		},
	}
	assertCodeGen(t, fieldList, `return "", err`)
}
