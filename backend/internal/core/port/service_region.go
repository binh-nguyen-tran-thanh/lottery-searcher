package port

import "backend/internal/core/domain"

type RegionServicer interface {
	GetAllRegion() ([]domain.Region, error)
}
