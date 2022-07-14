package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type QuizQuestion struct {
	Question string
	Result string
}

type CsvOpeningError struct {
	fileName string
}

func (error *CsvOpeningError) Error() string {
	return fmt.Sprintf("Unable to open CSV file='%s'", error.fileName)
}

type CsvParsingError struct {
	fileName string
}

func (error *CsvParsingError) Error() string {
	return fmt.Sprintf("Unable to parse CSV file='%s'", error.fileName)
}

func ProcessQuizCsv(path string) ([]*QuizQuestion, error) {
	quizQuestions := make([]*QuizQuestion, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, &CsvOpeningError{
			fileName: filepath.Base(path),
		}
	}
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, &CsvParsingError{
			fileName: filepath.Base(path),
		}
	}
	for _, line := range lines {
		quizQuestions = append(quizQuestions, &QuizQuestion{
			Question: line[0],
			Result: line[1],
		})
	}
	return quizQuestions, nil
}
