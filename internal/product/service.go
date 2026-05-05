package product

import (
	"context"
	"errors"
)

type RepositoryInterface interface {
	FindAll(ctx context.Context) ([]Product, error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAll(ctx context.Context) ([]Product, error) {
	products, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("no products found")
	}

	return products, nil
}
