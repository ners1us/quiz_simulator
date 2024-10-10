package main

import (
	"flag"
	"fmt"
	"quiz_simulator/src/quiz"
	"quiz_simulator/src/utils"
	"strings"
)

func main() {
	printWelcomeMessage()

	var choice int
	if !getUserChoice(&choice) || choice == 2 {
		utils.Exit("Goodbye!")
	}

	timeVal, err := getTimeInput()
	if err != nil {
		utils.Exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}

	fileName := flag.String("f", "quiz.csv", "path of csv file")
	timerVal := flag.Int("t", timeVal, "timer for the quiz")
	flag.Parse()

	problems, err := quiz.ProblemPuller(*fileName)
	if err != nil {
		utils.Exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}

	startQuiz(problems, *timerVal)
}

func printWelcomeMessage() {
	fmt.Println("Welcome to the quiz simulator!")
	fmt.Println()
	fmt.Println("Choose an action:")
	fmt.Println("1. Continue")
	fmt.Println("2. Exit")
}

func getUserChoice(choice *int) bool {
	_, err := fmt.Scan(choice)

	return err == nil && (*choice == 1 || *choice == 2)
}

func getTimeInput() (int, error) {
	var timeVal int

	fmt.Print("Enter time for the timer in seconds: ")
	_, err := fmt.Scan(&timeVal)

	if err != nil {
		return 0, err
	}

	err = utils.CheckTime(&timeVal)

	return timeVal, err
}

func startQuiz(problems []quiz.Problem, timeVal int) {
	countCorrectAns := 0
	timerObj := utils.StartTimer(timeVal)

	answerChannel := make(chan string)

problemLoop:
	for i, p := range problems {
		var ans string

		fmt.Printf("Problem %d -> %s: ", i+1, p.Q)

		go func() {
			_, err := fmt.Scanf("%s", &ans)
			if err != nil {
				utils.Exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
			}

			answerChannel <- ans
		}()

		select {
		case <-timerObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-answerChannel:
			if iAns == p.A || iAns == strings.ToLower(p.A) {
				countCorrectAns++
			}
			if i == len(problems)-1 {
				close(answerChannel)
			}
		}
	}
	fmt.Printf("Your result is %d out of %d\n", countCorrectAns, len(problems))
}
