package handler

import (
	"context"
	"database/sql"
	"time"
)

func InsertEmployees(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO employees(first_name, last_name, position) VALUES ('John', 'Doe', 'Manager'), ('Jane', 'Smith', 'Waitres'), ('Robert', 'Brown', 'Cook') ")
	return err
}

func InsertMenuItems(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO menu_items(name, description, price, category) VALUES ('Spaghetti Carbonara', 'Traditional Italian dish with eggs, cheese, pancetta, and pepper', 12.50, 'Main Course'), ('Caesar Salad', 'Fresh lettuce with Caesar dressing, croutons and parmesan', 6.00, 'Starter'), ('Tiramisu', 'Classic Italian desert with coffee-soaked spone and mascarpone', 5.50, 'Desert')")
	return err
}

func InsertOrders(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO orders(table_number, employee_id, order_date, status) VALUES (10, 1, '2023-08-09', 'Placed'), (5, 2, '2023-08-09', 'Completed')")
	return err
}

func InsertOrderItems(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO order_items(order_id, item_id, quantity, subtotal) VALUES (1, 1, 2, 25.00), (2, 2, 1, 6.00), (2, 3, 1, 5.50)")
	return err
}

func InsertPayments(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO payments(order_id, payment_date, payment_method, total_amount) VALUES (1, '2023-08-09', 'Credit Card', 25.00), (2, '2023-08-09', 'Cash', 11.50) ")
	return err
}
