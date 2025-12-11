package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hse-football/internal/domain"
)

type teamRepo struct {
	db *pgxpool.Pool
}

func NewTeamRepo(db *pgxpool.Pool) domain.TeamRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) Create(ctx context.Context, e *domain.Team) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO Team (team_name, budget, club_id)
		 VALUES ($1, $2, $3) RETURNING team_id`,
		e.Name, e.Budget, e.ClubID,
	).Scan(&id)
	return id, err
}

func (r *teamRepo) GetByID(ctx context.Context, id int64) (*domain.Team, error) {
	var e domain.Team
	err := r.db.QueryRow(
		ctx,
		`SELECT team_id, team_name, budget, club_id
		 FROM Team WHERE team_id=$1`,
		id,
	).Scan(&e.ID, &e.Name, &e.Budget, &e.ClubID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *teamRepo) Update(ctx context.Context, e *domain.Team) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE Team
		 SET team_name=$1, budget=$2, club_id=$3
		 WHERE team_id=$4`,
		e.Name, e.Budget, e.ClubID, e.ID,
	)
	return err
}

func (r *teamRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM Team WHERE team_id=$1`, id)
	return err
}

func (r *teamRepo) List(ctx context.Context, limit, offset int) ([]*domain.Team, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT team_id, team_name, budget, club_id
		 FROM Team
		 ORDER BY team_id
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*domain.Team
	for rows.Next() {
		e := new(domain.Team)
		err := rows.Scan(&e.ID, &e.Name, &e.Budget, &e.ClubID)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, rows.Err()
}
