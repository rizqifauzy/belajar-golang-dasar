package main

import "fmt"

func getFullName(name string) (string, string, string) {
	return "Hello " + name, "Kurniawan", "Khannedy"
}

func main() {
	firstName, _, _ := getFullName("rizqi")
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}
