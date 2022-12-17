package main

import "fmt"

func main() {
	fizzBuzz(30)
}

// Solution for the FizzBuzz problem
func fizzBuzz(n int) {
	for i := 1; i < n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("fizz buzz\n")
			break
		case i%3 == 0:
			fmt.Printf("fizz\n")
			break
		case i%5 == 0:
			fmt.Printf("buzz\n")
			break
		default:
			fmt.Printf("%d\n", i)
			break
		}
	}
}
