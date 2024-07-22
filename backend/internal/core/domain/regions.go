package domain

type Region struct {
	Name     string
	Code     string
	IsActive bool
}

func NewRegion(arg Region) *Region {
	return &Region{
		Name:     arg.Name,
		Code:     arg.Code,
		IsActive: arg.IsActive,
	}
}
