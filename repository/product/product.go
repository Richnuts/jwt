package product

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type ProductRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) GetProduct() ([]entities.Product, error) {
	var products []entities.Product
	result, err := pr.db.Query("select products.id, products.price, users.id, users.name, users.email, users.password from products JOIN users ON products.user_id = users.id")
	if err != nil {
		return products, err
	}
	defer result.Close()
	for result.Next() {
		var product entities.Product
		err := result.Scan(&product.Id, &product.Price, &product.User.Id, &product.User.Name, &product.User.Email, &product.User.Password)
		if err != nil {
			return products, fmt.Errorf("product not found")
		}
		products = append(products, product)
	}
	return products, nil
}
