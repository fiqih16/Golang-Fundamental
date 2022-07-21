package main

import "fmt"

func main() {
	// kondisi menggunakan if else
	// number := 2
	// number2 := 3
	// if number%2 == 0 && number2%2 == 1 {
	// 	fmt.Println("Nilai Genap dan Ganjil")
	// } else if number%2 == 0 {
	// 	fmt.Println("Nilai Genap")
	// } else {
	// 	fmt.Println("Nilai Ganjil")
	// }

	// if waktu := 3; waktu >= 6 && waktu <= 10 {
	// 	fmt.Println("Sarapan Pagi")
	// 	} else if waktu > 10 && waktu <= 15 {
	// 		fmt.Println("Makan Siang")
	// 	} else {
	// 		fmt.Println("Makan Malam")
	// 	}

	// kondisi menggunakan switch case
	//CARA SATU
	// case diisi dengan nilai yang akan di cek
	tipeProduk := "makanan"
	switch tipeProduk {
	case "makanan":
		fmt.Println("Produk Makanan")
	case "minuman":
		fmt.Println("Produk Minuman")
	default:
		fmt.Println("Produk Lainnya")
	} 

	//CARA DUA
	var nilai = 70
	switch {
	case nilai > 80:
		fmt.Println("Nilai Anda A")
	case nilai > 70:
		fmt.Println("Nilai Anda B")
	default:
		fmt.Println("Nilai Anda C")
	}

	// KUIS SELEKSI KONDISI
	// Bantulah Susi untuk menentukan jadwal sehari-harinya dengan sebuah program sederhana. Akan diberikan inputan jam berupa integer lalu tampilkanlah pada console kegiatan yang akan dilakukan Susi sesuai jam yang diberikan. Kegiatan yang akan dilakukan Susi adalah berikut :
	// Jika jam 4-5 akan mencetak “Bangun Pagi”
	// Jika jam 6-7 akan mencetak “Mandi dan Sarapan”.
	// Jika jam 8-11 akan mencetak “Berangkat Sekolah”.
	// Jika jam 12 akan mencetak “Pulang Sekolah”.
	// Jika jam 13-15 akan mencetak “Tidur Siang”.
	// Jika Jam 16-17 akan mencetak “Waktu Main”
	// Selain dari itu akan mencetak “Waktu Belajar dan Istirahat”
	// Bila user menginput melebihi 24 jam maka akan mencetak “Hanya ada 24 jam dalam sehari”
	// Sample input 1:
	// 9
	// Sample Input 2:
	// 25
	// Sample Input 3:
	// 23
	// Sample output 1:
	// Berangkat Sekolah
	// Sample Output 2:
	// Hanya ada 24 Jam dalam sehari
	// Sample Output 3:
	// Waktu Belajar dan Istirahat
	var jam int
		fmt.Scanln(&jam)	
	switch {
		case jam >= 4 && jam <= 5:
			fmt.Println("Bangun Pagi")
		case jam >= 6 && jam <= 7:
			fmt.Println("Mandi dan Sarapan")
		case jam >= 8 && jam <= 11:
			fmt.Println("Berangkat Sekolah")
		case jam == 12:
			fmt.Println("Pulang Sekolah")
		case jam >= 13 && jam <= 15:
			fmt.Println("Tidur Siang")
		case jam >= 16 && jam <= 17:
			fmt.Println("Waktu Main")
		case jam >= 24:
			fmt.Println("Hanya ada 24 Jam dalam sehari")
		default:
			fmt.Println("Waktu Belajar dan Istirahat")
		} 
}