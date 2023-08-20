package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// proccessQuestions - parse questions and answers into 2 slices
func processQuestions(data [][]string) ([]string, []string) {
	var questions []string
	var answers []string
	for i := 0; i < len(data); i++ {
		questions = append(questions, data[i][0])
		answers = append(answers, data[i][1])
	}

	return questions, answers
}

func main() {
	// check args for fileName or default to data.csv
	args := os.Args
	fileName := "data.csv"
	if len(args) > 1 {
		fileName = args[1]
	}

	// try to open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// try to read from the csv file
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// initialize score
	score := 0
	questions, answers := processQuestions(data)

	// create a timer read it from args or default it to 30 seconds
	var timeout time.Duration = 30
	if len(args) > 2 {
		customTime, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}

		timeout = time.Duration(customTime)
	}

	timer1 := time.NewTimer(timeout * time.Second)
	for i := range questions {
		// in case user runs out of time display the current score and exit the program
		go func() {
			for range timer1.C {
				fmt.Println("Your final score is:", score)
				os.Exit(0)
			}
		}()

		// print questions and wait for user answer
		fmt.Println(questions[i])
		var userAnswer string
		fmt.Scanln(&userAnswer)

		// if it's correct answer add +1 to score
		if answers[i] == userAnswer {
			score++
		}
	}

	// print final score in case user got to the end without the time to end
	fmt.Println("Your final score is:", score)
}
