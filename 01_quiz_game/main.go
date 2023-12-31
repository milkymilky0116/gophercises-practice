package main

import (
	"fmt"
	"log"
	"os"

	"github.com/milkymilky0116/gophercises-practice/01_quiz_game/cli"
	"github.com/milkymilky0116/gophercises-practice/01_quiz_game/quiz"
)

func main() {
	args := os.Args[1:]
	config, err := cli.ParseArgs(args)
	if err != nil {
		log.Println(err)
	}
	questions, err := quiz.ParseCsv(config.FileName, config)
	if err != nil {
		log.Println(err)
	}
	result, err := questions.ParseUserInput()

	if err != nil {
		log.Println(err)
	}

	resultOutput := fmt.Sprintf("You scored %d out of %d", result.Result, len(result.Quiz))

	fmt.Println(resultOutput)
}
