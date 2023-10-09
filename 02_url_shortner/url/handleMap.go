package urlshortner

import (
	"net/http"
)

func (app *AppConfig) CmdMap(pathToUrls map[string]string) {
	app.mapHandler(pathToUrls)
}
func (app *AppConfig) mapHandler(pathsToUrls map[string]string) {
	for url, redirect := range pathsToUrls {
		newUrl := url
		newRedirect := redirect
		app.Mux.HandleFunc(newUrl, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, newRedirect, http.StatusPermanentRedirect)
		})
	}
}
