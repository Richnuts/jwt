package graph

import (
	"sirclo/repository/auth"
	"sirclo/repository/book"
	"sirclo/repository/product"
	_userRepo "sirclo/repository/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authRepo    auth.Auth
	userRepo    _userRepo.User
	bookRepo    book.Book
	productRepo product.Product
}

func NewResolver(ur _userRepo.User, br book.Book, pr product.Product, ar auth.Auth) *Resolver {
	return &Resolver{
		authRepo:    ar,
		userRepo:    ur,
		bookRepo:    br,
		productRepo: pr,
	}
}
