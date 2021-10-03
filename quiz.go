package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
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

func main() {
	fptr := flag.String("fpath", "problems.csv", "file path to read from")
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
	for idx, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", idx+1, problem.question)
		var text string
		fmt.Scanf("%s\n", &text)
		if text == problem.answer {
			cnt = cnt + 1
		}
	}
	fmt.Printf("You scored %d from 12", cnt)
}
