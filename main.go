package main

import (
	"bufio"
	"fmt"
	"os"
)

type answer struct {
	text      string
	isCorrect bool
}

type question struct {
	question string
	answer1  answer
	answer2  answer
	answer3  answer
}

type questionBank struct {
	questions []question
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func printQuestions(questions questionBank) {
	correct := 0
	tries := 0
	for i := 0; i < len(questions.questions); i++ {
		fmt.Println("Question:  " + questions.questions[i].question)
		fmt.Println("A:\t" + questions.questions[i].answer1.text)
		fmt.Println("B:\t" + questions.questions[i].answer2.text)
		fmt.Println("C:\t" + questions.questions[i].answer3.text)

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		checkError(err)

		switch char {
		case 'A':
			if questions.questions[i].answer1.isCorrect == true {
				fmt.Println("You are correct")
				correct++
				tries++
			} else {
				fmt.Println("Please try again")
				i--
				tries++

			}

		case 'B':
			if questions.questions[i].answer2.isCorrect == true {
				fmt.Println("You are correct")
				correct++
				tries++
			} else {
				fmt.Println("Please try again")
				i--
				tries++
			}

		case 'C':
			if questions.questions[i].answer3.isCorrect == true {
				fmt.Println("You are correct")
				correct++
				tries++
			} else {
				fmt.Println("Please try again")
				i--
				tries++
			}

		default:
			fmt.Println("Invalid Answer")
			i--
		}

	}
	fmt.Printf("Score %d out of %d tries\n", correct, tries)
}

func ChooseLesson() error {
	fmt.Print("Please choose a lesson number 1-7")

	reader := bufio.NewReader(os.Stdin)
	//char, _, err := reader.ReadRune()
	lessonNum, _, err := reader.ReadRune()
	checkError(err)
	switch lessonNum {
	case '1':
		printQuestions(questionsL1)
	case '2':
		printQuestions(questionsL2)
	case '3':
		printQuestions(questionsL3)
	case '4':
		printQuestions(questionsL4)
	case '5':
		printQuestions(questionsL5)
	case '6':
		printQuestions(questionsL6)
	case '7':
		printQuestions(questionsL7)
	default:
		fmt.Println("Invalid Input")

	}
	return err

}

func main() {
	ChooseLesson()
}
