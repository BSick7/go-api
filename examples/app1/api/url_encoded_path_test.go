package app1

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/BSick7/go-api/json"
)

func TestUrlEncodedPath(t *testing.T) {
	tests := json.Tests{
		{
			Name:    "not-necessary-to-encode",
			Request: httptest.NewRequest("GET", "/paths/xyz/detail", nil),
			Want:    json.ExpectOK("xyz"),
		},
		{
			Name:    "encoded-target",
			Request: httptest.NewRequest("GET", "/paths/"+url.PathEscape("http://bishopfox.com")+"/detail", nil),
			Want:    json.ExpectOK("http://bishopfox.com"),
		},
		{
			Name:    "encoded-url",
			Request: httptest.NewRequest("GET", "/paths/"+url.PathEscape("http://bishopfox.com/?x=1&y1=test&z=&amp;copyright;")+"/detail", nil),
			Want:    json.ExpectOK("http://bishopfox.com/?x=1&y1=test&z=&amp;copyright;"),
		},
	}
	tests.Run(t, Server(nil))
}
