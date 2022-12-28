package main

import "fmt"

func main() {
	name := "Eko"
	counter := 0

	increment := func() {
		name := "Budi"
		//name = "Budi" //ini akan mengubah variable yang di atas "Eko"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
