package repository

import (
	"TechShop/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user models.User) (string, error) {
	var id string

	query := fmt.Sprintf("INSERT INTO %s (name,username,email,password_hash) values ($1,$2,$3,$4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}
func (r *AuthPostgres) GetUser(username, password, email string) (models.User, error) {
	var user models.User
	fmt.Println(username, " ", password, " ", email)
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2 AND email=$3", usersTable)
	err := r.db.Get(&user, query, username, password, email)
	return user, err
}
