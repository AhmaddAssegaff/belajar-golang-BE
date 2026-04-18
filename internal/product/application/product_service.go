package application

import (
	"context"

	"belajar-go-be/internal/product/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(r domain.ProductRepository) *ProductService {
	return &ProductService{
		repo: r,
	}
}

func (s *ProductService) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repo.FindAll(ctx)
}
