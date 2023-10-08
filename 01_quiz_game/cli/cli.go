package cli

import "flag"

type Config struct {
	FileName string
	Timer    int
}

func ParseArgs(args []string) (Config, error) {
	var c Config
	fs := flag.NewFlagSet("Quiz Game", flag.ContinueOnError)
	fs.StringVar(&c.FileName, "filename", "problems.csv", "enter your csv file name")
	fs.IntVar(&c.Timer, "timer", 30, "Time Limit to solving quiz")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	return c, nil
}
