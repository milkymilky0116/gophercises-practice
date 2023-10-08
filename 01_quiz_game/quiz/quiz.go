package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/milkymilky0116/gophercises-practice/01_quiz_game/cli"
)

type Quiz struct {
	quiz   map[string]string
	result int
	config cli.Config
}

func ParseCsv(filename string, config cli.Config) (*Quiz, error) {
	var q Quiz
	q.quiz = map[string]string{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range data {
		q.quiz[row[0]] = row[1]
	}
	q.config = config

	return &q, nil
}

func (q *Quiz) ParseUserInput() (*Quiz, error) {
	for question, answer := range q.quiz {
		var userAnswer int
		outputQuestion := fmt.Sprintf("\nQuestion : %s ?\n", question)

		channel := make(chan bool)

		go func(done chan bool) {
			timer := q.config.Timer
			fmt.Println(outputQuestion)

			fmt.Print("Type Your Answer : ")
			fmt.Scan(&userAnswer)

			for i := 0; i < timer; i++ {
				time.Sleep(1 * time.Second)
				saveCursorPosition := "\033[s"
				fmt.Print(saveCursorPosition)
				fmt.Print(i)
			}
			clearLine := "\033[u\033[K"
			fmt.Print(clearLine)

			done <- true
		}(channel)
		timeoutChan := time.After(time.Duration(q.config.Timer) * time.Second)
		select {
		case <-channel:
			parsedAnswer, err := strconv.Atoi(answer)
			if err != nil {
				return nil, err
			}

			if userAnswer == parsedAnswer {
				q.result++
			}
		case <-timeoutChan:
			fmt.Println("\n\nOops! Timeout!")
		}

	}
	return q, nil
}
