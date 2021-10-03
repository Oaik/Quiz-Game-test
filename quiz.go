package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "problems.csv", "file path to read from")
	flag.Parse()
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	cnt := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(record)
		var text string
		fmt.Scanf("%s\n", &text)

		if text == record[1] {
			cnt = cnt + 1
		}
	}
	fmt.Printf("You scored %d from 12", cnt)
}
