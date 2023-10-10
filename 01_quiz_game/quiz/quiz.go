package quiz

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/milkymilky0116/gophercises-practice/01_quiz_game/cli"
)

type Quiz struct {
	Quiz   map[string]string
	Result int
	index  int
	config cli.Config
}

func ParseCsv(filename string, config cli.Config) (*Quiz, error) {
	var q Quiz
	q.Quiz = map[string]string{}
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
	r, _ := regexp.Compile("[0-9]+")
	for _, row := range data {
		q.Quiz[row[0]] = strings.TrimSpace(r.FindString(row[1]))
	}
	q.config = config
	return &q, nil
}

func (q *Quiz) ParseUserInput() (*Quiz, error) {
	q.index = 1
	for question, answer := range q.Quiz {
		var userAnswer string

		channel := make(chan bool)
		errs := make(chan error, 1)
		outputQuestion := fmt.Sprintf("Problem #%d %s : ", q.index, question)

		go func() {
			fmt.Print(outputQuestion)
			reader := bufio.NewReader(os.Stdin)
			result, err := reader.ReadString('\n')
			if err != nil {
				errs <- err
			}
			result = result[:len(result)-1]
			userAnswer = strings.TrimSpace(result)
			if err != nil {
				errs <- err
			}
			channel <- true
		}()

		timeoutChan := time.After(time.Duration(q.config.Timer) * time.Second)
		select {
		case <-channel:
			userAnswer = strings.ToLower(userAnswer)
			answer = strings.ToLower(answer)
			if userAnswer == answer {
				q.Result++
			}
			q.index++

		case <-timeoutChan:
			err := errors.New("timeout exceed")
			return q, err
		}

	}
	return q, nil
}
