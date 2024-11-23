package student

import (
	"errors"
	"strings"

	"gofr.dev/pkg/gofr"
	"hire-ease/model"

	service "hire-ease/service/student"
)

type handler struct {
	service service.Student
}

// New - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func New(s service.Student) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var std model.Student

	err := ctx.Bind(&std)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	resp, err := h.service.Create(ctx, &std)
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

	resp, err := h.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var std model.Student

	id := ctx.PathParam("id")
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("missing param: ID")
	}

	err := ctx.Bind(&std)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	std.ID = id

	resp, err := h.service.Update(ctx, &std)
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

	err := h.service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h handler) Login(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.New("missing param: ID")
	}

	var auth model.Auth

	err := ctx.Bind(&auth)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	err = h.service.Login(ctx, id, auth)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
