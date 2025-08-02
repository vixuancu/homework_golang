package main

import "fmt"

// tính tổng diện tích hình học với interface
// định nghĩa interface Shape với các phương thức Area và Perimeter
type Shape interface {
	Area() float64
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
// hàm tính tổng diện tích của các hình
func totalArea(shapes []Shape) float64 {
	total:= 0.0
	for _,s := range shapes {
		total += s.Area()
	}
	return total
}




func main() {
	// tạo slice chứa các hình
	shapes := []Shape{
		Rectangle{Width: 4, Height: 5},
		Circle{Radius: 3},
		Rectangle{Width: 2, Height: 7},
	}

	total := totalArea(shapes)
	fmt.Printf("Tổng diện tích các hình: %.2f\n", total)
}