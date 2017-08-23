package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: iferr file.go POSITION")
	}
	filePath := os.Args[1]
	_position, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	position := token.Pos(_position)

	fset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	currentFunc := getCurrentFunc(parsedFile, position)

	code := codeGen(currentFunc.Type.Results)

	fmt.Println("if err != nil {")
	fmt.Println(code)
	fmt.Println("}")
}

func getCurrentFunc(file *ast.File, position token.Pos) *ast.FuncDecl {
	for _, decl := range file.Decls {
		if decl.Pos() < position && decl.End() > position {
			return decl.(*ast.FuncDecl)
		}
	}
	return nil
}

func codeGen(fieldList *ast.FieldList) string {
	if !containsError(fieldList) {
		return "log.Fatal(err)"
	}

	args := []string{}

	for _, field := range fieldList.List {
		args = append(args, zeroValue(field.Type))
	}

	return "return " + strings.Join(args, ", ")
}

func containsError(fieldList *ast.FieldList) bool {
	if fieldList == nil {
		return false
	}
	for _, field := range fieldList.List {
		if field.Type.(*ast.Ident).Name == "error" {
			return true
		}
	}
	return false
}

func zeroValue(expr ast.Expr) string {
	// Simple identifier
	// i.e. string, bool, Ident
	if ident, ok := expr.(*ast.Ident); ok {
		switch ident.Name {
		case "string":
			return `""`
		case "bool":
			return "false"
		case "int", "int64", "int32", "float", "float64", "float32":
			return "0"
		case "error":
			return "err"
		default:
			return ident.Name + "{}"
		}
	}
	// Selector expression
	// i.e. ast.Ident, time.Time
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		x := selector.X.(*ast.Ident)
		sel := selector.Sel
		return x.Name + "." + sel.Name + "{}"
	}
	// Everything else
	// i.e. *ast.Ident
	return "nil"
}
