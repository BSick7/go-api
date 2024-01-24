package app1

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/BSick7/go-api/json"
)

func TestSimple(t *testing.T) {
	tests := json.Tests{
		{
			Name:    "missing-input",
			Request: httptest.NewRequest("GET", "/simple", nil),
			Want:    json.ExpectBadRequest(1, map[string]string{"message": "missing data"}),
		},
		{
			// NOTE: This is a generic error message so it's obscured by API middleware
			Name:    "invalid-input",
			Request: httptest.NewRequest("GET", "/simple?data=hey", nil),
			Want:    json.ExpectInternalError(fmt.Errorf("We have encountered an unexpected error.")),
		},
		{
			Name:    "success",
			Request: httptest.NewRequest("GET", "/simple?data=10", nil),
			Want:    json.ExpectOK(10).WithHeader("Have", "data"),
		},
	}
	tests.Run(t, Server())
}
