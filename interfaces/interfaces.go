package main

import (
	"fmt"
	"log"
)

func ProcessTransaction(p PaymentProcessor, amount float64) {
	// business logic here
	result, err := p.ProcessPayment(amount)
	if err != nil {
		log.Printf("ProcessPayment failed: %s", err)
		return
	}
	fmt.Println(result)
}

func main() {
	var payment PaymentProcessor

	payment = PayPal{}
	ProcessTransaction(payment, 100.0)

	payment = Stripe{}
	ProcessTransaction(payment, 75.0)
}

type PaymentProcessor interface {
	ProcessPayment(amount float64) (string, error)
}

// PayPal Payment Processor
type PayPal struct{}

func (p PayPal) ProcessPayment(amount float64) (string, error) {
	return fmt.Sprintf("Processed payment of $%.2f via PayPal", amount), nil
}

// Stripe Payment Processor
type Stripe struct{}

func (s Stripe) ProcessPayment(amount float64) (string, error) {
	// return "", fmt.Errorf("something went wrong in Stripe")
	return fmt.Sprintf("Processed payment of $%.2f via Stripe", amount), nil
}
