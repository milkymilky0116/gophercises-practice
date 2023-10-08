package cli

import (
	"flag"
	"os"
)

type Config struct {
	Ymlfile  string
	Jsonfile string
}

func ParseArgs(args []string) (*Config, error) {
	var c Config
	fs := flag.NewFlagSet("URL Shortener", flag.ContinueOnError)
	fs.StringVar(&c.Ymlfile, "yaml", "test.yaml", "Give yaml file name that you want to parse.")
	fs.StringVar(&c.Jsonfile, "json", "test.json", "Give json file name that you want to parse.")
	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func OpenFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
