package base

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgresMethods struct {
	Db *sqlx.DB
}

func (r *PostgresMethods) CreateBalance(userId int, amount int) {
	query := fmt.Sprintf("INSERT INTO balances (id, balance) values ($1, $2)")
	r.Db.QueryRow(query, userId, amount)
}

func (r *PostgresMethods) PutBalance(userId int, amount int) {
	query := fmt.Sprintf("UPDATE balances SET balance=$2 WHERE id=$1")
	r.Db.QueryRow(query, userId, amount)
}

func (r *PostgresMethods) GetBalance(userId int) (int, error) {
	var balance int
	query := fmt.Sprintf("SELECT balance FROM balances WHERE id=$1")
	err := r.Db.Get(&balance, query, userId)

	return balance, err
}

func (r *PostgresMethods) PutTransaction(transaction Transaction) {
	query := fmt.Sprintf(
		"INSERT INTO transactions (id_order, id_user, id_service, amount, is_debit) values ($1, $2, $3, $4, $5)")

	r.Db.QueryRow(query, transaction.IdOrder, transaction.Id, transaction.IdService, transaction.Amount, true)
}

func (r *PostgresMethods) CompletionTransaction(transaction Transaction) {
	query := fmt.Sprintf(
		"UPDATE transactions SET is_debit=true WHERE (id_order=$1, id_user=$2, id_service=$3, amount=$4)")
	r.Db.QueryRow(query, transaction.IdOrder, transaction.Id, transaction.IdService, transaction.Amount)
}

func (r *PostgresMethods) PutCharge(userId int, amount int) {
	query := fmt.Sprintf("INSERT INTO charges (id_user, amount) values ($1, $2)")
	r.Db.QueryRow(query, userId, amount)
}
