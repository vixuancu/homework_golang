package main

func sumNumbers(ch chan int) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	ch <- sum // Gửi kết quả vào kênh
}
func main() {
	ch := make(chan int) // Tạo một kênh để nhận kết quả

	go sumNumbers(ch) // Gọi hàm sumNumbers trong goroutine
	result := <-ch // Nhận kết quả từ kênh
	println("Tổng các số từ 1 đến 100 là:", result) 
}
/*
Channel không có buffer (make(chan int)):
Khi main gọi result := <-ch, nó sẽ block (chờ) cho đến khi có dữ liệu được gửi vào ch.
Goroutine sumNumbers phải chạy xong vòng lặp và thực hiện ch <- sum thì main mới tiếp tục được.

Cơ chế đồng bộ tự nhiên:
Nhờ channel, goroutine sumNumbers và main được đồng bộ hóa.
main không thể kết thúc trước khi nhận được giá trị từ goroutine (vì đang bị block).

Nếu không có channel (hoặc không chờ):
Nếu bạn bỏ dòng result := <-ch, thì main có thể chạy xong ngay và kết thúc chương trình, khiến goroutine chưa kịp tính toán.
*/
