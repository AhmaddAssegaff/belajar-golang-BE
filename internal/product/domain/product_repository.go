package domain

import "context"

type ProductRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
}
