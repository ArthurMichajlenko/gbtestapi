// Generated by https://quicktype.io

package main

import "encoding/json"

// Clients is array of client
type Clients []Client

// UnmarshalClients decode clients from JSON
func UnmarshalClients(data []byte) (Clients, error) {
	var r Clients
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal encode clients to JSON
func (r *Clients) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Client is a single clients
type Client struct {
	ID      int64  `json:"id"`      
	Name    string `json:"name"`    
	Tel     string `json:"tel"`     
	Address string `json:"address"` 
	OrderID int64  `json:"order_id"`
}