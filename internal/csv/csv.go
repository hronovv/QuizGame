package csv

import (
	"quizgame/internal/problem"
	"strings"
)

func ParseLines(lines [][]string) []problem.Problem {
	problems := make([]problem.Problem, len(lines))
	for i, line := range lines {
		problems[i] = problem.Problem{
			Question: strings.TrimSpace(line[0]),
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}
