package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/startSaga", startSaga)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startSaga(w http.ResponseWriter, r *http.Request) {
	orderID := createOrder()
	if err := processPayment(orderID); err != nil {
		compensateOrder(orderID)
		return
	}
	if err := reserveInventory(orderID); err != nil {
		compensatePayment(orderID)
		compensateOrder(orderID)
		return
	}
	arrangeShipment(orderID)
}

func createOrder() string {
	// Logic to create order
	return "orderID123"
}

func processPayment(orderID string) error {
	// Logic to process payment
	return nil
}

func reserveInventory(orderID string) error {
	// Logic to reserve inventory
	return nil
}

func arrangeShipment(orderID string) {
	// Logic to arrange shipment
}

func compensateOrder(orderID string) {
	// Logic to compensate order creation
}

func compensatePayment(orderID string) {
	// Logic to compensate payment
}
