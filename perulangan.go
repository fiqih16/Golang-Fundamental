package main

import (
	"fmt"
)

func main() {
	//CARA SATU
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//CARA DUA
	// for i := 1; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	//CARA TIGA
	// i := 5
	// for {
	// 	fmt.Println("Loop ke %d \n",i)
		
	// 	if i == 0 {
	// 		break
	// 	}
	// 	i--
	// 	time.Sleep(1 * time.Second)
	// }

	//CARA EMPAT
	// for n := 0; n<= 10; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Bilangan Ganjil %d \n",n)
	// }

	// KUIS
	//Pada tugas ini kamu diminta untuk mencetak kalimat "I will become a go developer" pada console sebanyak inputan yang akan diberikan. Agar tantangan ini menjadi lebih menarik tambahkanlah suatu angka mundur di depan setiap kalimatnya.
	//Kriteria :
	//> Decrement
	//> Siapkan variabel untuk menerima inputan dari console
	// Sample input : 2
	// Sample output :
	// 2  I will become a go developer
	// 1  I will become a go developer

	var input int
	fmt.Scan(&input)
	for i := input; i > 0; i-- {
		fmt.Println(i, " I will become a go developer")
	}
}