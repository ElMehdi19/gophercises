package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, 3)
	for i, line := range lines {
		problems[i] = problem{q: line[0], a: line[1]}
		if i == 2 {
			break
		}
	}
	return problems
}

func main() {
	fname := flag.String("csv", "problems.csv", "CSV file in `question,answer` format")
	flag.Parse()

	file, err := os.Open(*fname)
	if err != nil {
		exit(fmt.Sprintf("Couldn't open file %s", *fname))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Couldn't parse file\nError: %s", err))
	}

	problems := parseLines(lines)
	counter := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s= ", i+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.a {
			counter++
		}
	}

	fmt.Printf("You've scored %d out of %d", counter, len(problems))
}
