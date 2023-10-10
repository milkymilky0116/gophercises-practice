package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/milkymilky0116/gophercises-practice/03_adventure/internal/models"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	mux           *http.ServeMux
	templateCache map[string]*template.Template
	story         *models.AdventureModel
}

func main() {
	mux := http.NewServeMux()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	stories, err := models.ParseJson("gopher.json")

	if err != nil {
		errorLog.Fatal(err)
	}
	cache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		mux:           mux,
		templateCache: cache,
		story: &models.AdventureModel{
			Stories: stories,
		},
	}
	srv := http.Server{
		Addr:     ":4000",
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Server starting on %s\n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
