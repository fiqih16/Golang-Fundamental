package main

import "fmt"

// type Member func(string) bool

// func registerMember(name string, member Member) {
// 	if member(name) {
// 		fmt.Println("Hello", name)
// 	} else {
// 		fmt.Println("Sorry no member", name)
// 	}
// }

// func main() {
// 	memberList := func(name string) bool {
// 		return name == "member"
// 	}

// 	registerMember("member", memberList)
// 	registerMember("non member", memberList)

// }

// KUIS (Cut off the bamboo)
// Setiap bambu memiliki sekat. Buatlah aplikasi yang dapat menentukan jumlah bambu yang kamu miliki melalui inputan, setelah itu kamu dapat menentukan jumlah sekat dari setiap bambu yang kamu miliki. Jika sudah, potonglah bambu beberapa kali berdasarkan inputan yang kamu inginkan. Setiap potongan yang kamu lakukan akan menghilangkan 1 sekat disetiap bambu yang kamu miliki. Lalu tampilkanlah pada console sisa sekat bambu setiap pemotongan.
// Kriteria:
//   > Gunakan variabel untuk menerima inputan dari console
//   > Gunakan function

// Sample Input :
// Masukan banyak bamboo : 2
// sekat bamboo ke-1 : 5
// sekat bamboo ke-2 : 3
// Berapa kali potong? 2

// Sample Output:
// Potongan ke- 1
// 4
// 2
// Potongan ke- 2
// 3
// 1

// Input pertama:
// 2
// 3
// 4
// 2
// Output pertama:
// Potongan ke- 1
// 2
// 3
// Potongan ke- 2
// 1
// 2
// Input kedua:
// 5
// 2
// 5
// 4
// 1
// 3
// 3
// Output kedua:
// Potongan ke- 1
// 1
// 4
// 3
// 0
// 2
// Potongan ke- 2
// 0
// 3
// 2
// 0
// 1
// Potongan ke- 3
// 0
// 2
// 1
// 0
// 0

func main() {
    var kumpulanSekatBamboo []int
    var banyakBamboo int
    var cuts int
 
   
    fmt.Scanln(&banyakBamboo)
    for i := 0; i < banyakBamboo; i++ {
        var SekatBamboo int
        fmt.Scanln(&SekatBamboo)
        kumpulanSekatBamboo = append(kumpulanSekatBamboo, SekatBamboo)
    }

    fmt.Scanln(&cuts)
 
    for i := 1; i <= cuts; i++ {
        fmt.Println("Potongan ke-", i)
        kumpulanSekatBamboo = decreaseBamboos(kumpulanSekatBamboo)
        printBamboos(kumpulanSekatBamboo)
    }
 
}
 
func decreaseBamboos(bamboos []int) []int {
    for i := 0; i < len(bamboos); i++ {
        if bamboos[i] == 0 {
            continue
        }
        bamboos[i] -= 1
    }
    return bamboos
}
 
func printBamboos(bamboos []int) {
    for i := 0; i < len(bamboos); i++ {
        printBamboo(bamboos[i])
    }
}
 
func printBamboo(bamboo int) {
    fmt.Println(bamboo)
}