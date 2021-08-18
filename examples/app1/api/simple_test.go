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
			Want:    json.ExpectBadRequest(map[string]string{"data": "missing data"}),
		},
		{
			Name:    "invalid-input",
			Request: httptest.NewRequest("GET", "/simple?data=hey", nil),
			Want:    json.ExpectInternalError(fmt.Errorf("invalid syntax")),
		},
		{
			Name:    "success",
			Request: httptest.NewRequest("GET", "/simple?data=10", nil),
			Want:    json.ExpectOK(10).WithHeader("Have", "data"),
		},
	}
	tests.Run(t, Server())
}
