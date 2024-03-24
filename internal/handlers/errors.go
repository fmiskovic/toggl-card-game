package handlers

import (
	"encoding/json"
	"fmt"
)

// ApiError represents a custom error struct that contains error and http code.
type ApiError struct {
	Err  string `json:"errorMessage"`
	Code int    `json:"httpCode"`
}

func NewApiError(err string, code int) ApiError {
	return ApiError{
		Err:  err,
		Code: code,
	}
}

func (x ApiError) Error() string {
	data, err := json.Marshal(x)
	if err != nil {
		return "{error: " + x.Err + ", code: " + fmt.Sprint(x.Code) + "}"
	}
	return string(data)
}
