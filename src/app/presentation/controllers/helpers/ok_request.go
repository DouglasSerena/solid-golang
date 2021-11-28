package helpers

import (
	"net/http"
)

type OkRequest struct {
	Status int32       `json:"status"`
	Data   interface{} `json:"data"`
}

func NewOkRequest(data interface{}) *OkRequest {
	return &OkRequest{Status: http.StatusOK, Data: data}
}
