package main

import "fmt"

// Struct Author
type Author struct {
	Name  string
	Email string
}

// Struct Book kết hợp với Author
type Book struct {
	Title  string
	Pages  int
	Author Author
}

// Method DisplayInfo với value receiver
func (b Book) DisplayInfo() {
	// defer thông báo khi kết thúc hàm
	defer fmt.Println("Kết thúc hiển thị thông tin sách.")
	fmt.Println("Thông tin sách:")
	fmt.Printf("Tiêu đề: %s\n", b.Title)
	fmt.Printf("Số trang: %d\n", b.Pages)
	fmt.Printf("Tác giả: %s\n", b.Author.Name)
	fmt.Printf("Email tác giả: %s\n", b.Author.Email)

}

func main() {
	// tạp 1 instance của book (instance là một đối tượng cụ thể của struct)
	book := Book{
		Title: "Go Programming",
		Pages: 300,
		Author: Author{
			Name:  "Vi Xuân Cử",
			Email: "vixuancu2004@gmail.com",
		},
	}
	// gọi method để hiển thị thông tin
	book.DisplayInfo()
}
