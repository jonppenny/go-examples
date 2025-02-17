package main

import "fmt"

var x = map[int]string{
	0: "one",
	1: "two",
	2: "three",
	3: "four",
}

func main() {
	for k, v := range x {
		fmt.Printf("%d : %s\n", k, v)
	}
}
