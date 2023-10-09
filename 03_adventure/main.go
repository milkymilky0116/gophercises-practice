package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/milkymilky0116/gophercises-practice/03_adventure/cmd/cli"
	"github.com/milkymilky0116/gophercises-practice/03_adventure/cmd/web"
	"github.com/milkymilky0116/gophercises-practice/03_adventure/internal/models"
)

func main() {
	stories, err := models.ParseJson("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stories["intro"])

	config, err := cli.ParseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	app := web.Application{
		Filename: config.Filename,
		Router:   httprouter.New(),
		Stories:  stories,
	}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: app.Routes(),
	}

	log.Printf("Server Listening on :%d", config.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
