// Generated by https://quicktype.io

package main

import "encoding/json"

// Couriers is array of courier
type Couriers []Courier

// UnmarshalCouriers decode couriers from JSON
func UnmarshalCouriers(data []byte) (Couriers, error) {
	var r Couriers
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode couriers to JSON
func (r *Couriers) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Courier is a single courier
type Courier struct {
	ID              int64           `json:"id"`              
	Name            string          `json:"name"`            
	CarNumber       string          `json:"car_number"`      
	CurrentLocation CurrentLocation `json:"current_location"`
}

// CurrentLocation is a location of the courier
type CurrentLocation struct {
	Latitude  float64 `json:"latitude"` 
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`  
}