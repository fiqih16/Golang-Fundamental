package main

import "fmt"

type Address struct {
	City, Provice string
}

func templateAddress(address *Address) {
	address.Provice = "DKI Jakarta"
}

func main() {
	address := Address{
		City:    "Jaksel",
		Provice: "",
	}

	// variabel address menggunakan pointer
	var addressPointer *Address = &address
	templateAddress(addressPointer)

	fmt.Println(address)

}