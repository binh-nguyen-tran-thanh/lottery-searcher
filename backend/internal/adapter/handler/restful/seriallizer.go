package restful

import (
	"backend/internal/core/domain"
	"encoding/json"
)

type regionResponse struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	IsActive bool   `json:"isActive"`
}

type resultResponse struct {
	OpenTime string   `json:"open_time"`
	Detail   []string `json:"detail"`
	Region   string   `json:"region"`
	ID       uint     `json:"id"`
}

type openNumbResponse struct {
	Result resultResponse `json:"result"`
	Numbs  string         `json:"numbs"`
	Rank   int8           `json:"rank"`
}

func newResultResponse(in domain.Result) (out resultResponse) {
	var details []string
	if err := json.Unmarshal([]byte(in.Detail), &details); err != nil {
		return
	}

	out = resultResponse{
		OpenTime: in.OpenTime,
		Detail:   details,
		Region:   in.Region,
		ID:       in.ID,
	}
	return
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
		out = append(out, newResultResponse(r))
	}

	return
}

func serializeOpenNumList(openNumbs []domain.OpenNum) (out []openNumbResponse) {
	for _, r := range openNumbs {
		out = append(out, openNumbResponse{
			Result: newResultResponse(r.Result),
			Numbs:  r.Numbs,
			Rank:   r.Rank,
		})
	}

	return
}
