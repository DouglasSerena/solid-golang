package helpers

import (
	"net/http"
)

type BadRequest struct {
	Status int32  `json:"status"`
	Error  string `json:"error"`
}

func NewBadRequest(err error) *BadRequest {
	return &BadRequest{Status: http.StatusBadRequest, Error: err.Error()}
}
