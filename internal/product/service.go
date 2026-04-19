package product

import "context"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAll(ctx context.Context) ([]Product, error) {
	return s.repo.FindAll(ctx)
}
