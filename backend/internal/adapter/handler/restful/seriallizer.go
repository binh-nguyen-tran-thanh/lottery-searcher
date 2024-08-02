package restful

import "backend/internal/core/domain"

type regionResponse struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	IsActive bool   `json:"isActive"`
}

func serializeRegion(regions []domain.Region) (result []regionResponse) {

	for _, region := range regions {
		result = append(result, regionResponse{
			Name:     region.Name,
			Code:     region.Code,
			IsActive: region.IsActive,
		})
	}

	return
}
