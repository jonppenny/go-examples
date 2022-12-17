package main

import "fmt"

func main() {
	a, b := 12, 432
	fmt.Println(swap(&a, &b))
}

func swap(i *int, j *int) (int, int) {
	var t int
	t = *i
	*i = *j
	*j = t
	return *i, *j
}
