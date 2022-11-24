package database

import (
	"database/sql"

	"github.com/shopspring/decimal"
	"github.com/ssergomol/E-Wallet/pkg/models"
)

type BalanceRepo struct {
	store *Storage
}

func (r *BalanceRepo) CreateBalance(balance models.Balance) {
	r.store.db.QueryRow(
		`INSERT INTO balances (sum) VALUES ($1)`,
		balance.Sum,
	)
}

func (r *BalanceRepo) GetBalance(userID uint) (models.Balance, error) {
	var sum string
	if err := r.store.db.QueryRow("SELECT sum from balances WHERE user_id = $1", userID).Scan(&sum); err != nil {
		if err == sql.ErrNoRows {
			r.store.db.QueryRow(
				`INSERT INTO balances (user_id, sum) VALUES ($1, $2)`,
				userID, "0.00",
			)
			balance := models.Balance{
				ID:  userID,
				Sum: "0.00",
			}
			return balance, nil
		}
		return models.Balance{}, err
	}

	balance := models.Balance{
		ID:  userID,
		Sum: sum,
	}
	return balance, nil
}

func (r *BalanceRepo) ReplenishBalance(balance models.Balance, order models.Order) error {
	var oldSum string
	if err := r.store.db.QueryRow("SELECT sum from balances WHERE user_id = $1", balance.ID).Scan(&oldSum); err != nil {
		if err == sql.ErrNoRows {
			r.store.db.QueryRow(
				`INSERT INTO balances (user_id, sum) VALUES ($1, $2)`,
				balance.ID, balance.Sum,
			)

			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	err := r.store.db.QueryRow("SELECT sum from balances WHERE user_id = $1", balance.ID).Scan(&oldSum)
	if err != nil {
		return err
	}

	sum, err := decimal.NewFromString(oldSum)
	if err != nil {
		return err
	}

	deltaSum, err := decimal.NewFromString(order.Price)
	if err != nil {
		return err
	}
	newSum := sum.Add(deltaSum)

	r.store.db.QueryRow("UPDATE balances SET sum = $1 WHERE user_id = $2", newSum.String(), order.UserID)
	return r.store.Order().CreateOrder(order)
}

func (r *BalanceRepo) SpendFunds(balance models.Balance, order models.Order) error {
	sum, err := decimal.NewFromString(balance.Sum)
	if err != nil {
		return err
	}

	deltaSum, err := decimal.NewFromString(order.Price)
	if err != nil {
		return err
	}

	r.store.db.QueryRow("UPDATE balances SET sum = $1 WHERE user_id = $2", sum.Sub(deltaSum), order.UserID)
	return r.store.Order().CreateOrder(order)
}
