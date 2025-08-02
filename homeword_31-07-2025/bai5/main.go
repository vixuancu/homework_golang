package main

import "fmt"

// interface kế thừa interface
type Reader interface {
	Read() string
}
type Writer interface {
	Write(data string)
}

// định nghĩa interface kế thừa từ Reader và Writer
type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	Content string
}
// phương thức write cho File
func (f *File) Write(data string) {
	f.Content = data
	fmt.Println("Đã ghi dữ liệu:", data)
}
// phương thức read cho File
func (f *File) Read() string {
	return f.Content
}
// hàm thao tác với ReadWriter
func process(rw ReadWriter,input string) {
	rw.Write(input)
	fmt.Println("Đã Đọc dữ liệu:", rw.Read())
}

func main() {
	f := &File{} // tạo đối tượng File với con trỏ
	process(f, "Hello, World!") // gọi hàm process với đối tượng File
}