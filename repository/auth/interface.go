package auth

import "sirclo/entities"

type Auth interface {
	Login(email string, password string) (string, entities.User, error)
}
