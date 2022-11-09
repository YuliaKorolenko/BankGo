package base

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgresMethods struct {
	Db *sqlx.DB
}

func (r *PostgresMethods) Create(userId int, amount int) {
	query := fmt.Sprintf("INSERT INTO balances (id, balance) values ($1, $2)")
	r.Db.QueryRow(query, userId, amount)
	fmt.Sprintf("hi, %s, %t", userId, amount)
}

func (r *PostgresMethods) PutBalance(userId int, amount int) {
	query := fmt.Sprintf("UPDATE balances SET balance=$2 WHERE id=$1")
	r.Db.QueryRow(query, userId, amount)
	fmt.Sprintf("hi, %s, %t", userId, amount)
}

func (r *PostgresMethods) GetBalance(userId int) (int, error) {
	var balance int
	query := fmt.Sprintf("SELECT balance FROM balances WHERE id=$1")
	err := r.Db.Get(&balance, query, userId)

	return balance, err
}

func (r *PostgresMethods) PutTransaction(transaction Transaction) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO transactions (id_order, id_user, id_service, amount, flag) values (DEFAULT, $2, $3, $4, $5)")

	row := r.Db.QueryRow(query, transaction.Id, transaction.IdService, transaction.Amount, transaction.Flag)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
