package main

import "fmt"

// phân biệt kiểu dữ liệu với empty interface 
func printType(val interface{}) {
	switch val.(type) {// sử dụng type assertion để phân biệt kiểu dữ liệu
	case int:
		fmt.Println("Đây là kiểu số nguyên")
	case float64:
		fmt.Println("đây là kiểu số thực")
	case string:
		fmt.Println("đây là kiểu chuỗi")
	default:
		fmt.Println("Kiểu dữ liệu không xác định")
	}
}
func main() {
	printType(10)
	printType(3.14)
	printType("Hello")
	printType(true)
}
