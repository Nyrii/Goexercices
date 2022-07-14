package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

var timeout = 10 * time.Second

func PlayGame(input *os.File, quizQuestions []*QuizQuestion) (result int) {
	if input == nil {
		input = os.Stdin
	}

	ch := make(chan int, len(quizQuestions))

	go func() {
		for _, quizQuestion := range quizQuestions {
			fmt.Println(quizQuestion.Question)
			userInput := readInput(input)
			if cleanUserInput(userInput) == quizQuestion.Result {
				ch <- 1
			}
		}
		close(ch)
	}()
	for {
		select {
		case _, hasValue := <-ch:
			if hasValue {
				result += 1
			} else {
				return result
			}
		case <-time.After(timeout):
			if ch != nil {
				close(ch)
			}
			return result
		}
	}
}

func readInput(input *os.File) string {
	var userInput string

	_, err := fmt.Fscanf(input, "%s\n", &userInput)
	if err != nil {
		fmt.Println("Fill a valid answer")
		return readInput(input)
	}
	return userInput
}

func cleanUserInput(userInput string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, userInput)
}
