package repository

import (
	productdb "belajar-go-be/internal/product/adapter/repository/sqlc"
	"belajar-go-be/internal/product/domain"
	"context"
	"database/sql"
)

type productRepository struct {
	q *productdb.Queries
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{
		q: productdb.New(db),
	}
}

func (r *productRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.q.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	var result []domain.Product

	for _, row := range rows {
		desc := ""
		if row.Description.Valid {
			desc = row.Description.String
		}

		image := ""
		if row.ImageUrl.Valid {
			image = row.ImageUrl.String
		}

		stock := 0
		if row.Stock.Valid {
			stock = int(row.Stock.Int32)
		}

		result = append(result, domain.Product{
			ID:          row.ID.String(),
			Name:        row.Name,
			Description: desc,
			Price:       row.Price,
			Stock:       stock,
			ImageURL:    image,
		})
	}

	return result, nil
}
