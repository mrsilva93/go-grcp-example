package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	db                *sql.DB
	ID                string
	TransactionTypeId string
	UserId            string
	Date              time.Time
	Value             int32
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{db: db}
}

func (s *Transaction) Create(transactionTypeId string, userId string, value int32, date time.Time) (Transaction, error) {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO transactions (id, transaction_type_id, user_id, value, data) VALUES (?, ?, ?, ?, ?)",
		id, transactionTypeId, userId, value, date)

	if err != nil {
		return Transaction{}, err
	}

	return Transaction{ID: id, TransactionTypeId: transactionTypeId, UserId: userId, Value: value, Date: date}, nil
}

func (s *Transaction) FindAll(transactionTypeId *string) ([]Transaction, error) {
	query := "SELECT id, transaction_type_id, value, user_id, data FROM transactions WHERE 1=1"
	args := []interface{}{}

	if transactionTypeId != nil {
		query += " AND transactionTypeId = ?"
		args = append(args, *transactionTypeId)
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var t Transaction
		var dateStr string

		if err := rows.Scan(&t.ID, &t.TransactionTypeId, &t.Value, &t.UserId, &dateStr); err != nil {
			return nil, err
		}

		t.Date, err = time.Parse("2006-01-02 15:04:05-07:00", dateStr)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
