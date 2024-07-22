package restful

import "backend/internal/core/util/exception"

type BaseResponse struct {
	IsSuccess bool          `json:"isSuccess"`
	Data      interface{}   `json:"data"`
	Errors    exception.Err `json:"errors"`
}

func NewSuccessResponse(data interface{}) *BaseResponse {
	return &BaseResponse{
		IsSuccess: true,
		Data:      data,
	}
}

func NewErrorResponse(err *exception.Exception) *BaseResponse {
	return &BaseResponse{
		IsSuccess: false,
		Errors:    err.Errors,
		Data:      nil,
	}
}
