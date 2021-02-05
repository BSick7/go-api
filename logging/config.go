package logging

import (
	"net/http"
	"time"
)

type Config struct {
	// Prefix is prepended to all log statements
	Prefix string

	// Log100s will determine whether an HTTP 1xx response is logged
	Log100s bool

	// Log200s will determine whether an HTTP 2xx response is logged
	Log200s bool

	// Log300s will determine whether an HTTP 3xx response is logged
	Log300s bool

	// Log400s will determine whether an HTTP 4xx response is logged
	Log400s bool

	// Log500s will determine whether an HTTP 5xx response is logged
	Log500s bool

	// If specified, OnRequest is called on every request
	// This allows APIs to add error reporting to the existing logging
	OnRequest func(r *http.Request, data ResponseData, duration time.Duration)
}

func (c Config) ShouldLog(statusCode int) bool {
	if statusCode >= 100 && statusCode < 200 {
		return c.Log100s
	}
	if statusCode >= 200 && statusCode < 300 {
		return c.Log200s
	}
	if statusCode >= 300 && statusCode < 400 {
		return c.Log300s
	}
	if statusCode >= 400 && statusCode < 500 {
		return c.Log400s
	}
	if statusCode >= 500 {
		return c.Log500s
	}
	return false
}

type ResponseData interface {
	StatusCode() int
	Body() string
}
