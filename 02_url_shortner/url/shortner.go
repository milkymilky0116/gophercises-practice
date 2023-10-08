package urlshortner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

type urls []url

type jsonUrls struct {
	Urls []url `json:"urls"`
}
type url struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func (app *AppConfig) MapHandler(pathsToUrls map[string]string) {
	for url, redirect := range pathsToUrls {
		newUrl := url
		newRedirect := redirect
		app.Mux.HandleFunc(newUrl, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, newRedirect, http.StatusPermanentRedirect)
		})
	}
}

func (app *AppConfig) YAMLHandler(yml []byte) error {
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

func (app *AppConfig) JsonHandler(jsonBuff []byte) error {
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
