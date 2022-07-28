package main

func main() {
	// CARA 1
	// var myVal int = 64
	// var myValPointer *int = &myVal
	// fmt.Println("myVal:", myVal)
	// fmt.Println("myValPointer:", myValPointer)

	// CARA 2
	// var myVal2 *int
	// fmt.Println("myVal2:", myVal2)

}

// KUIS
// Aldi ingin memberi Budi beberapa kelereng yang dimilikinya. Pada tugas kali ini kamu hanya diminta untuk membuat sebuah fungsi bernama giveMarble dimana fungsi ini memiliki deskripsi sebagai berikut :
// - Memiliki 3 parameter
// - Parameter pertama sebagai pemberi kelereng
// - Parameter kedua sebagai penerima kelereng
// - Parameter ketiga sebagai jumlah kelereng yang akan diberikan
// - Jika Jumlah kelereng yang diberikan melebihi jumlah yang dimiliki oleh pemberi maka akan menampilkan pada console “Kelereng tidak cukup untuk dibagikan”
// - Fungsi tidak memiliki nilai balikan namun dapat mengubah value dari pemberi dan penerima

// Deskripsi Input :
// - Input baris pertama sebagai banyak kelereng pemberi
// - Input baris kedua sebagai banyak kelereng penerima
// - Input baris ketiga sabagai banyak kelereng yang akan dibagikan

// Deskripsi output:
// - Baris pertama sebagai sisa kelereng pemberi
// - Baris kedua sebagai banyak kelereng penerima

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	giver, _ := strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	receiver, _ := strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	marble, _ := strconv.Atoi(scanner.Text())

// 	giveMarble(&giver, &receiver, marble)
// 	fmt.Printf("%d \n", giver)
// 	fmt.Printf("%d", receiver)
// }

// Jawaban Sendiri
// func giveMarble(giver *int, receiver *int, marble int) {
// 	if *giver < marble {
// 		fmt.Println("Kelereng tidak cukup untuk dibagikan")
// 	} else {
// 		*giver -= marble
// 		*receiver += marble
// 	}
// }

// Jawaban Enigcamp
// func giveMarble(giver, receiver *int, total int) {
// if *giver < total {
//     fmt.Println("Kelereng tidak cukup untuk dibagikan")
//     return
// }
// *giver -= total
// *receiver += total
// }