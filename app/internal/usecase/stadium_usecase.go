package usecase

import (
	"context"
	"hse-football/internal/domain"
)

type stadiumUsecase struct {
	repo domain.StadiumRepository
}

func NewStadiumUsecase(r domain.StadiumRepository) domain.StadiumUsecase {
	return &stadiumUsecase{repo: r}
}

func (u *stadiumUsecase) Create(ctx context.Context, e *domain.Stadium) (int64, error) {
	return u.repo.Create(ctx, e)
}

func (u *stadiumUsecase) GetByID(ctx context.Context, id int64) (*domain.Stadium, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *stadiumUsecase) Update(ctx context.Context, e *domain.Stadium) error {
	return u.repo.Update(ctx, e)
}

func (u *stadiumUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *stadiumUsecase) List(ctx context.Context, limit, offset int) ([]*domain.Stadium, error) {
	return u.repo.List(ctx, limit, offset)
}
