package quiz

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

type Problem struct {
	Q string
	A string
}

func ProblemPuller(filename string) ([]Problem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s file; %s", filename, err.Error())
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading data from %s file; %s", filename, err.Error())
	}

	return parseProblem(lines), nil
}

func parseProblem(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for i := range lines {
		problems[i] = Problem{Q: lines[i][0], A: lines[i][1]}
	}

	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})

	return problems
}
