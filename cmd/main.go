package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	CSV "quizgame/internal/csv"
	"time"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	csvFilename := flag.String("csv", "problems/problems.csv", "a csv file in the format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz(seconds)")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		msg := fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename)
		exit(msg)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := CSV.ParseLines(lines)
	var count int
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.Question)
		ansCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			ansCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou answered %d/%d questions correctly", count, len(problems))
			return
		case ans := <- ansCh:
			if ans != problem.Answer {
				fmt.Printf("Correct answer is: %s\n", problem.Answer)
			} else {
				count++
			}
		}
	}
	fmt.Printf("You answered %d/%d questions correctly", count, len(problems))
}
