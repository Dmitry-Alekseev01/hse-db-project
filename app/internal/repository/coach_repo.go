package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hse-football/internal/domain"
)

type coachRepo struct {
	db *pgxpool.Pool
}

func NewCoachRepo(db *pgxpool.Pool) domain.CoachRepository {
	return &coachRepo{db: db}
}

func (r *coachRepo) Create(ctx context.Context, e *domain.Coach) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO Coach (coach_name, coach_surname, salary, phone, team_id)
		 VALUES ($1, $2, $3, $4, $5) RETURNING coach_id`,
		e.Name, e.Surname, e.Salary, e.Phone, e.TeamID,
	).Scan(&id)
	return id, err
}

func (r *coachRepo) GetByID(ctx context.Context, id int64) (*domain.Coach, error) {
	var e domain.Coach
	err := r.db.QueryRow(
		ctx,
		`SELECT coach_id, coach_name, coach_surname, salary, phone, team_id
		 FROM Coach WHERE coach_id=$1`,
		id,
	).Scan(&e.ID, &e.Name, &e.Surname, &e.Salary, &e.Phone, &e.TeamID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *coachRepo) Update(ctx context.Context, e *domain.Coach) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE Coach
		 SET coach_name=$1, coach_surname=$2, salary=$3, phone=$4, team_id=$5
		 WHERE coach_id=$6`,
		e.Name, e.Surname, e.Salary, e.Phone, e.TeamID, e.ID,
	)
	return err
}

func (r *coachRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM Coach WHERE coach_id=$1`, id)
	return err
}

func (r *coachRepo) List(ctx context.Context, limit, offset int) ([]*domain.Coach, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT coach_id, coach_name, coach_surname, salary, phone, team_id
		 FROM Coach
		 ORDER BY coach_id
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*domain.Coach
	for rows.Next() {
		e := new(domain.Coach)
		err := rows.Scan(&e.ID, &e.Name, &e.Surname, &e.Salary, &e.Phone, &e.TeamID)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, rows.Err()
}
