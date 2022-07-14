package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	filePath := flag.String("file", "resources/problems.csv", "a CSV file in the format: question,answer")
	flag.Parse()

	quizQuestions, err := ProcessQuizCsv(*filePath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to process the file='%s'", *filePath), err.Error())
	}
	result := PlayGame(nil, quizQuestions)
	fmt.Printf("Result: %d/%d\n", result, len(quizQuestions))
}
