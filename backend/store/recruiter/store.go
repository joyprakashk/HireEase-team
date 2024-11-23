package recruiter

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
)

type store struct{}

// New is a factory function for store layer.
func New() Recruiter {
	return store{}
}

func (s store) Create(ctx *gofr.Context, rec *model.Recruiter) (*model.Recruiter, error) {
	const uniqueViolation = "23505"

	createdAt := time.Now().UTC().Format(time.RFC3339)

	_, err := ctx.SQL.ExecContext(ctx, "INSERT INTO recruiters (id, name, company, jobs, password, created_at) VALUES ($1,$2,$3,$4,$5,$6)",
		rec.ID, rec.Name, rec.Company, pq.Array(rec.Jobs), rec.Password, createdAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == uniqueViolation { // unique_violation
				return nil, errors.New("entity already exists")
			}
		}

		return nil, err
	}

	return rec, nil
}

func (s store) GetAll(ctx *gofr.Context) ([]model.Recruiter, error) {
	rows, err := ctx.SQL.QueryContext(ctx, "SELECT id, name, company, jobs FROM recruiters WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var recs []model.Recruiter

	for rows.Next() {
		var rec model.Recruiter

		err = rows.Scan(&rec.ID, &rec.Name, &rec.Company, pq.Array(&rec.Jobs))
		if err != nil {
			return nil, err
		}

		recs = append(recs, rec)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return recs, nil
}

func (s store) GetByID(ctx *gofr.Context, id string) (*model.Recruiter, error) {
	var rec model.Recruiter

	err := ctx.SQL.QueryRowContext(ctx, "SELECT id, name, company, jobs FROM recruiters WHERE id=$1 and deleted_at IS NULL", id).
		Scan(&rec.ID, &rec.Name, &rec.Company, pq.Array(&rec.Jobs))

	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	return &rec, nil
}

func (s store) Update(ctx *gofr.Context, rec *model.Recruiter) (*model.Recruiter, error) {
	updatedAt := time.Now().UTC().Format(time.RFC3339)

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE recruiters SET name=$1, company=$2, jobs=$3, password=$4, updated_at=$5 WHERE id=$6 and deleted_at IS NULL",
		rec.Name, rec.Company, pq.Array(rec.Jobs), rec.Password, updatedAt, rec.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return nil, errors.New("entity not found")
	}

	return rec, nil
}

func (s store) Delete(ctx *gofr.Context, id string) error {
	deletedAt := time.Now().UTC().Format(time.RFC3339)
	updatedAt := deletedAt

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE recruiters SET deleted_at=$1, updated_at=$2 WHERE id=$3 and deleted_at IS NULL", deletedAt, updatedAt, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return err
	}

	return nil
}

func (s store) Login(ctx *gofr.Context, id string, auth model.Auth) error {
	var pass string

	err := ctx.SQL.QueryRowContext(ctx, "SELECT password FROM recruiters WHERE id=$1 and deleted_at IS NULL", id).
		Scan(&pass)
	if err != nil {
		return err
	}

	if auth.Key != pass {
		return errors.New("invalid password")
	}

	return nil
}
