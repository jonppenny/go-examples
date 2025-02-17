package main

import (
	"fmt"

	"go-examples/queue/helpers"
)

func main() {
	q := helpers.Queue{}
	fmt.Println(q)
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
