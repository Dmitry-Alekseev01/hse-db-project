package usecase

import (
	"context"
	"hse-football/internal/domain"
)

type playerUsecase struct {
	repo domain.PlayerRepository
}

func NewPlayerUsecase(r domain.PlayerRepository) domain.PlayerUsecase {
	return &playerUsecase{repo: r}
}

func (u *playerUsecase) Create(ctx context.Context, e *domain.Player) (int64, error) {
	return u.repo.Create(ctx, e)
}

func (u *playerUsecase) GetByID(ctx context.Context, id int64) (*domain.Player, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *playerUsecase) Update(ctx context.Context, e *domain.Player) error {
	return u.repo.Update(ctx, e)
}

func (u *playerUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *playerUsecase) List(ctx context.Context, limit, offset int) ([]*domain.Player, error) {
	return u.repo.List(ctx, limit, offset)
}
