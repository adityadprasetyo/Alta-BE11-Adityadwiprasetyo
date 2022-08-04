package main

import "fmt"

func main() {
	var t float64 = 20
	var r float64 = 4
	var pi float64 = 3.14
	answer := (2 * pi * r * r) + (2 * pi * r * t)
	fmt.Println(answer)
}
