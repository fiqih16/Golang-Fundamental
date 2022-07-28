package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//func main() {
// CARA SATU
// direction := [4]string{"north", "south", "east", "west"}
// var vertikaldirection = direction[0:2]

// fmt.Printf("%v , data type %T \n", direction, direction)
// fmt.Printf("%v , data type %T \n", vertikaldirection, vertikaldirection)

// CARA DUA
// direction := []string{"north", "south", "east", "west"}
// fmt.Println(direction)

// CARA TIGA
// Menggunakan make
// s := make([]string, 3)
// fmt.Println("angka", s)

// CARA EMPAT
// direction := []string{"north", "south", "east", "west"}
// var vertikaldirection = direction[2:]

// fmt.Printf("%v , data type %T \n", direction, direction)
// fmt.Printf("%v , data type %T \n", vertikaldirection, vertikaldirection)
// fmt.Printf("%s located in %p \n", direction[2], direction[2])
// fmt.Printf("%s located in %p \n", vertikaldirection[0], vertikaldirection[0])

// CARA LIMA
// slice dan append
// direction := []string{"north", "south", "east", "west"}
// var vertikaldirection = direction[0:2]
// Menggunakan SLICE
// fmt.Printf("isi slice %v, jumlah element %d, kapasitas slice %d, in located %p \n", vertikaldirection, len(vertikaldirection), cap(vertikaldirection), vertikaldirection)
// fmt.Printf("isi slice %v, jumlah element %d, kapasitas slice %d, in located %p \n", direction, len(direction), cap(direction), direction)
// Menggunakan APPEND
// newVerticalDirection := append(vertikaldirection, "Northwest")
// fmt.Printf("isi slice %v, jumlah element %d, kapasitas slice %d, in located %p \n", newVerticalDirection, len(newVerticalDirection), cap(newVerticalDirection), newVerticalDirection)
// fmt.Printf("%p", &direction[0])
// fmt.Printf("%p", &vertikaldirection[0])
// fmt.Printf("%p", &newVerticalDirection[0])

// KUIS
// Pada tugas kali ini kamu diminta untuk mengumpulkan nama-nama yang memiliki jumlah karakter genap dari daftar nama yang diberikan.
// Function description:
// Lengkapilah function evenNames dalam editor. Function evenNames memiliki parameter berupa slice string dan harus mengembalikan nilai berupa slice string juga.
// Input :
// Herman budi jenny kevin aris
// Barry Levi Willy Ferry
// Output :
// Herman budi aris
// Levi
// KUIS CARA PERTAMA (SOLUSI Enicamp)
	func evenNames(slice []string) []string {
		var names []string
		for _, v := range slice {
			if len(v)%2 == 0 {
				names = append(names, v)
			}
		}
		return names
	}
 
	func main() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		x := scanner.Text()
		slice := strings.Split(x, " ")
		names := evenNames(slice)
		for _, name := range names {
			fmt.Print(name + " ")
		}
		
	}

	// KUIS CARA KEDUA (SOLUSI Sendiri)
	// func evenNames(slice []string) []string {
	// 	//Tulis kode disini
	// 	var names []string
	// 	for _, name := range slice {
	// 		if len(name) % 2 == 0 {
	// 			fmt.Printf("%s ", name)
	// 		}
	// 	}
	// 	return names
	// }
	

	// func main() {
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	x := scanner.Text()
	// 	slice := strings.Split(x, " ")
	// 	evenNames(slice)
	// }

	
//}