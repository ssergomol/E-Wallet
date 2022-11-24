package database

import (
	"github.com/ssergomol/Balance-Manager/pkg/models"
)

type OrderRepo struct {
	store *Storage
}

func (r *OrderRepo) CreateOrder(order models.Order) error {
	_, err := r.store.db.Query("INSERT INTO orders(user_id, service_id, price, description)"+
		" VALUES ($1, $2, $3, $4)",
		order.UserID, order.ServiceID, order.Price, order.Description,
	)
	return err
}
