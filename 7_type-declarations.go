package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	type Number int32

	var noKtpEko NoKTP = "18741982741897419874"
	var marriedStatus Married = true
	var Angka Number = 100
	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
	fmt.Println(Angka)
}
