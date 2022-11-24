package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/shopspring/decimal"
	"github.com/ssergomol/E-Wallet/pkg/models"
)

func (s *APIserver) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello! This is the E-Wallet</h1>"))
	}
}

func (s *APIserver) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.logger.Info("/balance POST request")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.logger.Fatal(err)
		}

		order := models.Order{}

		err = json.Unmarshal([]byte(body), &order)
		if err != nil {
			s.logger.Fatal(err)
		}

		balance, err := s.db.Balance().GetBalance(order.UserID)
		if err != nil {
			s.logger.Fatal(err)
		}

		switch order.ServiceID {
		case 0:
			order.Description = "Replenish balance"
			err = s.db.Balance().ReplenishBalance(balance, order)

		case 1:
			sum, err := decimal.NewFromString(balance.Sum)
			if err != nil {
				s.logger.Fatal(err)
			}

			price, err := decimal.NewFromString(order.Price)
			if err != nil {
				s.logger.Fatal(err)
			}

			if sum.LessThan(price) {
				w.WriteHeader(http.StatusBadRequest)
				message, err := json.Marshal("Not enough funds on the balance")
				if err != nil {
					s.logger.Fatal(err)
				}

				w.Write(message)
				return
			}

			order.Description = "Spend funds"
			err = s.db.Balance().SpendFunds(balance, order)
		}
		if err != nil {
			s.logger.Fatal(err)
		}

		var ordersToSend []models.Order
		if orders, ok := s.db.Cache[order.UserID]; !ok {
			ordersToSend, err = s.db.Order().RecoverCache(order.UserID)
			if err != nil {
				s.logger.Fatal(err)
			}
		} else {
			ordersToSend = orders
		}

		bytes, err := json.Marshal(ordersToSend)
		if err != nil {
			s.logger.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)

	case http.MethodGet:
		s.logger.Info("/balance get request")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.logger.Fatal(err)
		}

		balance := models.Balance{}

		err = json.Unmarshal([]byte(body), &balance)
		if err != nil {
			s.logger.Fatal(err)
		}

		getBalance, err := s.db.Balance().GetBalance(balance.ID)
		if err != nil {
			s.logger.Fatal(err)
		}

		bytes, err := json.Marshal(getBalance)
		if err != nil {
			s.logger.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
