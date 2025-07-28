package main

import "fmt"

type Customer struct {
	Name          string
	AccountNumber string
}
type BankAccount struct {
	Balance  float64 // Số dư tài khoản
	Customer Customer
}

// hàm rút tiền từ tài khoản. nếu số tiền lớn hơn số dư hiện tại gọi hàm panic thông báo "số dư không đủ"
func (b *BankAccount) Withdraw(amount float64) string {
	if amount > b.Balance {
		panic("Số dư không đủ")
	}
	b.Balance -= amount
	return "Rút tiền thành công"

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Lỗi:", r.(string))
		}
	}()

	account := BankAccount{
		Balance: 1000.0,
		Customer: Customer{
			Name:          "VI Xuân Cử",
			AccountNumber: "123456789",
		},
	}

	fmt.Printf("Khách hàng: %s\n", account.Customer.Name)
	fmt.Printf("Số tài khoản: %s\n", account.Customer.AccountNumber)
	fmt.Printf("Số dư ban đầu: %.2f\n", account.Balance)
	var amount float64
	amount = 200.0 // Số tiền muốn rút
	fmt.Printf("Số tiền muốn rút: %.2f\n", amount)

	// Gọi hàm Withdraw để rút tiền
	result := account.Withdraw(amount)
	fmt.Println(result)
}
