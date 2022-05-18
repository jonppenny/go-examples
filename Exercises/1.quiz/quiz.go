package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type csvRecords struct {
	question string
	answer   string
}

func main() {
	var score int

	records, err := readFile("problems.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, record := range records {
		data := csvRecords{
			question: record[0],
			answer:   record[1],
		}
		fmt.Println("What is " + data.question)
		var ua string
		fmt.Scanln(&ua)
		if ua == data.answer {
			score++
		}
	}

	fmt.Printf("You scored: %d\n", score)
}

// Read a file
func readFile(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	cr := csv.NewReader(f)
	records, err := cr.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
