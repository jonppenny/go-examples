package main

func main() {
	a, b := 12, 432
	swap(&a, &b)
	println(a, b)
}

func swap(i *int, j *int) (int, int) {
	var t int
	t = *i
	*i = *j
	*j = t
	return *i, *j
}
