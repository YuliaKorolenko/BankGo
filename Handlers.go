package main

import (
	"awesomeProject/base"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	postgresMethods base.PostgresMethods
}

func MakeTransaction(w http.ResponseWriter, r *http.Request) base.Transaction {
	var transaction base.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return base.Transaction{}
	}
	return transaction
}

func (handler *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	transaction := MakeTransaction(w, r)
	if (base.Transaction{}) == transaction {
		return
	}
	isAccount, _ := handler.postgresMethods.IsExistBalance(transaction.Id)
	if isAccount {
		http.Error(w, "Account already exist", http.StatusBadRequest)
		return
	}
	handler.postgresMethods.CreateBalance(transaction.Id, transaction.Amount)
	handler.postgresMethods.PutCharge(transaction.Id, transaction.Amount)
	fmt.Fprintf(w, "Create Account : %+v", transaction.Id)
}

// DepositMoney Метод начисления средств на баланс.
func (handler *Handler) DepositMoney(w http.ResponseWriter, r *http.Request) {
	transaction := MakeTransaction(w, r)
	balance, _ := handler.postgresMethods.GetBalance(transaction.Id)
	handler.postgresMethods.PutBalance(transaction.Id, balance+transaction.Amount)
	handler.postgresMethods.PutCharge(transaction.Id, transaction.Amount)
	fmt.Fprintf(w, "Deposit : %+v", balance+transaction.Amount)
}

// Reserve Метод резервирования средств с основного баланса на отдельном счете.
func (handler *Handler) Reserve(w http.ResponseWriter, r *http.Request) {
	transaction := MakeTransaction(w, r)
	isReserve, _ := handler.postgresMethods.IsExistReserveTransaction(transaction)
	if isReserve {
		http.Error(w, "Money for this order has been reserved already", http.StatusBadRequest)
		return
	}
	balance, err := handler.postgresMethods.GetBalance(transaction.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if balance-transaction.Amount < 0 {
		http.Error(w, "No money", http.StatusBadRequest)
		return
	}
	handler.postgresMethods.PutBalance(transaction.Id, balance-transaction.Amount)
	handler.postgresMethods.PutTransaction(transaction)
	fmt.Fprintf(w, "Reserve : %+v", transaction)
}

// PersonalBalance Метод получения баланса пользователя.
func (handler *Handler) PersonalBalance(w http.ResponseWriter, r *http.Request) {
	transaction := MakeTransaction(w, r)
	balance, err := handler.postgresMethods.GetBalance(transaction.Id)
	if err != nil {
		http.Error(w, "No money", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "PersonalBalance: %+v", balance)
}

// Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии.
func (handler *Handler) Debit(w http.ResponseWriter, r *http.Request) {
	transaction := MakeTransaction(w, r)
	isReserve, _ := handler.postgresMethods.IsExistReserveTransaction(transaction)
	if !isReserve {
		http.Error(w, "Money for this order has not been reserved", http.StatusBadRequest)
		return
	}
	handler.postgresMethods.CompletionTransaction(transaction)
	fmt.Fprintf(w, "Debit : %+v", transaction)
}
