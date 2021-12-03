package usecase

import (
	"github.com/Ratatoouille/model"
	"github.com/Ratatoouille/product"
)

type ProductUseCase struct {
	productRepo product.Repository
}

func NewProductUseCase(productRepo product.Repository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

func (p ProductUseCase) CreateProduct(product *model.Product) error {
	return p.productRepo.CreateProduct(product)
}

func (p ProductUseCase) GetProducts() ([]*model.Product, error) {
	return p.productRepo.GetProducts()
}

func (p ProductUseCase) EditProduct(product *model.Product, id int) error {
	return p.productRepo.EditProduct(product, id)
}

func (p ProductUseCase) UpdateProduct(product *model.Product) error {
	return p.productRepo.UpdateProduct(product)
}

func (p ProductUseCase) DeleteProduct(id int) error {
	return p.productRepo.DeleteProduct(id)
}
