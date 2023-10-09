package urlshortner

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type jsonConfig struct {
	filename string
}

func (app *AppConfig) CmdJson(w io.Writer, args []string) error {
	var c jsonConfig
	fs := flag.NewFlagSet("JSON Handler", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.filename, "filename", "test.json", "Set json filename that you want to parse.")
	err := fs.Parse(args)
	if err != nil {
		switch {
		case err.Error() == flag.ErrHelp.Error():
			return flag.ErrHelp
		default:
			return err
		}
	}
	buff, err := OpenFile(c.filename)
	if err != nil {
		return err
	}
	err = app.jsonHandler(buff)
	if err != nil {
		return err
	}
	return nil
}
func (app *AppConfig) jsonHandler(jsonBuff []byte) error {
	var u jsonUrls
	err := json.Unmarshal(jsonBuff, &u)
	if err != nil {
		return err
	}
	fmt.Println(u)
	for _, config := range u.Urls {
		path := config.Path
		url := config.Url
		app.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		})
	}
	return nil
}
