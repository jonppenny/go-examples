package main

import "fmt"

var nums = []int{1, 3, 4, 7, 8, 2}
var revString = "test"

func main() {
	tot := sum(nums)

	fmt.Printf("Sum is %d\n", tot)

	var index int
	fmt.Printf("Please enter a number to search for:\n")
	fmt.Scanln(&index)

	search := find(index, nums)

	fmt.Printf("Index of %d is %d\n\n", index, search)

	fmt.Printf("%s", reverse(revString))
}

// Sum the numbers in a slice
func sum(s []int) int {
	t := 0
	for _, i := range s {
		t += i
	}
	return t
}

// Find the index of an int in a slice
func find(a int, s []int) int {
	var idx = 0
	for i, v := range s {
		if v == a {
			idx = i
			return idx
		}
	}
	return -1
}

// Reverse a string
func reverse(w string) string {
	var s string
	for i := len(w) - 1; i >= 0; i-- {
		s += string(w[i])
	}

	return s
}
