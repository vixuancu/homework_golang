package main

import "fmt"

type Vehicle interface {
	Start() string
	Stop() string
}
type Car struct{}
type Bike struct{}

func (c Car) Start() string {
	return "Car is starting"
}
func (c Car) Stop() string {
	return "Car is stopping"
}
func (b Bike) Start() string {
	return "Bike is starting"
}
func (b Bike) Stop() string {
	return "Bike is stopping"
}

// hàm nhận vào 1 interface Vehicle và gọi các phương thức Start và Stop
func operateVehicle(v Vehicle) {
	fmt.Println(v.Start())
	fmt.Println(v.Stop())
}

func main() {

	car := Car{}
	operateVehicle(car)
	bike := Bike{}
	operateVehicle(bike)
}