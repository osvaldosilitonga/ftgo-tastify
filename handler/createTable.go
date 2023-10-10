package handler

import (
	"context"
	"database/sql"
	"time"
)

func CreateEmployees(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS employees (id INT PRIMARY KEY AUTO_INCREMENT, first_name VARCHAR(50), last_name VARCHAR(50), position VARCHAR(50) )")
	return err
}

func CreateMenuItems(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS menu_items (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100), description VARCHAR(255), price DECIMAL(10, 2), category VARCHAR(50) )")
	return err
}

func CreateOrders(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS orders (id INT PRIMARY KEY AUTO_INCREMENT, table_number INT, employee_id INT, order_date DATE, status VARCHAR(50), FOREIGN KEY (employee_id) REFERENCES employees(id) )")
	return err
}

func CreateOrderItems(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS order_items (id INT PRIMARY KEY AUTO_INCREMENT, order_id INT, item_id INT, quantity INT, subtotal DECIMAL(10, 2), FOREIGN KEY (order_id) REFERENCES orders(id), FOREIGN KEY (item_id) REFERENCES menu_items(id) )")
	return err
}

func CreatePayments(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS payments (id INT PRIMARY KEY AUTO_INCREMENT, order_id INT, payment_date DATE, payment_method VARCHAR(50), total_amount DECIMAL(10, 2), FOREIGN KEY (order_id) REFERENCES orders(id) )")
	return err
}
