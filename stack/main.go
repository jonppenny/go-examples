package main

import "fmt"

func main() {
	s := Stack{}
	fmt.Println(s)
	s.Push(100)
	s.Push(200)
	s.Push(300)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}
