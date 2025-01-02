package aircraft

import (
	"net/http"
	"strconv"

	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/aircraft/<id>", res.get)
	r.Get("/aircraft", res.query)

	r.Use(authHandler)

	// The following endpoints require a valid JWT
	r.Post("/aircraft", res.create)
	r.Put("/aircraft/<id>", res.update)
	r.Delete("/aircraft/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.BadRequest("invalid ID format")
	}
	aircraft, err := r.service.Get(c.Request.Context(), id)
	if err != nil {
		return err
	}

	return c.Write(aircraft)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	aircraft, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = aircraft
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateAircraftRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("invalid input")
	}
	aircraft, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(aircraft, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.BadRequest("invalid ID format")
	}

	var input UpdateAircraftRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("invalid input")
	}

	aircraft, err := r.service.Update(c.Request.Context(), id, input)
	if err != nil {
		return err
	}

	return c.Write(aircraft)
}

func (r resource) delete(c *routing.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.BadRequest("invalid ID format")
	}
	aircraft, err := r.service.Delete(c.Request.Context(), id)
	if err != nil {
		return err
	}

	return c.Write(aircraft)
}
