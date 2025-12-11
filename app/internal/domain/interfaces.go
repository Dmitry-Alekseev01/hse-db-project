package domain

import "context"

type ClubRepository interface {
	Create(ctx context.Context, club *Club) (int64, error)
	GetByID(ctx context.Context, id int64) (*Club, error)
	Update(ctx context.Context, club *Club) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]*Club, error)
}
