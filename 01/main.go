package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	filePath := flag.String("file", "resources/problems.csv", "a CSV file in the format: question,answer")
	quizDuration := flag.Int("duration", 10, "the duration of the quiz in seconds")
	flag.Parse()

	quizQuestions, err := ProcessQuizCsv(*filePath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to process the file='%s'", *filePath), err.Error())
	}
	result := PlayGame(nil, quizQuestions, time.Duration(*quizDuration) * time.Second)
	fmt.Printf("Result: %d/%d\n", result, len(quizQuestions))
}
