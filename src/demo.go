package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")
	fmt.Println(Sum(3)) // 0+1+2
}

// Nhận tham số truyền vào là int và trả về kiểu int
func Sum(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}
