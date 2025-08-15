package main

import (
	"fmt"
	"time"
)

func main() {
	ch1:=make(chan string)
	ch2:=make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Giả lập một tác vụ mất thời gian
		ch1 <- "from one" // Gửi giá trị vào kênh ch1
	}()
	go func() {
		time.Sleep(1 * time.Second) // Giả lập một tác vụ mất thời gian
		ch2 <- "from two" // Gửi giá trị vào kênh ch2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}
}
