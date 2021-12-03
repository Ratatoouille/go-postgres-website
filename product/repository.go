package product

import (
	"github.com/Ratatoouille/model"
)

type Repository interface {
	CreateProduct(product *model.Product) error
	GetProducts() ([]*model.Product, error)
	EditProduct(product *model.Product, id int) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int) error
}
