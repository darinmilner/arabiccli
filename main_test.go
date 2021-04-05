package main

import (
	"fmt"
	"testing"
)

var mockQ1 = questionBank{
	[]question{
		{question: "How many letters are in the Arabic alphabet?",
			answer1: answer{text: "28", isCorrect: false},
			answer2: answer{text: "29", isCorrect: true},
			answer3: answer{text: "30", isCorrect: false}},
	},
}

var mockQ2 = questionBank{
	[]question{
		{question: "How many letters are in the Arabic alphabet?",
			answer1: answer{text: "28", isCorrect: false},
			answer2: answer{text: "29", isCorrect: true},
			answer3: answer{text: "30", isCorrect: false}},
	},
}

func TestChooseLesson(t *testing.T) {
	err := MockChooseLesson()
	if err != nil {
		t.Error("chooseLesson()", err)
	}
}

func MockChooseLesson() error {
	fmt.Print("Please choose a lesson number")
	// reader := bufio.NewReader(os.Stdin)
	// //char, _, err := reader.ReadRune()
	// lessonNum, _, err := reader.ReadRune()
	var err error
	lessonNum := '1'
	if err != nil {
		checkError(err)
	}

	switch lessonNum {
	case '1':
		fmt.Println(mockQ1.questions[0])
	case '2':
		fmt.Println(mockQ2.questions[0])
	default:
		fmt.Println("Invalid Input")

	}
	return err

}
