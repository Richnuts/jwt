package auth

import (
	"database/sql"
	"fmt"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
)

type AuthRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) Login(email string, password string) (string, entities.User, error) {
	var user entities.User
	// input email = asd, password = 123
	result, err := a.db.Query("select * from users where email = ? and password = ?", email, password)
	if err != nil {
		fmt.Println(err)
		return "", user, err
	}
	for result.Next() {
		err_scan := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err_scan != nil {
			return "", user, err_scan
		}
	}
	if user.Email == email && user.Password == password {
		token, err_token := middlewares.CreateToken(user.Id)
		if err_token != nil {
			return "", user, err
		}
		return token, user, nil
	}
	// tidak error tapi usernya tidak ada
	return "", user, fmt.Errorf("user not found")
}
