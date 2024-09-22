package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("order.created", func(m *nats.Msg) {
		processPayment(string(m.Data))
	})

	createOrder(nc)
}

func createOrder(nc *nats.Conn) {
	// Logic to create order
	nc.Publish("order.created", []byte("orderID123"))
}

func processPayment(orderID string) {
	// Logic to process payment
}
