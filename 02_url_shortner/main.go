package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	urlshortner "github.com/milkymilky0116/gophercises-practice/02_url_shortner/url"
)

func main() {
	app := urlshortner.AppConfig{
		Mux: http.NewServeMux(),
	}
	err := handleCmd(os.Stdout, os.Args[1:], app)
	if err != nil {
		os.Exit(1)
	}
	app.Run()

}
func printUsage(w io.Writer, command string, app urlshortner.AppConfig) {
	fmt.Fprintf(w, "Usage: URL Shortner [map|yaml|json] -h\n")
	switch command {
	case "map":
		fmt.Println("Print Usage of Map Handler")
	case "yaml":
		app.CmdYaml(w, []string{"-h"})
	case "json":
		app.CmdJson(w, []string{"-h"})
	case "db":
		app.CmdDB(w, []string{"-h"})
	default:
		fmt.Println("Print Usage of Map Handler")
		app.CmdYaml(w, []string{"-h"})
		app.CmdJson(w, []string{"-h"})
		app.CmdDB(w, []string{"-h"})
	}
}
func handleCmd(w io.Writer, args []string, app urlshortner.AppConfig) error {
	var err error
	var cmdType string
	if len(args) < 1 {
		err = urlshortner.ErrInvalidNumofArgs
	} else {
		switch args[0] {
		case "map":
			pathsToUrls := map[string]string{
				"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
				"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
			}
			app.CmdMap(pathsToUrls)
			cmdType = "map"
		case "yaml":
			err = app.CmdYaml(w, args[1:])
			cmdType = "yaml"
		case "json":
			err = app.CmdJson(w, args[1:])
			cmdType = "json"
		case "db":
			err = app.CmdDB(w, args[1:])
			cmdType = "db"
		case "-h":
			printUsage(w, "", app)
		case "-help":
			printUsage(w, "", app)
		default:
			err = urlshortner.ErrInvalidSubCmd
		}
	}

	if err != nil {
		fmt.Fprintln(w, err)
		printUsage(w, cmdType, app)
	}

	return err
}
