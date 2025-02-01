package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Printf("%d", int(C.random()))
}
