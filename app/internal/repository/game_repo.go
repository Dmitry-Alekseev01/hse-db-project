package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hse-football/internal/domain"
)

type gameRepo struct {
	db *pgxpool.Pool
}

func NewGameRepo(db *pgxpool.Pool) domain.GameRepository {
	return &gameRepo{db: db}
}

func (r *gameRepo) Create(ctx context.Context, e *domain.Game) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO Game (stadium_id, team_1_id, team_2_id, match_date)
		 VALUES ($1, $2, $3, $4) RETURNING match_id`,
		e.StadiumID, e.Team1ID, e.Team2ID, e.MatchDate,
	).Scan(&id)
	return id, err
}

func (r *gameRepo) GetByID(ctx context.Context, id int64) (*domain.Game, error) {
	var e domain.Game
	err := r.db.QueryRow(
		ctx,
		`SELECT match_id, stadium_id, team_1_id, team_2_id, match_date
		 FROM Game WHERE match_id=$1`,
		id,
	).Scan(&e.ID, &e.StadiumID, &e.Team1ID, &e.Team2ID, &e.MatchDate)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *gameRepo) Update(ctx context.Context, e *domain.Game) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE Game
		 SET stadium_id=$1, team_1_id=$2, team_2_id=$3, match_date=$4
		 WHERE match_id=$5`,
		e.StadiumID, e.Team1ID, e.Team2ID, e.MatchDate, e.ID,
	)
	return err
}

func (r *gameRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM Game WHERE match_id=$1`, id)
	return err
}

func (r *gameRepo) List(ctx context.Context, limit, offset int) ([]*domain.Game, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT match_id, stadium_id, team_1_id, team_2_id, match_date
		 FROM Game
		 ORDER BY match_id
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*domain.Game
	for rows.Next() {
		e := new(domain.Game)
		err := rows.Scan(&e.ID, &e.StadiumID, &e.Team1ID, &e.Team2ID, &e.MatchDate)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, rows.Err()
}
