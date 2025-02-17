package main

import "fmt"

func main() {
	sum := 0
	nums := []int{1, 2, 4, 8}

	for _, v := range nums {
		sum = sum + v
	}

	fmt.Printf("%d\n", sum)
}
