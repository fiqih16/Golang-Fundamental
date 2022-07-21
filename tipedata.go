package main

import "fmt"

// Pada tugas kali ini kamu akan diberikan dua variabel yang pertama dalam bentuk integer dan yang kedua dalam bentuk float32. Tulislah code yang akan menampilkan pada console berupa :

// Baris pertama menampilkan penjumlahan dari dua variabel yang diberikan dalam bentuk integer

// Baris kedua menampilkan pengurangan dalam bentuk float dengan 2 koma

// Baris ketiga menampilkan perkalian dalam bentuk integer

func main() {
    var x int
    var y float64

	x = 5
	y = 12.32
    // scanner := bufio.NewScanner(os.Stdin)
    // scanner.Scan()
    // x, _ = strconv.Atoi(scanner.Text())
    // scanner.Scan()
    // y, _ = strconv.ParseFloat(scanner.Text(), 32)

	// Expected output
	// 17
	// -7.32
	// 60
    //Tulis kode disini
	result := (x + int(y))
	resultdua := (float64(x) - y)	
	resulttiga := (x * int(y))
	fmt.Println(result)
	fmt.Printf("%.2f\n", resultdua)
	fmt.Println(resulttiga)

}