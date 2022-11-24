package models

type User struct {
	ID        uint `json:"id"`
	BalanceID uint `json:"balance_id"`
	OrderID   uint `json:"order_id"`
}
