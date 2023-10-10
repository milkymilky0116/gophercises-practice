package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, page string, data *TemplateData, status int) {
	ts, ok := app.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exists", page)
		app.errorLog.Println(err)
		return
	}

	buff := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buff, "base", data)

	if err != nil {
		app.errorLog.Println(err)
	}
	w.WriteHeader(status)

	buff.WriteTo(w)
}
