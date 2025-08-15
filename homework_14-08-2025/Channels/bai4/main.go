package main

func sumNumbers(ch chan int) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	ch <- sum // Gửi kết quả vào kênh
}
func main() {
	ch := make(chan int)

	go sumNumbers(ch) // Gọi hàm sumNumbers trong goroutine
	result := <-ch // Nhận kết quả từ kênh
	println("Tổng các số từ 1 đến 100 là:", result) 
}
