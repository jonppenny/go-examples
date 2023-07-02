package main

import "fmt"

func main() {
	fmt.Println("Enter a whole number")
	var num int
	fmt.Scanln(&num)

	for i := 0; i <= num; i++ {
		fizzBuzz(i)
	}
}

// Solution for the FizzBuzz problem
func fizzBuzz(i int) {
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
