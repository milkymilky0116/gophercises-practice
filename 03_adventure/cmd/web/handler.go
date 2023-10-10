package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	intro, err := app.story.Get("intro")
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	data := newTemplateData(r)
	data.Story = intro
	app.render(w, "view.tmpl.html", data, http.StatusOK)
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	arc := r.URL.Query().Get("arc")
	page, err := app.story.Get(arc)

	if err != nil {
		app.errorLog.Println(err)
		return
	}

	data := newTemplateData(r)
	data.Story = page
	app.render(w, "view.tmpl.html", data, http.StatusOK)
}
