package urlshortner

import "errors"

type urls []url

type jsonUrls struct {
	Urls []url `json:"urls"`
}
type url struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

var ErrInvalidNumofArgs error = errors.New("invalid number of arguments")
var ErrInvalidSubCmd error = errors.New("invalid sub command")
var ErrInvalidPositionalArgs error = errors.New("invalid number of positional arguments")
