package urlshortner

import (
	"flag"
	"io"
	"net/http"

	"gopkg.in/yaml.v3"
)

type yamlConfig struct {
	filename string
}

func (app *AppConfig) CmdYaml(w io.Writer, args []string) error {
	var c yamlConfig
	fs := flag.NewFlagSet("YAML Handler", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.filename, "filename", "test.yaml", "Set yaml filename that you want to parse.")
	err := fs.Parse(args)
	if err != nil {
		switch {
		case err.Error() == flag.ErrHelp.Error():
			return flag.ErrHelp
		default:
			return err
		}
		//return err
	}
	buff, err := OpenFile(c.filename)
	if err != nil {
		return err
	}
	err = app.yamlHandler(buff)
	if err != nil {
		return err
	}
	return nil
}
func (app *AppConfig) yamlHandler(yml []byte) error {
	var u urls
	err := yaml.Unmarshal(yml, &u)
	if err != nil {
		return err
	}
	for _, config := range u {
		path := config.Path
		url := config.Url
		app.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		})
	}
	return nil
}
