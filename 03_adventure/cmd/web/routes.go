package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	fileServer := http.FileServer(http.Dir("./ui/static"))
	app.mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	app.mux.HandleFunc("/", app.home)
	app.mux.HandleFunc("/view", app.view)
	return app.mux
}
