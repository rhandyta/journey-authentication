package repository

import (
	"context"
	"database/sql"
	"journey-user/model"
	"log"
)

type InMemoryUserRepository struct {
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}

func (rUser *InMemoryUserRepository) Get(ctx context.Context, tx *sql.Tx) []model.User {
	query := "SELECT * FROM users"

	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}
	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Firstname, &user.Lastname, &user.Age, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	return users
}

// func (rUser *InMemoryUserRepository) Registration(ctx context.Context, tx *sql.Tx) model.User {
// 	query := "INSERT INTO users"
// }
