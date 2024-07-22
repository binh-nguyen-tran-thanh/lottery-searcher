package port

import "backend/internal/core/domain"

type RegionRepository interface {
	GetRegions() ([]domain.Region, error)
}
