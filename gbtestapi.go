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

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// OrderCancel = -1 order cancelled, OrderWait = 0 order in queue, OrderComplete = 1 order finished
// const (
// 	OrderCancel = iota - 1
// 	OrderWait
// 	OrderComplete
// )

func main() {
	db, err := sqlx.Connect("sqlite3", "gelibert.db")
	// db, err := sqlx.Connect("mysql", "arthur:Nfnmzyf@tcp(217.12.127.253:3306)/gelibert")
	if err != nil {
		log.Println(err)
	}

	log.Println("Start create Couriers JSON file...")
	var couriersList Couriers
	for i := 0; i < 20; i++ {
		var courier Courier
		courier.ID = i
		courier.Imei = 123456789012300 + int64(i)
		courier.Tel = fmt.Sprintf("Tel_%d", i)
		courier.Name = fmt.Sprintf("Name_%d", i)
		courier.CarNumber = fmt.Sprintf("CarNumber_%d", i)
		courier.Latitude = 0
		courier.Longitude = 0
		courier.Address = fmt.Sprintf("Address_%d", i)
		couriersList = append(couriersList, courier)
	}
	for _, courier := range couriersList {
		_, err := db.NamedExec(`INSERT INTO couriers (id, imei, tel, name, car_number, latitude, longitude, address)
			VALUES (:id, :imei, :tel, :name, :car_number, :latitude, :longitude, :address)`, &courier)
		if err != nil {
			log.Println(err)
		}
	}
	writeData, _ := couriersList.Marshal()
	err = ioutil.WriteFile("couriers.json", writeData, 0666)
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
	for _, client := range clientsList {
		_, err := db.NamedExec(`INSERT INTO clients (id, name, tel, address)
		VALUES (:id, :name, :tel, :address)`, &client)
		if err != nil {
			log.Println(err)
		}
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
		var consist Consist
		order.ID = i
		order.CourierID = i
		order.ClientID = i
		order.PaymentMethod = "Cash"
		for j := 0; j < 4; j++ {
			if j%2 == 0 {
				consist.Delivery = true
			} else {
				consist.Delivery = false
			}
			if consist.Delivery {
				consist.Product = fmt.Sprintf("ProductTo_%d/%d", i, j)
				consist.Quantity = float64(j + 1)
			} else {
				consist.Product = fmt.Sprintf("ProductFrom_%d/%d", i, j)
				consist.Quantity = float64(j + 1)
			}
			order.Consists = append(order.Consists, consist)
		}
		order.OrderCost = 10 * (float64(i) + 0.5)
		order.Delivered = false
		order.DeliveryDelay = 0
		order.DateStart = strings.Split(time.Now().String(), ".")[0]
		order.DateFinish = ""
		// order.DateFinish = strings.Split(time.Now().String(), ".")[0]
		ordersList = append(ordersList, order)
	}
	for _, order := range ordersList {
		// _, err := db.NamedExec(`INSERT INTO orders (id, courier_id, client_id, payment_method, order_cost, delivered, delivery_delay, date_start)
		// 	 VALUES (:id, :courier_id, :client_id, :payment_method, :order_cost, :delivered, :delivery_delay, :date_start)`, &order)
		_, err := db.NamedExec(`INSERT INTO orders (id, courier_id, client_id, payment_method, order_cost, delivered, delivery_delay, date_start, date_finish)
			 VALUES (:id, :courier_id, :client_id, :payment_method, :order_cost, :delivered, :delivery_delay, :date_start, :date_finish)`, &order)
		if err != nil {
			log.Println(err)
		}
		for _, consist := range order.Consists {
			_, err := db.Exec(`INSERT INTO consists (id, product, quantity, delivery) 
			VALUES (?, ?, ?, ?)`, order.ID, consist.Product, consist.Quantity, consist.Delivery)
			if err != nil {
				log.Println(err)
			}
		}
	}
	writeData, _ = ordersList.Marshal()
	err = ioutil.WriteFile("orders.json", writeData, 0666)
	if err != nil {
		log.Println(err)
	}
	log.Println("Stop create Orders JSON file...")
}
