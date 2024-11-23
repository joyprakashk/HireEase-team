package job

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
)

type store struct{}

// New is a factory function for store layer.
func New() Job {
	return store{}
}

func (s store) Create(ctx *gofr.Context, job *model.Job) (*model.Job, error) {
	const uniqueViolation = "23505"

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("error in creating uuid")
	}

	createdAt := time.Now().UTC().Format(time.RFC3339)

	_, err = ctx.SQL.ExecContext(ctx, "INSERT INTO jobs (id, title, description, company, location, domain, skills, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)",
		id, job.Title, job.Desc, job.Company, job.Location, job.Domain, job.Skills, createdAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == uniqueViolation { // unique_violation
				return nil, errors.New("entity already exists")
			}
		}

		return nil, err
	}

	job.ID = id

	return job, nil
}

func (s store) GetAll(ctx *gofr.Context) ([]model.Job, error) {
	rows, err := ctx.SQL.QueryContext(ctx, "SELECT id, title, description, company, location, domain, skills FROM jobs WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []model.Job

	for rows.Next() {
		var job model.Job

		err = rows.Scan(&job.ID, &job.Title, &job.Desc, &job.Company, &job.Location, &job.Domain, &job.Skills)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (s store) GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Job, error) {
	var job model.Job

	err := ctx.SQL.QueryRowContext(ctx, "SELECT id, title, description, company, location, domain, skills FROM jobs WHERE id=$1 and deleted_at IS NULL", id).
		Scan(&job.ID, &job.Title, &job.Desc, &job.Company, &job.Location, &job.Domain, &job.Skills)

	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	return &job, nil
}

func (s store) Update(ctx *gofr.Context, job *model.Job) (*model.Job, error) {
	updatedAt := time.Now().UTC().Format(time.RFC3339)

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE jobs SET title=$1, description=$2, company=$3, location=$4, domain=$5, skills=$6, updated_at=$7 WHERE id=$8 and deleted_at IS NULL",
		job.Title, job.Desc, job.Company, job.Location, job.Domain, job.Skills, updatedAt, job.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return nil, errors.New("entity not found")
	}

	return job, nil
}

func (s store) Delete(ctx *gofr.Context, id uuid.UUID) error {
	deletedAt := time.Now().UTC().Format(time.RFC3339)
	updatedAt := deletedAt

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE jobs SET deleted_at=$1, updated_at=$2 WHERE id=$3 and deleted_at IS NULL", deletedAt, updatedAt, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return err
	}

	return nil
}
