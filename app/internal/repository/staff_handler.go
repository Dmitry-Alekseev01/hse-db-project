package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"hse-football/internal/domain"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) domain.StaffRepository {
	return &staffRepo{db: db}
}

func (r *staffRepo) Create(ctx context.Context, e *domain.Staff) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO Staff (staff_name, staff_surname, salary, specification_id, club_id)
		 VALUES ($1, $2, $3, $4, $5) RETURNING staff_id`,
		e.Name, e.Surname, e.Salary, e.SpecificationID, e.ClubID,
	).Scan(&id)
	return id, err
}

func (r *staffRepo) GetByID(ctx context.Context, id int64) (*domain.Staff, error) {
	var e domain.Staff
	err := r.db.QueryRow(
		ctx,
		`SELECT staff_id, staff_name, staff_surname, salary, specification_id, club_id
		 FROM Staff WHERE staff_id=$1`,
		id,
	).Scan(&e.ID, &e.Name, &e.Surname, &e.Salary, &e.SpecificationID, &e.ClubID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *staffRepo) Update(ctx context.Context, e *domain.Staff) error {
	_, err := r.db.Exec(
		ctx,
		`UPDATE Staff
		 SET staff_name=$1, staff_surname=$2, salary=$3, specification_id=$4, club_id=$5
		 WHERE staff_id=$6`,
		e.Name, e.Surname, e.Salary, e.SpecificationID, e.ClubID, e.ID,
	)
	return err
}

func (r *staffRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `DELETE FROM Staff WHERE staff_id=$1`, id)
	return err
}

func (r *staffRepo) List(ctx context.Context, limit, offset int) ([]*domain.Staff, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT staff_id, staff_name, staff_surname, salary, specification_id, club_id
		 FROM Staff
		 ORDER BY staff_id
		 LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*domain.Staff
	for rows.Next() {
		e := new(domain.Staff)
		err := rows.Scan(&e.ID, &e.Name, &e.Surname, &e.Salary, &e.SpecificationID, &e.ClubID)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, rows.Err()
}
