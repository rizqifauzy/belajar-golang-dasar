package main

import "fmt"

func main() {
	const (
		firstName string = "Eko"
		lastName         = "Khannedy"
		value            = 1000
	)
	//firstName = "Tidak berubah" //error
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
