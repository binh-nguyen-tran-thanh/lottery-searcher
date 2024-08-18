package restful

import (
	"backend/internal/core/port"
	"time"
)

type pagingParamPayload struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type TimeRangeFilter struct {
	FilterRangeFrom time.Time `form:"range_from" time_format:"2006-01-02 15:04:05"`
	FilterRangeTo   time.Time `form:"range_to" time_format:"2006-01-02 15:04:05"`
}

type FilterOpenNumbPayload struct {
	pagingParamPayload
	TimeRangeFilter
	FilterTye   int    `form:"type"`
	Region      string `form:"region"`
	ID          uint   `form:"id"`
	ResultId    uint   `form:"result_id"`
	FilterMode  string `form:"mode"`
	FilterValue string `form:"value"`
	Ranks       []int  `form:"ranks"`
}

type FilterResultPayload struct {
	pagingParamPayload
	TimeRangeFilter
	Region string `form:"region"`
	ID     uint   `form:"id"`
}

func AsPortFilterOpenNumbPayload(arg FilterOpenNumbPayload) port.FilterOpenNumbPayload {
	return port.FilterOpenNumbPayload{
		FilterRangeFrom: arg.FilterRangeFrom,
		FilterRangeTo:   arg.FilterRangeTo,
		Region:          arg.Region,
		ID:              arg.ID,
		FilterTye:       arg.FilterTye,
		ResultId:        arg.ResultId,
		PagingParamPayload: port.PagingParamPayload{
			Limit:  arg.Limit,
			Offset: arg.Offset,
		},
		FilterMode:  arg.FilterMode,
		FilterValue: arg.FilterValue,
		Ranks:       arg.Ranks,
	}
}

func AsPortFilterResultPayload(arg FilterResultPayload) port.FilterResultPayload {
	return port.FilterResultPayload{
		FilterRangeFrom: arg.FilterRangeFrom,
		FilterRangeTo:   arg.FilterRangeTo,
		Region:          arg.Region,
		ID:              arg.ID,
		PagingParamPayload: port.PagingParamPayload{
			Limit:  arg.Limit,
			Offset: arg.Offset,
		}}
}
