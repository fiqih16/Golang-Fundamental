package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Cara pertama
	// bulan := map[int]string{
	// 	1: "Januari",
	// 	2: "Februari",
	// 	3: "Maret",
	// }
	// bulan[4] = "April"
	// fmt.Println(bulan)

	// Cara kedua
	// var sales = make(map[string]int)
	// sales["Penjualan"] = 100
	// sales["Pembelian"] = 50
	// fmt.Println(sales)
	// for k, v := range sales {
	// 	fmt.Println("sales %s sejumlah %d \n", k, v)
	// }

	// Cara ketiga
	// var value, isExist = sales["Penjualan"]
	// if isExist {
	// 	fmt.Println("value:", value)
	// } else {
	// 	fmt.Println("value:", "tidak ada")
	// }

	// KUIS
	// Pak Tejo ingin melihat penjualan pada beberapa bulan. Kali ini tugas kamu adalah membantu Pak Tejo untuk mencari penjualan pada bulan yang diinginkannya. Pak Tejo akan menentukan bulannya dengan angka urutan pada bulan yang diinginkannya. Kamu akan disediakan slice yang berisikan nama-nama bulan secara berurutan dan map yang menampung data penjualan Pak Tejo. Tampilkanlah pada console sesuai dengan format  sebagai berikut :
	// “Bulan (nama bulan) : (banyak penjualan) tusuk”
	// Explanation :
	// Input baris pertama pak tejo menentukan pengecekan penjualan mulai dari bulan apa.
	// Input baris kedua menentukan bulan terakhir untuk pengecekan.
	// Lalu tampilkanlah sesuai dengan format yang sudah ditentukan.
	// Constraints :
	// 0 < (input ke-1) < 13
	// 0 < (input ke-2) < 13
	// Output :
	// Bulan februari : 3282 tusuk
	// Bulan maret : 787 tusuk
	// Bulan april : 6238 tusuk
	// Bulan mei : 1992 tusuk
	month := [12]string{"januari", "februari", "maret", "april", "mei", "juni",
        "juli", "agustus", "september", "oktober", "november", "desember"}

    penjualan := map[string]int{
        "januari":   2836,
        "februari":  3282,
        "maret":     787,
        "april":     6238,
        "mei":       1992,
        "juni":      824,
        "juli":      2903,
        "agustus":   602,
        "september": 930,
        "oktober":   1002,
        "november":  922,
        "desember":  3219,
    }

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    bulan1, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    bulan2, _ := strconv.Atoi(scanner.Text())

	// Jawaban Sendiri
	//Tulis kode disini
	for i := bulan1; i <= bulan2; i++ {
		fmt.Printf("Bulan %s : %d tusuk\n", month[i-1], penjualan[month[i-1]])
	}

	// Jawaban Enigcamp
	// for i := bulan1 - 1; i < bulan2; i++ {
    //     for k, v := range penjualan {
    //         if month[i] == k {
    //             fmt.Printf("Bulan %v : %d tusuk\n", k, v)
    //         }
    //     }
    // }
}