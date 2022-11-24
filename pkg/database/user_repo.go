package database

import "github.com/ssergomol/E-Wallet/pkg/models"

type UserRepo struct {
	store *Storage
}

func (r *UserRepo) CreateUser(user models.User) {
	r.store.db.Query("INSERT INTO (balance_id, order_id) VALUES ($1, $2)",
		user.BalanceID, user.OrderID,
	)
}
