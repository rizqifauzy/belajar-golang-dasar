package main

import "fmt"

func main() {

	//var name1 = "Eko"
	//name1 := "Eko"
	//var name2 = "Budi"
	//name2 := "Budi"
	//var result bool = name1 == name2
	name1, name2 := "Eko", "Budi"
	result := name1 == name2
	fmt.Println(result) //false

	value1, value2 := 100, 200

	fmt.Println(value1 > value2)  //false
	fmt.Println(value1 < value2)  //true
	fmt.Println(value1 == value2) //false
	fmt.Println(value1 != value2) //true

}
