package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func newOrder(id string, amount float32, status string, time time.Time) *order {
	return &order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time,
	}
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) changeAmmount(amount float32) {
	o.amount = amount
}

func main() {
	var myorder order = order{
		id: "1",
		//amount: 1.0,
		//status: "pending",
	}
	myorder.createdAt = time.Now()
	fmt.Println(myorder.id, myorder.amount, myorder.status, myorder.createdAt)
	fmt.Println(myorder)

	myorder.changeStatus("paid")
	fmt.Println(myorder.status)

	myorder.changeAmmount(567.09)
	fmt.Println(myorder.amount)

	muOrder2 := order{
		id:        "2",
		amount:    100.0,
		status:    "delivered",
		createdAt: time.Now(),
	}

	fmt.Println(muOrder2)

	order3 := newOrder("3", 100.0, "paid", time.Now())
	fmt.Println(order3)
	fmt.Println(order3.status)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)
}
