package aircraft

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates use-case logic for aircraft.
type Service interface {
	Get(ctx context.Context, id int) (Aircraft, error)
	Query(ctx context.Context, offset, limit int) ([]Aircraft, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateAircraftRequest) (Aircraft, error)
	Update(ctx context.Context, id int, input UpdateAircraftRequest) (Aircraft, error)
	Delete(ctx context.Context, id int) (Aircraft, error)
}

// Aircraft represents the data about an aircraft.
type Aircraft struct {
	entity.Aircraft
}

// CreateAircraftRequest represents an aircraft creation request.
type CreateAircraftRequest struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

// Validate validates the CreateAircraftRequest fields.
func (m CreateAircraftRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Status, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Name, validation.Required, validation.Length(0, 255)),
	)
}

// UpdateAircraftRequest represents an aircraft update request.
type UpdateAircraftRequest struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

// Validate validates the UpdateAircraftRequest fields.
func (m UpdateAircraftRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Status, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Name, validation.Required, validation.Length(0, 255)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new aircraft service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the aircraft with the specified aircraft ID.
func (s service) Get(ctx context.Context, id int) (Aircraft, error) {
	aircraft, err := s.repo.Get(ctx, id)
	if err != nil {
		return Aircraft{}, err
	}
	return Aircraft{aircraft}, nil
}

// Create creates a new aircraft.
func (s service) Create(ctx context.Context, req CreateAircraftRequest) (Aircraft, error) {
	if err := req.Validate(); err != nil {
		return Aircraft{}, err
	}
	now := time.Now()
	aircraft := entity.Aircraft{
		Status:    req.Status,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := s.repo.Create(ctx, aircraft)
	if err != nil {
		return Aircraft{}, err
	}
	return Aircraft{aircraft}, nil
}

// Update updates the aircraft with the specified ID.
func (s service) Update(ctx context.Context, id int, req UpdateAircraftRequest) (Aircraft, error) {
	if err := req.Validate(); err != nil {
		return Aircraft{}, err
	}

	aircraft, err := s.Get(ctx, id)
	if err != nil {
		return aircraft, err
	}
	aircraft.Status = req.Status
	aircraft.Name = req.Name
	aircraft.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, aircraft.Aircraft); err != nil {
		return aircraft, err
	}
	return aircraft, nil
}

// Delete deletes the aircraft with the specified ID.
func (s service) Delete(ctx context.Context, id int) (Aircraft, error) {
	aircraft, err := s.Get(ctx, id)
	if err != nil {
		return Aircraft{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Aircraft{}, err
	}
	return aircraft, nil
}

// Count returns the number of aircraft.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the aircraft with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Aircraft, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Aircraft{}
	for _, item := range items {
		result = append(result, Aircraft{item})
	}
	return result, nil
}
