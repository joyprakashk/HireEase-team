package student

import (
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
	store "hire-ease/store/student"
)

type service struct {
	store store.Student
}

// New - is a factory function to inject store in service.
func New(s store.Student) Student {
	return service{store: s}
}

func (s service) Create(ctx *gofr.Context, order *model.Student) (*model.Student, error) {
	return s.store.Create(ctx, order)
}

func (s service) GetAll(ctx *gofr.Context) ([]model.Student, error) {
	return s.store.GetAll(ctx)
}

func (s service) GetByID(ctx *gofr.Context, id string) (*model.Student, error) {
	return s.store.GetByID(ctx, id)
}

func (s service) Update(ctx *gofr.Context, order *model.Student) (*model.Student, error) {
	return s.store.Update(ctx, order)
}

func (s service) Delete(ctx *gofr.Context, id string) error {
	return s.store.Delete(ctx, id)
}

func (s service) Login(ctx *gofr.Context, id string, auth model.Auth) error {
	return s.store.Login(ctx, id, auth)
}
