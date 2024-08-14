package restful

import "backend/internal/core/domain"

type regionResponse struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	IsActive bool   `json:"isActive"`
}

type resultResponse struct {
	OpenTime string `json:"open_time"`
	Detail   string `json:"detail"`
	Region   string `json:"region"`
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

func serializeResults(results []domain.Result) (out []resultResponse) {
	for _, r := range results {
		out = append(out, resultResponse{
			OpenTime: r.OpenTime,
			Detail:   r.Detail,
			Region:   r.Region,
		})
	}

	return
}
