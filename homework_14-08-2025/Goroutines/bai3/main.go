package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchURL(url string,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Fetching URL: %s\n", url)
	time.Sleep(2 * time.Second) // Giả lập thời gian lấy dữ liệu
	fmt.Printf("Finished fetching URL: %s\n", url)
}

func main() {
	urls := []string{
		"url1.com",
		"url2.com",
		"url3.com",
		"url4.com",
		"url5.com",
	}
	var wg sync.WaitGroup 
	wg.Add(len(urls)) // Thêm số lượng goroutine cần chờ
	for _, url := range urls {
		go fetchURL(url, &wg) // Gọi hàm fetchURL trong goroutine
	}
	wg.Wait() // Chờ tất cả goroutine hoàn thành
	fmt.Println("All URLs have been fetched!")
}
