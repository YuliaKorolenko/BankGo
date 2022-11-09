package base

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgresMethods struct {
	Db *sqlx.DB
}

//balanceTable

func (r *PostgresMethods) CreateBalance(userId int, amount int) *sql.Row {
	query := fmt.Sprintf("INSERT INTO balances (id, balance) values ($1, $2)")
	err := r.Db.QueryRow(query, userId, amount)
	return err
}

func (r *PostgresMethods) PutBalance(userId int, amount int) *sql.Row {
	query := fmt.Sprintf("UPDATE balances SET balance=$2 WHERE id=$1")
	err := r.Db.QueryRow(query, userId, amount)
	return err
}

func (r *PostgresMethods) GetBalance(userId int) (int, error) {
	var balance int
	query := fmt.Sprintf("SELECT balance FROM balances WHERE id=$1")
	err := r.Db.Get(&balance, query, userId)
	return balance, err
}

func (r *PostgresMethods) IsExistBalance(userId int) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM balances WHERE (id=$1)")
	row := r.Db.QueryRow(query, userId)
	var tmp interface{}
	err := row.Scan(&tmp)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

//transactions

func (r *PostgresMethods) PutTransaction(transaction Transaction) *sql.Row {
	query := fmt.Sprintf(
		"INSERT INTO transactions (id_order, id_user, id_service, amount, is_debit) values ($1, $2, $3, $4, $5)")

	err := r.Db.QueryRow(query, transaction.IdOrder, transaction.Id, transaction.IdService, transaction.Amount, false)
	return err
}

func (r *PostgresMethods) CompletionTransaction(transaction Transaction) *sql.Row {
	query := fmt.Sprintf(
		"UPDATE transactions SET is_debit=true WHERE (id_order=$1, id_user=$2, id_service=$3, amount=$4)")
	err := r.Db.QueryRow(query, transaction.IdOrder, transaction.Id, transaction.IdService, transaction.Amount)
	return err
}

func (r *PostgresMethods) IsExistReserveTransaction(transaction Transaction) (bool, error) {
	query := fmt.Sprintf("SELECT id_user FROM transactions WHERE (id_order=$1)")
	row := r.Db.QueryRow(query, transaction.IdOrder)
	var tmp interface{}
	err := row.Scan(&tmp)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err

}

// charges

func (r *PostgresMethods) PutCharge(userId int, amount int) *sql.Row {
	query := fmt.Sprintf("INSERT INTO charges (id_user, amount) values ($1, $2)")
	err := r.Db.QueryRow(query, userId, amount)
	return err
}
