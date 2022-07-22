package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Macam macam array
	// var tipeProduk [2]string{"", ""}
	// tipeProduk := [2]string{"", ""}

	// CARA PERTAMA
	// var tipeProduk [2]string = [2]string{"Komik", "Game"}
	// fmt.Println(tipeProduk)

	// CARA KEDUA
	// tipeProduk := [2]string{
	// 	"Komik",
	// 	"Game",
	// }
	// fmt.Println(tipeProduk)

	// CARA KETIGA
	// tipeProduk := [...]string{
	// 	"Komik",
	// 	"Game",
	// 	"Novel",
	// 	"Musik",
	// 	"Film",
	// }
	// Cara ubah data array
	// tipeProduk[3] = "Musik diubah"
	// fmt.Println(tipeProduk)

	// CARA KEEMPAT
	// var num = [4][3]string{
	// 	[3]string{"1", "2", "3"},
	// 	[3]string{"4", "5", "6"},
	// 	[3]string{"", "0", ""},
	// }
	// for i := 0; i < len(num); i++ {
	// 	fmt.Printf("baris ke %d ,isinya %v \n", i+1, num[i])
	// }

	// CARA KELIMA
	// direction := [4]string{"north", "south", "east", "west"}
	// for i, direction := range direction {
	// 	fmt.Printf("Element %d : %s \n",i, direction)
	// }

	// KUIS
	// Doni ingin mengetahui ada berapa angka genap yang terdapat di sebuah array, hanya saja doni butuh bantuan untuk mengetahuinya karena elemen di dalam array tersebut sangat banyak. Tulislah code untuk menerima 2 inputan dari Doni,
	// - Inputan pertama adalah kapasitas dari array
	// - Inputan kedua adalah elemen dari array tersebut.
	// Lalu cetak angka-angka genap tersebut ke dalam console.
	// Input satu:
	// 2
	// 32 91
	// Input dua:
	// 	6
	// 22 90 45 43 2 57
	// Output satu:
	// 32
	// Output dua:
	// 	22
	// 90
	// 2

	// CARA YANG SEDERHANA
	// var input int
	// fmt.Scan(&input)
	// for i := 0; i < input; i++ {
	// 	var number int
	// 	fmt.Scan(&number)
	// 	if number % 2 == 0 {
	// 		fmt.Println(number)
	// 	}
	// }

	// CARA YANG BENAR
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	capacity, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, capacity)

	scanner.Scan()
	arrText := scanner.Text()
	arrText2 := strings.Split(arrText, " ")

	for i, v := range arrText2 {
		x, _ := strconv.Atoi(string(v))
		arr[i] = x
	}

	for _, v := range arr {
		if v%2 == 0 && v != 0 {
			fmt.Println(v)
		}
	}
}