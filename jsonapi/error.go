package jsonapi

import (
	"github.com/hashicorp/jsonapi"
	"net/http"
	"strconv"
)

func NewError(id string, statusCode int, title string, err error) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		ID:     id,
		Title:  title,
		Detail: err.Error(),
		Status: http.StatusText(statusCode),
		Code:   strconv.Itoa(statusCode),
	}
}
