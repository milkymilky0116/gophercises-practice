package main

import (
	"log"
	"net/http"
	"os"

	"github.com/milkymilky0116/gophercises-practice/02_url_shortner/cli"
	urlshortner "github.com/milkymilky0116/gophercises-practice/02_url_shortner/url"
)

func main() {
	app := urlshortner.AppConfig{
		Mux: http.NewServeMux(),
	}
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	// 	yaml := `- path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `
	app.MapHandler(pathsToUrls)
	config, err := cli.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
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
