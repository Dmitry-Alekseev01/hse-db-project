package usecase

import (
	"context"
	"hse-football/internal/domain"
)

type teamUsecase struct {
	repo domain.TeamRepository
}

func NewTeamUsecase(r domain.TeamRepository) domain.TeamUsecase {
	return &teamUsecase{repo: r}
}

func (u *teamUsecase) Create(ctx context.Context, e *domain.Team) (int64, error) {
	return u.repo.Create(ctx, e)
}

func (u *teamUsecase) GetByID(ctx context.Context, id int64) (*domain.Team, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *teamUsecase) Update(ctx context.Context, e *domain.Team) error {
	return u.repo.Update(ctx, e)
}

func (u *teamUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *teamUsecase) List(ctx context.Context, limit, offset int) ([]*domain.Team, error) {
	return u.repo.List(ctx, limit, offset)
}
