package user

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result, err := ur.db.Query("select id, name, email, password from users")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var user entities.User
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("user not found")
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	result, err := ur.db.Exec("delete from users where id = ?", id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (ur *UserRepository) EditUser(user entities.User, id int) error {
	result, err := ur.db.Exec("UPDATE users SET name= ?, email= ?, password= ? WHERE id = ?", user.Name, user.Email, user.Password, id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (ur *UserRepository) CreateUser(user entities.User) error {
	result, err := ur.db.Exec("INSERT INTO users(name, email, password) VALUES(?,?,?)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not created")
	}
	return nil
}
