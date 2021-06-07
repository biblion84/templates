package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tmplFuncs = template.FuncMap{
	"add":        add,
	"JoinString": JoinString,
}

var templates map[string]*template.Template

func templateExample(w http.ResponseWriter) {
	templates["maps.page.tmpl"].Execute(w, nil)
}

func InitTemplates() error {
	templates = make(map[string]*template.Template)

	files, err := filepath.Glob(filepath.Join("./html", "*.page.tmpl"))
	if err != nil {
		return err
	}
	for _, file := range files {
		name := filepath.Base(file)
		file, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		tmpl, err := template.New(name).Funcs(tmplFuncs).Parse(string(file))
		if err != nil {
			return err
		}
		tmpl, err = tmpl.ParseGlob(filepath.Join("./html", "*.layout.tmpl"))
		if err != nil {
			return err
		}
		templates[tmpl.Name()] = tmpl
	}
	return nil

}

func add(a, b, c int) int {
	return a + b + c
}

func JoinString(elements []string) string {
	return strings.ToUpper(strings.Join(elements, "-"))
}
