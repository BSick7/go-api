package expect

import "net/http"

type Response struct {
	Data       interface{}
	StatusCode int
	Err        string
}

func OK(data interface{}) Response {
	return Response{Data: data, StatusCode: http.StatusOK, Err: ""}
}

func NoContent() Response {
	return Response{Data: nil, StatusCode: http.StatusNoContent, Err: ""}
}

func NotFound(msg string) Response {
	return Response{Data: nil, StatusCode: http.StatusNotFound, Err: msg}
}

func BadRequest(msg string) Response {
	return Response{Data: nil, StatusCode: http.StatusBadRequest, Err: msg}
}

func InternalError(msg string) Response {
	return Response{Data: nil, StatusCode: http.StatusInternalServerError, Err: msg}
}
