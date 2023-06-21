package errors

type ResponseErrorer interface {
	ResponseError() error
}
