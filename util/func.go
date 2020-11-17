package util

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// Get the name and path of a func
func FuncPathAndName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// Get the name of a func (with package path)
func FuncName(f interface{}) string {
	splitFuncName := strings.Split(FuncPathAndName(f), ".")
	return splitFuncName[len(splitFuncName)-1]
}

func SimpleFuncDescription(f interface{}) string {
	desc := FuncDescription(f)
	tmp := strings.Split(desc, "\n")
	return strings.Trim(tmp[0], " \n")
}

// Get description of a func
func FuncDescription(f interface{}) string {
	fileName, _ := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).FileLine(0)
	funcName := FuncName(f)
	fset := token.NewFileSet()

	// Parse src
	parsedAst, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst

	importPath, _ := filepath.Abs("/")
	myDoc := doc.New(pkg, importPath, doc.AllDecls)
	for _, theFunc := range myDoc.Funcs {
		if theFunc.Name == funcName {
			return strings.TrimRight(theFunc.Doc, "\n")
		}
	}
	return ""
}

func SliceContains(slice []interface{}, data interface{}) bool {
	for _, v := range slice {
		if v == data {
			return true
		}
	}
	return false
}
