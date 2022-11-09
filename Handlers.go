package main

import (
	"awesomeProject/base"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type Handler struct {
	postgresMethods base.PostgresMethods
}

func (handler *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var transaction base.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	handler.postgresMethods.Create(transaction.Id, transaction.Amount)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// DepositMoney Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.
func (handler *Handler) DepositMoney(w http.ResponseWriter, r *http.Request) {
	var transaction base.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	balance, _ := handler.postgresMethods.GetBalance(transaction.Id)
	handler.postgresMethods.PutBalance(transaction.Id, balance+transaction.Amount)
	fmt.Fprintf(w, "Deposit: %+v", balance+transaction.Amount)
}

// Reserve Метод резервирования средств с основного баланса на отдельном счете.
// Принимает id пользователя, ИД услуги, ИД заказа, стоимость.
func (handler *Handler) Reserve(w http.ResponseWriter, r *http.Request) {
	var transaction base.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	balance, err := handler.postgresMethods.GetBalance(transaction.Id)
	if balance-transaction.Amount < 0 {
		http.Error(w, "No money", http.StatusBadRequest)
		return
	}
	handler.postgresMethods.PutBalance(transaction.Id, balance-transaction.Amount)
	handler.postgresMethods.
		PutTransaction(
			base.Transaction{transaction.Id, transaction.IdService,
				transaction.IdOrder, transaction.Amount, 1})
	fmt.Fprintf(w, "Welcome home!")
}

// PersonalBalance Метод получения баланса пользователя. Принимает id пользователя.
func (handler *Handler) PersonalBalance(w http.ResponseWriter, r *http.Request) {
	var transaction base.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	balance, err := handler.postgresMethods.GetBalance(transaction.Id)
	fmt.Fprintf(w, "PersonalBalance: %+v", balance)
}

// Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму.
func (handler *Handler) Debit(w http.ResponseWriter, r *http.Request) {
	var transaction base.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	handler.postgresMethods.PutTransaction(transaction)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
