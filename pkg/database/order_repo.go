package database

import (
	"github.com/ssergomol/E-Wallet/pkg/models"
)

type OrderRepo struct {
	store *Storage
}

func (r *OrderRepo) CreateOrder(order models.Order) error {
	_, err := r.store.db.Query("INSERT INTO orders(user_id, service_id, price, description)"+
		" VALUES ($1, $2, $3, $4)",
		order.UserID, order.ServiceID, order.Price, order.Description,
	)

	if _, ok := r.store.Cache[order.UserID]; !ok {
		r.store.Cache[order.UserID], err = r.store.Order().RecoverCache(order.UserID)
		if err != nil {
			return err
		}
		return nil
	}

	r.store.Cache[order.UserID] = append(r.store.Cache[order.UserID], order)
	return nil
}

func (r *OrderRepo) RecoverCache(userID uint) ([]models.Order, error) {
	rows, err := r.store.db.Query("SELECT user_id, service_id, price, description FROM orders where user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]models.Order, 0)
	var order models.Order
	for rows.Next() {
		err := rows.Scan(&order.UserID, &order.ServiceID, &order.Price, &order.Description)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
