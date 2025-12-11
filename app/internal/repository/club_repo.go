package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hse-football/internal/domain"
)

type clubRepo struct {
	db *pgxpool.Pool
}

func NewClubRepo(db *pgxpool.Pool) domain.ClubRepository {
	return &clubRepo{db: db}
}

func (r *clubRepo) Create(ctx context.Context, club *domain.Club) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO Club (club_name, creation_date, website)
		 VALUES ($1, $2, $3) RETURNING club_id`,
		club.Name, club.CreationDate, club.Website,
	).Scan(&id)

	return id, err
}

func (r *clubRepo) GetByID(ctx context.Context, id int64) (*domain.Club, error) {
	var c domain.Club

	err := r.db.QueryRow(
		ctx,
		`SELECT club_id, club_name, creation_date, website
		 FROM Club WHERE club_id=$1`,
		id,
	).Scan(&c.ID, &c.Name, &c.CreationDate, &c.Website)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *clubRepo) Update(ctx context.Context, club *domain.Club) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE Club
		 SET club_name=$1, creation_date=$2, website=$3
		 WHERE club_id=$4`,
		club.Name, club.CreationDate, club.Website, club.ID,
	)

	return err
}

func (r *clubRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM Club WHERE club_id=$1`, id)
	return err
}

func (r *clubRepo) List(ctx context.Context, limit, offset int) ([]*domain.Club, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT club_id, club_name, creation_date, website
		 FROM Club
		 ORDER BY club_id
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*domain.Club

	for rows.Next() {
		c := new(domain.Club)
		err := rows.Scan(&c.ID, &c.Name, &c.CreationDate, &c.Website)
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, rows.Err()
}
