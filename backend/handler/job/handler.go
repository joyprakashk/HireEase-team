package job

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"

	"hire-ease/model"
	svc "hire-ease/service/job"
)

type handler struct {
	service svc.Job
}

// New - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func New(s svc.Job) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var job model.Job

	err := ctx.Bind(&job)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	resp, err := h.service.Create(ctx, &job)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.New("missing param: ID")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid param: ID")
	}

	resp, err := h.service.GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var job model.Job

	id := ctx.PathParam("id")
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("missing param: ID")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid param: ID")
	}

	job.ID = uid

	err = ctx.Bind(&job)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	resp, err := h.service.Update(ctx, &job)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.New("missing param: ID")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid param: ID")
	}

	err = h.service.Delete(ctx, uid)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
