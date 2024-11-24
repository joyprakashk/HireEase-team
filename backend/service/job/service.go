package job

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
	store "hire-ease/store/job"
)

type service struct {
	store store.Job
}

// New - is a factory function to inject store in service.
func New(s store.Job) Job {
	return service{store: s}
}

func (s service) Create(ctx *gofr.Context, job *model.Job) (*model.Job, error) {
	return s.store.Create(ctx, job)
}

func (s service) GetAll(ctx *gofr.Context) ([]model.Job, error) {
	return s.store.GetAll(ctx)
}

func (s service) GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Job, error) {
	return s.store.GetByID(ctx, id)
}

func (s service) Update(ctx *gofr.Context, job *model.Job) (*model.Job, error) {
	return s.store.Update(ctx, job)
}

func (s service) Delete(ctx *gofr.Context, id uuid.UUID) error {
	return s.store.Delete(ctx, id)
}
