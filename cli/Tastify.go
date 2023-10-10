package cli

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/osvaldosilitonga/ftgo-tastify/config"
	"github.com/osvaldosilitonga/ftgo-tastify/handler"
)

func connDB(connURL string) *sql.DB {
	db, err := config.GetDb(connURL)
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}

	return db
}

func OrderDetails(connURL string) {
	db := connDB(connURL)
	defer db.Close()

	orders, err := handler.GetOrderDetails(db)
	if err != nil {
		log.Fatal("Failed to retrive order details data :", err)
	}

	for _, o := range orders {
		fmt.Println(o.Id, o.FirstName, o.LastName, o.Position, o.TableNumber, o.Status)
	}
}

func CreateTable(connURL string) {
	db := connDB(connURL)
	defer db.Close()

	// Create Employee Table
	err := handler.CreateEmployees(db)
	if err != nil {
		log.Fatal("Failed to create Employee Table")
	} else {
		fmt.Println("Success create table employee")
	}

	// Create Menu Items Table
	err = handler.CreateMenuItems(db)
	if err != nil {
		log.Fatal("Failed to create Menu Items Table")
	} else {
		fmt.Println("Success create table employee")
	}

	// Create Orders Table
	err = handler.CreateOrders(db)
	if err != nil {
		log.Fatal("Failed to create Orders Table")
	} else {
		fmt.Println("Success create table orders")
	}

	// Create Order Items Table
	err = handler.CreateOrderItems(db)
	if err != nil {
		log.Fatal("Failed to create Order Items Table")
	} else {
		fmt.Println("Success create table order items")
	}

	// Create Payments Table
	err = handler.CreatePayments(db)
	if err != nil {
		log.Fatal("Failed to create Payments Table")
	} else {
		fmt.Println("Success create table payments")
	}
}

func InsertData(connURL string) {
	db := connDB(connURL)
	defer db.Close()

	// Insert employees
	err := handler.InsertEmployees(db)
	if err != nil {
		log.Fatal("Failed insert data to employees")
	} else {
		fmt.Println("Inserting employess data success")
	}

	// Insert menu items
	err = handler.InsertMenuItems(db)
	if err != nil {
		log.Fatal("Failed insert data to menu items")
	} else {
		fmt.Println("Inserting menu items data success")
	}

	// Insert orders
	err = handler.InsertOrders(db)
	if err != nil {
		log.Fatal("Failed insert data to orders")
	} else {
		fmt.Println("Inserting orders data success")
	}

	// Insert order items
	err = handler.InsertOrderItems(db)
	if err != nil {
		log.Fatal("Failed insert data to order items")
	} else {
		fmt.Println("Inserting order items data success")
	}

	// Insert payments
	err = handler.InsertPayments(db)
	if err != nil {
		log.Fatal("Failed insert data to payments")
	} else {
		fmt.Println("Inserting payments data success")
	}
}
