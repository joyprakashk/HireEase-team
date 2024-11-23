package student

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
func New() Student {
	return store{}
}

func (s store) Create(ctx *gofr.Context, std *model.Student) (*model.Student, error) {
	const uniqueViolation = "23505"

	createdAt := time.Now().UTC().Format(time.RFC3339)

	_, err := ctx.SQL.ExecContext(ctx, "INSERT INTO students (id, name, gender, branch, cgpa, courses, github, linkedin, interests, skills, location, password, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)",
		std.ID, std.Name, std.Gender, std.Branch, std.CGPA, std.Courses, std.GitHub, std.LinkedIn, std.Interests, std.Skills, std.Location, std.Password, createdAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == uniqueViolation { // unique_violation
				return nil, errors.New("entity already exists")
			}
		}

		return nil, err
	}

	return std, nil
}

func (s store) GetAll(ctx *gofr.Context) ([]model.Student, error) {
	rows, err := ctx.SQL.QueryContext(ctx, "SELECT id, name, gender, branch, cgpa, courses, github, linkedin, interests, skills, location FROM students WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var stds []model.Student

	for rows.Next() {
		var std model.Student

		err = rows.Scan(&std.ID, &std.Name, &std.Gender, &std.Branch, &std.CGPA, &std.Courses, &std.GitHub,
			&std.LinkedIn, &std.Interests, &std.Skills, &std.Location)
		if err != nil {
			return nil, err
		}

		stds = append(stds, std)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return stds, nil
}

func (s store) GetByID(ctx *gofr.Context, id string) (*model.Student, error) {
	var std model.Student

	err := ctx.SQL.QueryRowContext(ctx, "SELECT id, name, gender, branch, cgpa, courses, github, linkedin, interests, skills, location FROM students WHERE id=$1 and deleted_at IS NULL", id).
		Scan(&std.ID, &std.Name, &std.Gender, &std.Branch, &std.CGPA, &std.Courses, &std.GitHub, &std.LinkedIn,
			&std.Interests, &std.Skills, &std.Location)

	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	return &std, nil
}

func (s store) Update(ctx *gofr.Context, std *model.Student) (*model.Student, error) {
	updatedAt := time.Now().UTC().Format(time.RFC3339)

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE students SET name=$1, gender=$2, branch=$3, cgpa=$4, courses=$5, github=$6, linkedin=$7, interests=$8, skills=$9, location=$10, password=$11, updated_at=$12 WHERE id=$13 and deleted_at IS NULL",
		std.Name, std.Gender, std.Branch, std.CGPA, std.Courses, std.GitHub, std.LinkedIn, std.Interests, std.Skills,
		std.Location, std.Password, updatedAt, std.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return nil, errors.New("entity not found")
	}

	return std, nil
}

func (s store) Delete(ctx *gofr.Context, id string) error {
	deletedAt := time.Now().UTC().Format(time.RFC3339)
	updatedAt := deletedAt

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE students SET deleted_at=$1, updated_at=$2 WHERE id=$3 and deleted_at IS NULL", deletedAt, updatedAt, id)
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

	err := ctx.SQL.QueryRowContext(ctx, "SELECT password FROM students WHERE id=$1 and deleted_at IS NULL", id).
		Scan(&pass)
	if err != nil {
		return err
	}

	if auth.Key != pass {
		return errors.New("invalid password")
	}

	return nil
}
