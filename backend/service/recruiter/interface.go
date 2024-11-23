package recruiter

import (
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
)

type Recruiter interface {
	Create(ctx *gofr.Context, order *model.Recruiter) (*model.Recruiter, error)
	GetAll(ctx *gofr.Context) ([]model.Recruiter, error)
	GetByID(ctx *gofr.Context, id string) (*model.Recruiter, error)
	Update(ctx *gofr.Context, order *model.Recruiter) (*model.Recruiter, error)
	Delete(ctx *gofr.Context, id string) error
	Login(ctx *gofr.Context, id string, auth model.Auth) error
}
