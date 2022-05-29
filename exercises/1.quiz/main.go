package main

import (
	"fmt"
	"log"
)

func main() {
	var score int

	records, err := ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	for i, record := range records {
		data := Record{
			question: record[0],
			answer:   record[1],
		}
		fmt.Printf("Question %d. What is %s\n", i+1, data.question)
		var ua string
		fmt.Scanln(&ua)
		if ua == data.answer {
			score++
		}
	}

	fmt.Printf("You scored: %d\n", score)
}
