package json

type StatusCoder interface {
	StatusCode() int
}

type ResponsePayloader interface {
	Payload() map[string]interface{}
}
