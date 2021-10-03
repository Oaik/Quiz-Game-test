package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func parseLine(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for idx, val := range lines {
		ret[idx] = problem{val[0], val[1]}
	}
	return ret
}

func getAnswer(ch chan string) {
	var text string
	fmt.Scanf("%s\n", &text)
	ch <- text
}

func main() {
	fptr := flag.String("fpath", "problems.csv", "file path to read from")
	timeLimit := flag.Int("limit", 1, "Time limit for the quiz")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	cnt := 0
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseLine(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for idx, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", idx+1, problem.question)
		inputAnswer := make(chan string)
		go getAnswer(inputAnswer)

		select {
		case <-timer.C:
			fmt.Printf("You scored %d from 12", cnt)
			return
		case text := <-inputAnswer:
			if text == problem.answer {
				cnt++
			}
		}
	}
	fmt.Printf("You scored %d from 12", cnt)
}
