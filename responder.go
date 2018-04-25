package api

type Responder interface {
	Send(data interface{}, context ...string)
	SendNotFound(msg string, context ...string)
	SendError(statusCode int, err error, context ...string)
}
