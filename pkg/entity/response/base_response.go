package response

import (
	"encoding/json"
	"io"
	"net/http"
)

type BaseResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  ErrorResponseData `json:"errors"`
	Data    interface{}       `json:"data"`
}

func NewBaseResponse(
	code int,
	message string,
	errors ErrorResponseData,
	data interface{},
) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}

func (baseResponse *BaseResponse) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(baseResponse)
}

func NewErrorResponse(code int, message string, errors ...ErrorResponseValue) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Errors: NewErrorResponseData(
			errors...,
		),
		Data: nil,
	}
}
func (baseResponse *BaseResponse) SendResponse(rw *http.ResponseWriter) error {
	(*rw).WriteHeader(baseResponse.Code)
	return json.NewEncoder(*rw).Encode(baseResponse)
}
