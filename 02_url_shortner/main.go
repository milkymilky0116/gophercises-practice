package main

import (
	"log"
	"net/http"
	"os"

	"github.com/milkymilky0116/gophercises-practice/02_url_shortner/cli"
	"github.com/milkymilky0116/gophercises-practice/02_url_shortner/db"
	urlshortner "github.com/milkymilky0116/gophercises-practice/02_url_shortner/url"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	app := urlshortner.AppConfig{
		Mux: http.NewServeMux(),
		DB:  db,
	}
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	config, err := cli.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	app.MapHandler(pathsToUrls)

	ymlfileBuff, err := cli.OpenFile(config.Ymlfile)
	if err != nil {
		log.Fatal(err)
	}
	err = app.YAMLHandler(ymlfileBuff)
	if err != nil {
		log.Fatal(err)
	}

	jsonBuff, err := cli.OpenFile(config.Jsonfile)
	if err != nil {
		log.Fatal(err)
	}
	err = app.JsonHandler(jsonBuff)
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
