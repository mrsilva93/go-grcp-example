package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	db   *sql.DB
	ID   string
	Name string
	Cnpj string
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (s *User) Create(name string, cnpj string) (User, error) {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO users (id, name, cnpj) VALUES (?, ?, ?)",
		id, name, cnpj)

	if err != nil {
		return User{}, err
	}

	return User{ID: id, Name: name, Cnpj: cnpj}, nil
}

func (s *User) FindAll() ([]User, error) {
	rows, err := s.db.Query("Select id, name, cnpj from users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := []User{}
	for rows.Next() {
		var id, name, cnpj string
		err := rows.Scan(&id, &name, &cnpj)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, User{ID: id, Name: name, Cnpj: cnpj})
	}

	return transactions, nil

}

func (s *User) FindById(id string) (User, error) {
	row := s.db.QueryRow("SELECT id, name, cnpj FROM users WHERE id = ?", id)
	user := User{}

	err := row.Scan(&user.ID, &user.Name, &user.Cnpj)

	if err != nil {
		return user, err
	}

	return user, nil

}
