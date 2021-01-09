package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
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
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{q: line[0], a: line[1]}
	}
	return problems
}

func main() {
	// program entrypoint goes here
	fname := flag.String("csv", "problems.csv", "CSV File in `question,answer` format.")
	flag.Parse()

	// open file and store file handler in a variable
	file, err := os.Open(*fname)
	if err != nil {
		exit(fmt.Sprintf("Coudln't open file %s", *fname))
	}
	defer file.Close()

	// parse csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Couldn't parse file! Error: %s", err))
	}

	problems := parseLines(lines)
	counter := 0
	total := 4

	timer := time.NewTimer(5 * time.Second)
	fmt.Println("TIMER STARTED, 2 SECONDS TO GO")
	for i, problem := range problems {
		if total == i {
			break
		}

		answerChan := make(chan string)
		go func() {
			fmt.Printf("Problem #%d: %s=", i+1, problem.q)
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou've scored %d out of %d\n", counter, total)
			return

		case answer := <-answerChan:
			if answer == problem.a {
				counter++
			}
		}
	}
	fmt.Printf("You've scored %d out of %d\n", counter, total)
}
