package aircraft

import (
	"context"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access aircraft from the data source.
type Repository interface {
	// Get returns the aircraft with the specified aircraft ID.
	Get(ctx context.Context, id int) (entity.Aircraft, error)
	// Count returns the number of aircraft.
	Count(ctx context.Context) (int, error)
	// Query returns the list of aircraft with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Aircraft, error)
	// Create saves a new aircraft in the storage.
	Create(ctx context.Context, aircraft entity.Aircraft) error
	// Update updates the aircraft with given ID in the storage.
	Update(ctx context.Context, aircraft entity.Aircraft) error
	// Delete removes the aircraft with given ID from the storage.
	Delete(ctx context.Context, id int) error
}

// repository persists aircraft in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new aircraft repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the aircraft with the specified ID from the database.
func (r repository) Get(ctx context.Context, id int) (entity.Aircraft, error) {
	var aircraft entity.Aircraft
	err := r.db.With(ctx).Select().Model(id, &aircraft)
	return aircraft, err
}

// Create saves a new aircraft record in the database.
func (r repository) Create(ctx context.Context, aircraft entity.Aircraft) error {
	return r.db.With(ctx).Model(&aircraft).Insert()
}

// Update saves the changes to an aircraft in the database.
func (r repository) Update(ctx context.Context, aircraft entity.Aircraft) error {
	return r.db.With(ctx).Model(&aircraft).Update()
}

// Delete deletes an aircraft with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id int) error {
	aircraft, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&aircraft).Delete()
}

// Count returns the number of the aircraft records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("aircraft").Row(&count)
	return count, err
}

// Query retrieves the aircraft records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Aircraft, error) {
	var aircraft []entity.Aircraft
	err := r.db.With(ctx).
		Select().
		OrderBy("aircraft_id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&aircraft)
	return aircraft, err
}
