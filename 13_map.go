package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	newmap := map[string]int{
		"angka1": 10,
		"angka2": 20,
	}

	fmt.Println(newmap)

	newmap2 := map[int]string{
		10: "sepuluh",
		20: "dua puluh",
	}

	fmt.Println(newmap2)

	delete(newmap2, 20)
	fmt.Println(newmap2)

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
