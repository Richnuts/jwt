package main

import (
	"fmt"
	"sirclo/config"

	"sirclo/delivery/controllers/graph"
	"sirclo/delivery/router"
	_bookRepo "sirclo/repository/book"
	_productRepo "sirclo/repository/product"
	_userRepo "sirclo/repository/user"
	_authRepo "sirclo/repository/auth"
	"sirclo/util"

	"github.com/labstack/echo/v4"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDriver(config)

	//initiate user model
	// authRepo := auth.New()
	userRepo := _userRepo.New(db)
	bookRepo := _bookRepo.New(db)
	productRepo := _productRepo.New(db)
	authRepo := _authRepo.New(db)
	//create echo http
	e := echo.New()
	client := graph.NewResolver(userRepo, bookRepo, productRepo, authRepo)
	srv := router.NewGraphQLServer(client)
	router.RegisterPath(e, srv)

	// run server
	address := fmt.Sprintf("localhost:%d", config.Port)
	if err := e.Start(address); err != nil {
		fmt.Errorf("shutting down the server")
	}
}
