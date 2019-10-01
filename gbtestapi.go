/*
 * Filename: /Users/arthur/go/src/gbtestapi/gbtestapi.go
 * Path: /Users/arthur/go/src/gbtestapi
 * Created Date: Wednesday, September 18th 2019, 3:13:21 pm
 * Author: Arthur Michajlenko
 *
 * Copyright (c) 2019 Infinit Loop SRL
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	log.Println("Start create Couriers JSON file...")
	var couriersList Couriers
	for i := 0; i < 20; i++ {
		var courier Courier
		courier.ID = i
		courier.Imei = 123456789012300 + int64(i)
		courier.Tel = fmt.Sprintf("Tel_%d", i)
		courier.Name = fmt.Sprintf("Name_%d", i)
		courier.CarNumber = fmt.Sprintf("CarNumber_%d", i)
		courier.CurrentLocation.Latitude = 0
		courier.CurrentLocation.Longitude = 0
		courier.CurrentLocation.Address = fmt.Sprintf("Address_%d", i)
		couriersList = append(couriersList, courier)
	}
	writeData, _ := couriersList.Marshal()
	err := ioutil.WriteFile("couriers.json", writeData, 0666)
	if err != nil {
		log.Println(err)
	}
	log.Println("Stop create Couriers JSON file...")

	log.Println("Start create Clients JSON file...")
	var clientsList Clients
	for i := 0; i < 20; i++ {
		var client Client
		client.ID = i
		client.Name = fmt.Sprintf("Name_%d", i)
		client.Tel = fmt.Sprintf("Tel_%d", i)
		client.Address = fmt.Sprintf("Address_%d", i)
		clientsList = append(clientsList, client)
	}
	writeData, _ = clientsList.Marshal()
	err = ioutil.WriteFile("clients.json", writeData, 0666)
	if err != nil {
		log.Println(err)
	}
	log.Println("Stop create Clients JSON file...")

	log.Println("Start create Orders JSON file...")
	var ordersList Orders
	for i := 0; i < 20; i++ {
		var order Order
		order.ID = i
		order.CourierID = i
		order.ClientID = i
		order.Product = fmt.Sprintf("Product_%d", i)
		order.PaymentMethod = "Cash"
		order.QuantityTo = float64(i + 1)
		order.QuantityFrom = float64(i)
		order.OrderCost = 10 * (float64(i) + 0.5)
		order.OrderStatus = 0
		order.DateStart = strings.Split(time.Now().String(), ".")[0]
		ordersList = append(ordersList, order)
	}
	writeData, _ = ordersList.Marshal()
	err = ioutil.WriteFile("orders.json", writeData, 0666)
	if err != nil {
		log.Println(err)
	}
	log.Println("Stop create Orders JSON file...")
}
