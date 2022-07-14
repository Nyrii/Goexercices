package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var folder = "resources"

func TestParseCsvFile(t *testing.T) {
	var testRecords = []struct {
		fileName              string
		expectedQuizQuestions []*QuizQuestion
		expectedError         interface{}
	}{
		{
			"valid_problems_test.csv",
			[]*QuizQuestion{
				{
					Question: "1+1",
					Result:   "2",
				},
				{
					Question: "1+2",
					Result:   "3",
				},
				{
					Question: "1+3",
					Result:   "4",
				},
			},
			nil,
		},
		{
			"unknown_file.csv",
			nil,
			&CsvOpeningError{fileName: fmt.Sprintf("%s/%s", folder, "unknown_file.csv")},
		},
		{
			"unformatted_problems_test.csv",
			nil,
			&CsvParsingError{fileName: fmt.Sprintf("%s/%s", folder, "unformatted_problems_test.csv")},
		},
		{
			"invalid_problems_test.csv",
			nil,
			&CsvParsingError{fileName: fmt.Sprintf("%s/%s", folder, "invalid_problems_test.csv")},
		},
	}
	for _, testRecord := range testRecords {
		path := fmt.Sprintf("%s/%s", folder, testRecord.fileName)
		actualQuizQuestions, err := ProcessQuizCsv(path)
		if len(testRecord.expectedQuizQuestions) == 0 {
			assert.NotNil(t, err, "expected an error")
			assert.IsType(t, testRecord.expectedError, err, "expected error of a certain type")
		} else {
			assert.Nil(t, err, "expected no error")
			assert.ElementsMatch(t, testRecord.expectedQuizQuestions, actualQuizQuestions,"expected quiz questions to be identical")
		}
	}
}
