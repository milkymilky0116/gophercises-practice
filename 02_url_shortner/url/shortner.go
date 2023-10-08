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
		app.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, redirect, http.StatusPermanentRedirect)
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
		app.Mux.HandleFunc(config.Path, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, config.Url, http.StatusPermanentRedirect)
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
		fmt.Println(config.Path, config.Url)
		app.Mux.HandleFunc(config.Path, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, config.Url, http.StatusPermanentRedirect)
		})
	}
	return nil
}
