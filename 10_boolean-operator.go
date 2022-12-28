package main

import "fmt"

func main() {

	ujian, absensi := 80, 75

	lulusUjian := ujian >= 80
	lulusAbsensi := absensi >= 80
	fmt.Println(lulusUjian)   // true
	fmt.Println(lulusAbsensi) //false

	lulus := lulusAbsensi && lulusUjian
	fmt.Println(lulus) //false

	fmt.Println(ujian >= 80 && absensi >= 80) //false

}
