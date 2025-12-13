package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hse-football/internal/domain"
)

type playerRepo struct {
	db *pgxpool.Pool
}

func NewPlayerRepo(db *pgxpool.Pool) domain.PlayerRepository {
	return &playerRepo{db: db}
}

func (r *playerRepo) Create(ctx context.Context, e *domain.Player) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO Player (player_name, player_surname, player_number, salary, phone, birth_date, team_id, status_id)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING player_id`,
		e.Name, e.Surname, e.Number, e.Salary, e.Phone, e.BirthDate, e.TeamID, e.StatusID,
	).Scan(&id)
	return id, err
}

func (r *playerRepo) GetByID(ctx context.Context, id int64) (*domain.Player, error) {
	var e domain.Player
	err := r.db.QueryRow(
		ctx,
		`SELECT player_id, player_name, player_surname, player_number, salary, phone, birth_date, team_id, status_id
		 FROM Player WHERE player_id=$1`,
		id,
	).Scan(&e.ID, &e.Name, &e.Surname, &e.Number, &e.Salary, &e.Phone, &e.BirthDate, &e.TeamID, &e.StatusID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *playerRepo) Update(ctx context.Context, e *domain.Player) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE Player
		 SET player_name=$1, player_surname=$2, player_number=$3, salary=$4, phone=$5, birth_date=$6, team_id=$7, status_id=$8
		 WHERE player_id=$9`,
		e.Name, e.Surname, e.Number, e.Salary, e.Phone, e.BirthDate, e.TeamID, e.StatusID, e.ID,
	)
	return err
}

func (r *playerRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM Player WHERE player_id=$1`, id)
	return err
}

func (r *playerRepo) List(ctx context.Context, limit, offset int) ([]*domain.Player, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT player_id, player_name, player_surname, player_number, salary, phone, birth_date, team_id, status_id
		 FROM Player
		 ORDER BY player_id
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*domain.Player
	for rows.Next() {
		e := new(domain.Player)
		err := rows.Scan(&e.ID, &e.Name, &e.Surname, &e.Number, &e.Salary, &e.Phone, &e.BirthDate, &e.TeamID, &e.StatusID)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, rows.Err()
}
