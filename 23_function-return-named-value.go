package main

import "fmt"

func getFullName2(name string) (firstName string, middleName string, lastName string) {
	firstName = "Hallo " + name
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return
}

func main() {
	a, b, c := getFullName2("Rizqi")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
