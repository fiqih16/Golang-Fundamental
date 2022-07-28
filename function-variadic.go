package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	total = len(number)
	return total
}

func main() {
	angka := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(angka)
}