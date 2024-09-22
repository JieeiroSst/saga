package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("payment.processed", func(m *nats.Msg) {
		if err := reserveInventory(string(m.Data)); err != nil {
			nc.Publish("payment.compensated", []byte(m.Data))
		} else {
			nc.Publish("inventory.reserved", []byte(m.Data))
		}
	})
}

func reserveInventory(orderID string) error {
	// Logic to reserve inventory
	return nil
}
