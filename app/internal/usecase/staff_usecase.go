package usecase

import (
	"context"
	"hse-football/internal/domain"
)

type staffUsecase struct {
	repo domain.StaffRepository
}

func NewStaffUsecase(r domain.StaffRepository) domain.StaffUsecase {
	return &staffUsecase{repo: r}
}

func (u *staffUsecase) Create(ctx context.Context, e *domain.Staff) (int64, error) {
	return u.repo.Create(ctx, e)
}

func (u *staffUsecase) GetByID(ctx context.Context, id int64) (*domain.Staff, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *staffUsecase) Update(ctx context.Context, e *domain.Staff) error {
	return u.repo.Update(ctx, e)
}

func (u *staffUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *staffUsecase) List(ctx context.Context, limit, offset int) ([]*domain.Staff, error) {
	return u.repo.List(ctx, limit, offset)
}
