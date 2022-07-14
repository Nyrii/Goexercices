package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
	"time"
)

func TestPlayGame(t *testing.T) {
	quizQuestions := []*QuizQuestion{
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
	}
	var testRecords = []struct {
		quizQuestions []*QuizQuestion
		userInputs []string
		expectedResult int
	}{
		{
			quizQuestions,
			[]string{" 2  ", "3", "4"},
			3,
		},
		{
			quizQuestions,
			[]string{"2", "100", "200"},
			1,
		},
		{
			quizQuestions,
			[]string{},
			0,
		},
		{
			[]*QuizQuestion{},
			[]string{},
			0,
		},
	}

	for _, testRecord := range testRecords {
		in, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}

		for _, userInput := range testRecord.userInputs {
			_, err = io.WriteString(in, userInput + "\n")
			if err != nil {
				t.Fatal("Unable to write user inputs", err)
			}
		}

		_, err = in.Seek(0, io.SeekStart)
		if err != nil {
			t.Fatal(err)
		}

		actualResult := PlayGame(in, testRecord.quizQuestions, 1 * time.Second)

		in.Close()
		assert.Equal(t, testRecord.expectedResult, actualResult, "expected identical outcome")
	}
}
