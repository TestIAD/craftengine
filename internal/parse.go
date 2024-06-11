package internal

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

const (
	model   = "./template/model.tmpl"
	pgsql   = "./template/pgsql.tmpl"
	restful = "./template/restful.tmpl"
	router  = "./template/router.tmpl"
	read    = "./template/service/read.tmpl"
	write   = "./template/service/write.tmpl"
	service = "./template/service/service.tmpl"
)

// Paths of template files
var templatePaths = []string{model, pgsql, restful, router, read, write, service}

// Output paths correspond to each template file
var outPath = map[string]string{
	model:   "/internal/models/%s",
	pgsql:   "/internal/models/pgsql/%s",
	restful: "/api/restful/%s",
	router:  "/api/router/%s",
	read:    "/internal/service/%s/read.go",
	write:   "/internal/service/%s/write.go",
	service: "/internal/service/%s/service.go",
}

func pascalToCamel(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func camelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// Service holds template data
type Service struct {
	Lower   string
	Capital string
	Module  string
	Snake   string
}

// Parse generates code files based on provided templates
func Parse(module, service, srcPath string) {
	setService(service, srcPath)
	fileName := fmt.Sprintf("%s.go", camelToSnake(service))
	services := Service{
		Lower:   pascalToCamel(service),
		Capital: service,
		Module:  module,
		Snake:   camelToSnake(service),
	}
	path := fmt.Sprintf("%s/internal/service/%s", srcPath, strings.ToLower(service))
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	} else {
		log.Printf("Nested folders created: %s", path)
	}

	for _, tmplPath := range templatePaths {
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			log.Printf("error parsing template file '%s': %s", tmplPath, err)
			continue
		}

		outFilePath := srcPath + outPath[tmplPath]
		if tmplPath == model || tmplPath == pgsql || tmplPath == restful || tmplPath == router {
			outFilePath = fmt.Sprintf(outFilePath, fileName)
		} else {
			outFilePath = fmt.Sprintf(outFilePath, strings.ToLower(service))
		}
		outputFile, err := os.Create(outFilePath)
		if err != nil {
			log.Printf("error creating output file '%s': %s", outFilePath, err)
			continue
		}
		defer outputFile.Close()

		err = tmpl.Execute(outputFile, services)
		if err != nil {
			log.Printf("error executing template for file '%s': %s", outFilePath, err)
		}
		log.Printf("file generated: '%s'", outFilePath)
	}
}

func setService(service, srcPath string) {
	newFieldName := service + "Service"
	newFieldType := strings.ToLower(service) + ".Service"

	filename := srcPath + "/internal/service/service.go"
	src, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(-1)
	}

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if ts.Name.Name == "Services" {
			structType, ok := ts.Type.(*ast.StructType)
			if !ok {
				return false
			}
			structType.Fields.List = append(structType.Fields.List, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent(newFieldName)},
				Type:  ast.NewIdent(newFieldType),
			})
		}

		return false
	})

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		log.Printf("add servcie struct err")
		os.Exit(-1)
	}
	log.Printf("success mkdir content success %s", filename)
}
