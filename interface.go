package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// type Student interface {
// 	StudentName() string
// }

// // Contoh interface kosong
// func Hi(name string) interface{}{
//     return "Hi " + name
// }

// func getStudent(student Student) {
// 	fmt.Println("Halo, nama saya ", student.StudentName())
// }

// type Biodata struct {
// 	Nama string
// }

// func (biodata Biodata) StudentName() string {
// 	return biodata.Nama
// }

// func main() {
// 	var biodata1 Biodata
// 	biodata1.Nama = "Doni"
// 	getStudent(biodata1)

// 	var kosong interface{} = Hi("Doni")
// 	fmt.Println(kosong)
// }

// KUIS
// Pada tugas kali ini kamu diminta untuk membuat kontrak antara struct Segitiga dan interface BangunDatar yang sudah disediakan.
// Untuk melakukan kontrak dengan sebuah interface, sebuah tipe data harus memilki (mengimplementasikan) setiap method yang dideklarasi interface tersebut.

type BangunDatar interface {
    Luas() int
}

//kode struct disini
type Segitiga struct {
    alas int
    tinggi int
}

//kode struct method disini
func (s *Segitiga) Luas() int {
    return s.alas * s.tinggi / 2
}

func getLuas(bangunDatar BangunDatar) {
    fmt.Printf("Luas bangun datar = %d \n", bangunDatar.Luas())
}

func main() {
   var segitiga1 Segitiga
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    segitiga1.alas, _ = strconv.Atoi(scanner.Text())
    scanner.Scan()
    segitiga1.tinggi, _ = strconv.Atoi(scanner.Text())
    getLuas(&segitiga1)

}