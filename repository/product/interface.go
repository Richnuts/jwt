package product

import (
	"sirclo/entities"
)

type Product interface {
	GetProduct() ([]entities.Product, error)
}
