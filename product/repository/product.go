package repository

import (
	"context"

	"github.com/Ratatoouille/model"
	"github.com/jackc/pgx/v4"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r ProductRepository) CreateProduct(product *model.Product) error {
	_, err := r.db.Exec(
		context.Background(),
		"INSERT INTO products (model, company, price) VALUES ($1, $2, $3)",
		product.Model,
		product.Company,
		product.Price,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepository) GetProducts() ([]*model.Product, error) {
	var products []*model.Product

	rows, err := r.db.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := &model.Product{}

		err = rows.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	rows.Close()

	return products, nil
}

func (r ProductRepository) EditProduct(product *model.Product, id int) error {
	row := r.db.QueryRow(
		context.Background(),
		"SELECT id, model, company, price  FROM products WHERE id = $1", id,
	)

	err := row.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepository) UpdateProduct(product *model.Product) error {
	_, err := r.db.Exec(
		context.Background(),
		"UPDATE products SET model = $1, company = $2, price = $3 WHERE id = $4",
		product.Model,
		product.Company,
		product.Price,
		product.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepository) DeleteProduct(id int) error {
	_, err := r.db.Exec(
		context.Background(),
		"DELETE FROM products WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
