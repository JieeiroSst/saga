package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("inventory.reserved", func(m *nats.Msg) {
		arrangeShipment(string(m.Data))
	})
}

func arrangeShipment(orderID string) {
	// Logic to arrange shipment
}
