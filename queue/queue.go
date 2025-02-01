package main

import "fmt"

func main() {
	q := Queue{}
	fmt.Println(q)
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
