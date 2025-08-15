package main

import (
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup // Khởi tạo WaitGroup để đợi các goroutine hoàn thành

	numGoroutines := 5 // Số lượng goroutine cần chạy
	wg.Add(numGoroutines) // Thêm số lượng goroutine vào WaitGroup

	for i := 1; i <= numGoroutines; i++ {
		go func (id int)  {
			defer wg.Done() // Đánh dấu goroutine đã hoàn thành khi kết thúc
			fmt.Printf("Goroutine %d: %d\n", id,i)
		}(i)
	}
	wg.Wait() // chờ tất cả goroutine hoàn thành
}