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
// OrderStatus 1 - completed, 0 - awaiting, -1 - canceled.
// DeliveryDelay in minutes (15-300 min step 15 min).
// Date format 2006-01-02 15:04:05.
type Order struct {
	ID            int     `json:"id"`
	CourierID     int     `json:"courier_id"`
	ClientID      int     `json:"client_id"`
	Product       string  `json:"product"`
	PaymentMethod string  `json:"payment_method"`
	QuantityTo    float64 `json:"quantity_to"`
	QuantityFrom  float64 `json:"quantity_from"`
	OrderCost     float64 `json:"order_cost"`
	OrderStatus   int     `json:"order_status"`
	DeliveryDelay int     `json:"delivery_delay"`
	DateStart     string  `json:"date_start"`
	DateFinish    string  `json:"date_finish"`
}
