package main

import (
	"fmt"
	"sync"
)

var (
	balance int        = 1000 // Số dư tài khoản ban đầu
	mu      sync.Mutex        // Khóa để đồng bộ truy cập vào biến balance
)

func deposit(wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu goroutine đã hoàn thành khi kết thúc

	for i := 0; i < 100; i++ {
		mu.Lock()      // Khóa để đảm bảo an toàn khi truy cập vào balance
		balance += 200 // Thêm 200 vào số dư
		mu.Unlock()    // Mở khóa sau khi cập nhật xong
	}

}

func withdraw(wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu goroutine đã hoàn thành khi kết thúc
	for i := 0; i < 100; i++ {
		mu.Lock()           // Khóa để đảm bảo an toàn khi truy cập vào balance
		if balance >= 100 { // Kiểm tra xem có đủ tiền để rút không
			balance -= 100 // Rút 100 từ số dư
		}
		mu.Unlock() // Mở khóa sau khi cập nhật xong
	}

}
func main() {
	var wg sync.WaitGroup // Khởi tạo WaitGroup để đợi các goroutine hoàn thành

	wg.Add(2) // Thêm 2 goroutine vào WaitGroup

	go deposit(&wg)                                            // Goroutine để gửi tiền
	go withdraw(&wg)                                           // Goroutine để rút tiền
	wg.Wait()                                                  // Chờ tất cả goroutine hoàn thành
	fmt.Printf("Số dư tài khoản sau giao dịch: %d\n", balance) 
}
