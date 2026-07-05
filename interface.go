package main

import "fmt"

type paymenter interface {
	pay(amt float64)
}
type payment struct {
	gateway paymenter
}

func (p *payment) makePayment(amt float64) {
	p.gateway.pay(1000)
}

type razorpay struct {
}

type stripe struct {
}

func (s *stripe) pay(amt float64) {
	fmt.Println("making payemnt through stripe : ", amt)
}
func (r *razorpay) pay(amt float64) {
	fmt.Println("making payment using razorpay : ", amt)
}

type fakepayment struct {
}

func (f *fakepayment) pay(amt float64) {
	fmt.Println("making payment using fakepayment : ", amt)
}

func main() {
	//razorpayGw := &razorpay{}
	newPayment := payment{
		gateway: &stripe{},
	}
	newPayment.makePayment(1000)
}
