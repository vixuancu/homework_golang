package main

import (
	"fmt"
)

// Hàm kiểm tra số nguyên tố
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var start, end int
	fmt.Print("Nhập số bắt đầu: ")
	fmt.Scanln(&start)

	fmt.Print("Nhập số kết thúc: ")
	fmt.Scanln(&end)

	// Đảm bảo start <= end
	if start > end {
		start, end = end, start
	}

	count := 0
	sum := 0

	for i := start; i <= end; i++ {
		if isPrime(i) {
			sum += i
			count++
		}
	}

	fmt.Printf("\nTổng các số nguyên tố trong khoảng [%d, %d] là: %d\n", start, end, sum)
	fmt.Printf("Có %d số nguyên tố trong khoảng này.\n", count)
}
