package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/milkymilky0116/gophercises-practice/03_adventure/internal/models"
)

type TemplateData struct {
	Story *models.AdventureStory
}

func indentation(line string) string {
	return "\t" + line
}

var functions = template.FuncMap{
	"indentation": indentation,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.gohtml")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.gohtml")

		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.gohtml")

		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

func newTemplateData(r *http.Request) *TemplateData {
	return &TemplateData{}
}
