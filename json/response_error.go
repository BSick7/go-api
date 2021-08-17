package json

type StatusCoder interface {
	StatusCode() int
}

type RequestIder interface {
	RequestId() string
}
