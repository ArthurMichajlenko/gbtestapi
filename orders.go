// Generated by https://quicktype.io

package main

import "encoding/json"

// Orders is array of orders
type Orders []Order

// UnmarshalOrders decode orders from JSON
func UnmarshalOrders(data []byte) (Orders, error) {
	var r Orders
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode orders to JSON
func (r *Orders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Order is a single order.
// Confirmation 1 - completed, 0 - awaiting, -1 - canceled.
// DeliveryDelay in minutes.
// Date format 2006-01-02 15:04:05.
type Order struct {
	ID            int64  `json:"id"`
	Product       string `json:"product"`
	PaymentMethod string `json:"payment_method"`
	QuantityTo    int64  `json:"quantity_to"`
	QuantityFrom  int64  `json:"quantity_from"`
	Confirmation  int64  `json:"confirmation"`
	DeliveryDelay int64  `json:"delivery_delay"`
	DateStart     string `json:"date_start"`
	DateFinish    string `json:"date_finish"`
}
