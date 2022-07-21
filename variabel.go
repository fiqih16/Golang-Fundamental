package main

import "fmt"

func main() {
	//  Pada tugas ini kamu diminta melengkapi coding berikut untuk menukar nilai variabel a dan b lalu tampilkanlah pada console.
	//  a = Camp, b = Enigma
	// maka nilai a menjadi Enigma dan b menjadi Camp.
	var a,b string
	a = "Camp"
	b = "Enigma"
	fmt.Println("a =",a,"b =",b)
	a,b = b,a
	fmt.Println("a =",a,"b =",b)
 	fmt.Scanln(&a)
 	fmt.Scanln(&b)

 	




}