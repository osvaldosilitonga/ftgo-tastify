package handler

import (
	"context"
	"database/sql"
	"time"

	"github.com/osvaldosilitonga/ftgo-tastify/entity"
)

func GetOrderDetails(db *sql.DB) ([]entity.OrdersDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT orders.id, employees.first_name, employees.last_name, employees.position, orders.table_number, orders.status FROM orders JOIN employees ON orders.employee_id = employees.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.OrdersDetail

	for rows.Next() {
		var o entity.OrdersDetail

		err := rows.Scan(&o.Id, &o.FirstName, &o.LastName, &o.Position, &o.TableNumber, &o.Status)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, nil
}
