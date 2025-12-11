package usecase

import (
	"context"
	"hse-football/internal/domain"
)

type gameUsecase struct {
	repo domain.GameRepository
}

func NewGameUsecase(r domain.GameRepository) domain.GameUsecase {
	return &gameUsecase{repo: r}
}

func (u *gameUsecase) Create(ctx context.Context, e *domain.Game) (int64, error) {
	return u.repo.Create(ctx, e)
}

func (u *gameUsecase) GetByID(ctx context.Context, id int64) (*domain.Game, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *gameUsecase) Update(ctx context.Context, e *domain.Game) error {
	return u.repo.Update(ctx, e)
}

func (u *gameUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *gameUsecase) List(ctx context.Context, limit, offset int) ([]*domain.Game, error) {
	return u.repo.List(ctx, limit, offset)
}
