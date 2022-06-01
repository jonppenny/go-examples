package main

import (
	"fmt"
	"log"
)

func main() {
	var score int

	rs, err := ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	for i, r := range rs {
		data := Record{
			question: r[0],
			answer:   r[1],
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
