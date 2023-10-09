package web

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Home")
}
func (app *Application) Routes() *httprouter.Router {
	app.Router.GET("/", Home)
	return app.Router
}
