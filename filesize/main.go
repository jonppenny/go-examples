package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sizeConvert(2854125))
}

func sizeConvert(bytes float64) string {
	var msg string

	if bytes >= 1099511627776.0 {
		msg = fmt.Sprintf("%f TB", getFloatPrecision(bytes/1099511627776.0, 2))
	} else if bytes >= 1073741824 {
		msg = fmt.Sprintf("%f GB", getFloatPrecision(bytes/1073741824.0, 2))
	} else if bytes >= 1048576.0 {
		msg = fmt.Sprintf("%f MB", getFloatPrecision(bytes/1048576.0, 2))
	} else if bytes >= 1024 {
		msg = fmt.Sprintf("%f KB", getFloatPrecision(bytes/1024.0, 2))
	} else {
		msg = fmt.Sprintf("%f B", getFloatPrecision(bytes/1.0, 2))
	}

	return msg
}

func getFloatPrecision(value, precision float64) float64 {
	return math.Floor(value*math.Pow(10, precision)+0.5) / math.Pow(10, precision)
}
