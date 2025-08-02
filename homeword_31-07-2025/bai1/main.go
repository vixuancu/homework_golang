package main

import "fmt"

// hình học với interface
// định nghĩa interface Shape với các phương thức Area và Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// định nghĩa struct Rectangle có width và height
type Rectangle struct {
	Width  float64
	Height float64
}
// định nghĩa struct Circle có bán kính
type Circle struct {
	Radius float64
}
// receiver method Area cho Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// receiver method Area cho Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Perimeter() float64 {
return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// hàm in diện tích và chu vi của hình 
func printShapeInfo(s Shape, name string) {
	fmt.Printf("%s: Area = %.2f, Perimeter = %.2f\n", name, s.Area(), s.Perimeter())
}



func main() {
	r:= Rectangle{Width: 3,Height: 4}
	C := Circle{Radius: 3}

	printShapeInfo(r, "Rectangle")
	printShapeInfo(C, "Circle")

}