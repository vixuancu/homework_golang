package main

import (
	"fmt"
	"math"
)

// Hàm tính chỉ số BMI = cân nặng / (chiều cao)^2
func calculateBMI(weight float64, height float64) float64 {
	if height == 0 {
		return 0
	}
	return weight / math.Pow(height, 2)
}

func main() {
	var weight, height float64

	fmt.Print("Nhập cân nặng (kg): ")
	fmt.Scanln(&weight)

	fmt.Print("Nhập chiều cao (m): ")
	fmt.Scanln(&height)

	bmi := calculateBMI(weight, height)

	fmt.Printf("\nChỉ số BMI của bạn là: %.2f\n", bmi)
	fmt.Print("Phân loại sức khỏe: ")

	if bmi < 18.5 {
		fmt.Println("Thiếu cân")
	} else if bmi < 25 {
		fmt.Println("Bình thường")
	} else if bmi < 30 {
		fmt.Println("Thừa cân")
	} else {
		fmt.Println("Béo phì")
	}
}
