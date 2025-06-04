package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type TransactionType struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewTransactionType(db *sql.DB) *TransactionType {
	return &TransactionType{db: db}
}

func (s *TransactionType) Create(name string, description string) (TransactionType, error) {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO transaction_type (id, name, description) VALUES (?, ?, ?)",
		id, name, description)

	if err != nil {
		return TransactionType{}, err
	}

	return TransactionType{ID: id, Name: name, Description: description}, nil
}

func (s *TransactionType) FindAll() ([]TransactionType, error) {
	rows, err := s.db.Query("Select id, name, description from transaction_type")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := []TransactionType{}
	for rows.Next() {
		var id, name, description string
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, TransactionType{ID: id, Name: name, Description: description})
	}

	return transactions, nil

}

func (s *TransactionType) FindById(id string) (TransactionType, error) {
	row := s.db.QueryRow("SELECT id, name, description FROM transaction_type WHERE id = ?", id)
	transactionType := TransactionType{}

	err := row.Scan(&transactionType.ID, &transactionType.Name, &transactionType.Description)

	if err != nil {
		return transactionType, err
	}

	return transactionType, nil
}
