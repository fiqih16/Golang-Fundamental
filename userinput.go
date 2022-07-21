package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cara 1
	// fmt.Printf("Masukan alamat : ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// homeAddress := scanner.Text()
	// fmt.Println("Alamat rumah anda : \n ", homeAddress)

	// Cara 2 dengan Menggunakan OS Argumen
	// argsWithProg := os.Args
	// fmt.Println(argsWithProg[1])
	// fmt.Println(argsWithProg[2])
	// fmt.Println(argsWithProg[3])

	// Cara 3 dengan Menggunakan OS Argumen fleg
	// fullNameArg := flag.String("Name", "", "Nama lengkap")
	// homeAddressArgs := flag.String("Address", "", "Alamat rumah")
	// isgoldMember := flag.Bool("Gold", false, "Jika Member Gold")
	// flag.Parse()
	// fmt.Println("Nama lengkap : ", *fullNameArg)
	// fmt.Println("Alamat lengkap : ", *homeAddressArgs)
	// fmt.Println("Member Gold : ", *isgoldMember)

	// KUIS USER INPUT
	// Tulislah code yang akan menerima inputan berupa string lalu tampilkanlah di console. Inputan yang akan diberikan berupa :
	// Nama
	// Alamat
	// Email
	// Kriteria : Siapkan variabel untuk menerima inputan dari console
	// Sample output :
	// Halo, saya Tian. Saya tinggal di Jakarta. Alamat email saya adalah Tian@gmail.com
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	scanner.Scan()
	address := scanner.Text()
	scanner.Scan()
	email := scanner.Text()
	fmt.Printf("Halo, saya %s. Saya tinggal di %s. Alamat email saya adalah %s", name, address, email)
}