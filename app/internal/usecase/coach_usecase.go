package usecase

import (
	"context"
	"hse-football/internal/domain"
)

type coachUsecase struct {
	repo domain.CoachRepository
}

func NewCoachUsecase(r domain.CoachRepository) domain.CoachUsecase {
	return &coachUsecase{repo: r}
}

func (u *coachUsecase) Create(ctx context.Context, e *domain.Coach) (int64, error) {
	return u.repo.Create(ctx, e)
}

func (u *coachUsecase) GetByID(ctx context.Context, id int64) (*domain.Coach, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *coachUsecase) Update(ctx context.Context, e *domain.Coach) error {
	return u.repo.Update(ctx, e)
}

func (u *coachUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *coachUsecase) List(ctx context.Context, limit, offset int) ([]*domain.Coach, error) {
	return u.repo.List(ctx, limit, offset)
}
