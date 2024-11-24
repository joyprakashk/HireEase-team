package job

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
)

type Job interface {
	Create(ctx *gofr.Context, order *model.Job) (*model.Job, error)
	GetAll(ctx *gofr.Context) ([]model.Job, error)
	GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Job, error)
	Update(ctx *gofr.Context, order *model.Job) (*model.Job, error)
	Delete(ctx *gofr.Context, id uuid.UUID) error
}
