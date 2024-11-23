package student

import (
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
)

type Student interface {
	Create(ctx *gofr.Context, order *model.Student) (*model.Student, error)
	GetAll(ctx *gofr.Context) ([]model.Student, error)
	GetByID(ctx *gofr.Context, id string) (*model.Student, error)
	Update(ctx *gofr.Context, order *model.Student) (*model.Student, error)
	Delete(ctx *gofr.Context, id string) error
	Login(ctx *gofr.Context, id string, auth model.Auth) error
}
