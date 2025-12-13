package usecase

import (
	"context"
	"hse-football/internal/domain"
)

type clubUsecase struct {
	repo domain.ClubRepository
}

func NewClubUsecase(r domain.ClubRepository) domain.ClubUsecase {
	return &clubUsecase{repo: r}
}

func (u *clubUsecase) Create(ctx context.Context, club *domain.Club) (int64, error) {
	return u.repo.Create(ctx, club)
}

func (u *clubUsecase) GetByID(ctx context.Context, id int64) (*domain.Club, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *clubUsecase) Update(ctx context.Context, club *domain.Club) error {
	return u.repo.Update(ctx, club)
}

func (u *clubUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *clubUsecase) List(ctx context.Context, limit, offset int) ([]*domain.Club, error) {
	return u.repo.List(ctx, limit, offset)
}
