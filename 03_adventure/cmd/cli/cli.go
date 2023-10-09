package cli

import (
	"flag"
	"io"
)

type httpConfig struct {
	Port     int
	Filename string
}

func ParseArgs(w io.Writer, args []string) (*httpConfig, error) {
	var c httpConfig
	fs := flag.NewFlagSet("Adventure Book", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&c.Port, "port", 4000, "Set port to run server")
	fs.StringVar(&c.Filename, "json", "gopher.json", "Set json adventure file to parse")
	err := fs.Parse(args)

	if err != nil {
		return nil, err
	}

	return &c, err
}
