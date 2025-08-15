package main

import (
	"fmt"
	"sync"
)

func worker(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu goroutine đã hoàn thành khi kết thúc
	for job := range jobs {
		fmt.Printf("Worker received %d, squared: %d\n ", job, job*job)
	}

}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 3)
	wg.Add(1) // Thêm 1 goroutine vào WaitGroup
	// Tạo worker goroutine
	go worker(jobs, &wg)

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	// Gửi xong thì đóng channel
	close(jobs)
	wg.Wait() // Chờ worker goroutine hoàn thành
	fmt.Println("Worker finished ")
}
