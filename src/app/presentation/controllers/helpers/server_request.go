package helpers

import (
	"net/http"
)

type ServerErrorRequest struct {
	Status int32  `json:"status"`
	Error  string `json:"error"`
}

func NewServerErrorRequest(err error) *ServerErrorRequest {
	return &ServerErrorRequest{Status: http.StatusInternalServerError, Error: "err"}
}
