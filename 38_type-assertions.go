package main

import "fmt"

func apasaja() interface{} {
	return "1"
}

func main() {
	var result interface{} = apasaja()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
