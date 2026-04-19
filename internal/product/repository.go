package product

import (
	"context"

	productdb "belajar-go-be/internal/product/sqlc"
)

type Repository struct {
	q *productdb.Queries
}

func NewRepository(q *productdb.Queries) *Repository {
	return &Repository{q: q}
}

func (repo *Repository) FindAll(ctx context.Context) ([]Product, error) {
	rows, err := repo.q.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	var result []Product

	for _, row := range rows {
		result = append(result, Product{
			ID:          row.ID.String(),
			Name:        row.Name,
			Description: row.Description.String,
			Price:       row.Price,
			Stock:       int(row.Stock.Int32),
			ImageURL:    row.ImageUrl.String,
		})
	}

	return result, nil
}
