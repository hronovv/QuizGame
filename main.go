package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func ParseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			Question: strings.TrimSpace(line[0]),
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func main() {
	csvFilename := flag.String("csv", "problems/problems.csv", "a csv file in the format 'question,answer'")
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
	problems := ParseLines(lines)
	var count int
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.Question)
		var ans string
		fmt.Scanf("%s", &ans)
		if ans != problem.Answer {
			fmt.Printf("Correct answer is: %s\n", problem.Answer)
		} else {
			count++
		}
	}
	fmt.Printf("You answered %d/%d questions correctly", count, len(problems))
}
