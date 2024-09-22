package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("order.created", func(m *nats.Msg) {
		if err := processPayment(string(m.Data)); err != nil {
			compensateOrder(string(m.Data))
		} else {
			nc.Publish("payment.processed", []byte(m.Data))
		}
	})
}

func processPayment(orderID string) error {
	// Logic to process payment
	return nil
}

func compensateOrder(orderID string) {
	// Logic to compensate order creation
}
