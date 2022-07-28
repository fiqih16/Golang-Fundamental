package main

import "fmt"

func sayHello(firstNamr string, lastName string) {
	fmt.Printf("Hello %s %s \n", firstNamr, lastName)
}

func increment(angka int) int {
	return angka + 1
}

func main() {
	sayHello("Steven", "Wiliam")
	angka := increment(5)
	fmt.Println(angka)
}