package urlshortner

import (
	"fmt"
	"net/http"

	"go.etcd.io/bbolt"
)

type AppConfig struct {
	Mux *http.ServeMux
	DB  *bbolt.DB
}

func (app *AppConfig) Run() *http.ServeMux {
	fmt.Println("Server starting on :4000")
	http.ListenAndServe(":4000", app.Mux)
	return app.Mux
}
