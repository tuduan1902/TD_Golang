package main

import "fmt"

// thư viện fmt

func main() {
	// var number int
	// number = 10
	// fmt.Println(number)
	// var number1 int = 11
	// fmt.Println(number1)

	// Type Inference
	// var email = "tudaun@gmail.com"
	// fmt.Println(email)

	// Khai báo nhiều biến
	// 1. Khai báo nhiều biến cùng kiểu dữ liệu
	// var a, b int = 10, 11
	// fmt.Println(a, b)
	// 2. Khai báo nhiều biến khác kiểu dữ liệu
	var (
		name    string = "Moriaty"
		address string = "TP.HCM"
		age     int    = 25
	)

	fmt.Println(name, address, age)

}
