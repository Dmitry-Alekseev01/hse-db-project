package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hse-football/internal/domain"
)

type stadiumRepo struct {
	db *pgxpool.Pool
}

func NewStadiumRepo(db *pgxpool.Pool) domain.StadiumRepository {
	return &stadiumRepo{db: db}
}

func (r *stadiumRepo) Create(ctx context.Context, e *domain.Stadium) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO Stadium (capacity, stadium_location, build_date)
		 VALUES ($1, $2, $3) RETURNING stadium_id`,
		e.Capacity, e.Location, e.BuildDate,
	).Scan(&id)
	return id, err
}

func (r *stadiumRepo) GetByID(ctx context.Context, id int64) (*domain.Stadium, error) {
	var e domain.Stadium
	err := r.db.QueryRow(
		ctx,
		`SELECT stadium_id, capacity, stadium_location, build_date
		 FROM Stadium WHERE stadium_id=$1`,
		id,
	).Scan(&e.ID, &e.Capacity, &e.Location, &e.BuildDate)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *stadiumRepo) Update(ctx context.Context, e *domain.Stadium) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE Stadium
		 SET capacity=$1, stadium_location=$2, build_date=$3
		 WHERE stadium_id=$4`,
		e.Capacity, e.Location, e.BuildDate, e.ID,
	)
	return err
}

func (r *stadiumRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM Stadium WHERE stadium_id=$1`, id)
	return err
}

func (r *stadiumRepo) List(ctx context.Context, limit, offset int) ([]*domain.Stadium, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT stadium_id, capacity, stadium_location, build_date
		 FROM Stadium
		 ORDER BY stadium_id
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*domain.Stadium
	for rows.Next() {
		e := new(domain.Stadium)
		err := rows.Scan(&e.ID, &e.Capacity, &e.Location, &e.BuildDate)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, rows.Err()
}
