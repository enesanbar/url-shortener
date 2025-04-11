package inmemory

import (
	"context"
	"errors"
	"sync"

	serviceErr "github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/url-shortener/internal/domain"
)

type MappingInmemoryAdapter struct {
	mtx       sync.RWMutex
	Redirects map[string]*domain.Mapping
}

func NewMappingInmemoryAdapter(db map[string]*domain.Mapping) *MappingInmemoryAdapter {
	return &MappingInmemoryAdapter{
		Redirects: db,
	}
}

func (rr *MappingInmemoryAdapter) Store(_ context.Context, m *domain.Mapping) (*domain.Mapping, error) {
	rr.mtx.Lock()
	defer rr.mtx.Unlock()

	rr.Redirects[m.Code] = m
	return m, nil
}

func (rr *MappingInmemoryAdapter) Update(ctx context.Context, m *domain.Mapping) (*domain.Mapping, error) {
	return rr.Store(ctx, m)
}

func (rr *MappingInmemoryAdapter) FindAll(_ context.Context, page int64, pageSize int64) ([]*domain.Mapping, error) {
	rr.mtx.RLock()
	defer rr.mtx.RUnlock()

	mappings := make([]*domain.Mapping, 0)

	for _, mapping := range rr.Redirects {
		mappings = append(mappings, mapping)
	}

	return mappings, nil
}

func (rr *MappingInmemoryAdapter) FindByCode(_ context.Context, code string) (*domain.Mapping, error) {
	rr.mtx.RLock()
	defer rr.mtx.RUnlock()

	if redirect, ok := rr.Redirects[code]; ok {
		return redirect, nil
	}

	return nil, serviceErr.Error{
		Err:  errors.New("mapping not found"),
		Code: serviceErr.ENOTFOUND,
	}
}

func (rr *MappingInmemoryAdapter) Delete(ctx context.Context, code string) error {
	return nil
}
