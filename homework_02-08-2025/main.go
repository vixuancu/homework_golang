package main

import (
	"fmt"
)

// Notifier interface định nghĩa hợp đồng chung cho mọi trình gửi thông báo
type Notifier interface {
	Send(message string) error
}

// ConsoleNotifier là component cơ bản, gửi thông báo ra console
type ConsoleNotifier struct{}

func (c *ConsoleNotifier) Send(message string) error {
	fmt.Printf("Thông báo qua Console: %s\n", message)
	return nil
}

// SmsDecorator là decorator bọc một Notifier để thêm hành vi gửi SMS
type SmsDecorator struct {
	notifier Notifier
}

// NewSmsDecorator khởi tạo SmsDecorator với một Notifier
func NewSmsDecorator(notifier Notifier) *SmsDecorator {
	return &SmsDecorator{notifier: notifier}
}

func (s *SmsDecorator) Send(message string) error {
	fmt.Println("[SMS Decorator] Gửi thông báo qua SMS...")
	if s.notifier != nil {
		return s.notifier.Send(message)
	}
	return nil
}

// SlackDecorator là decorator bọc một Notifier để thêm hành vi gửi Slack
type SlackDecorator struct {
	notifier Notifier
}

// NewSlackDecorator khởi tạo SlackDecorator với một Notifier
func NewSlackDecorator(notifier Notifier) *SlackDecorator {
	return &SlackDecorator{notifier: notifier}
}

func (s *SlackDecorator) Send(message string) error {
	fmt.Println("[Slack Decorator] Gửi thông báo qua Slack...")
	if s.notifier != nil {
		return s.notifier.Send(message)
	}
	return nil
}

func main() {
	message := "Hệ thống sẽ bảo trì vào lúc 2 giờ sáng."

	// Kịch bản 1: Chỉ gửi qua Console
	fmt.Println("Kịch bản 1: Chỉ gửi qua Console")
	var notifier Notifier = &ConsoleNotifier{}
	notifier.Send(message)

	// Kịch bản 2: Gửi qua Console, bọc bởi SMS
	fmt.Println("\nKịch bản 2: Console + SMS")
	var notifierWithSms Notifier = NewSmsDecorator(&ConsoleNotifier{})
	notifierWithSms.Send(message)

	// Kịch bản 3: Gửi qua Console, bọc bởi SMS, rồi bọc bởi Slack
	fmt.Println("\nKịch bản 3: Console + SMS + Slack")
	baseNotifier := &ConsoleNotifier{}
	smsNotifier := NewSmsDecorator(baseNotifier)
	slackNotifier := NewSlackDecorator(smsNotifier)
	slackNotifier.Send(message)
}