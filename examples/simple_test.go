package examples

import (
	"testing"

	"github.com/BishopFox/go-api/expect"
)

func TestSimple(t *testing.T) {
	tests := expect.Tests{
		// missing data
		{
			Send: expect.RequestData{},
			Want: expect.BadRequest("missing data"),
		},
		// invalid data
		{
			Send: expect.SendVars(map[string]string{"data": "hey"}),
			Want: expect.InternalError("invalid syntax"),
		},
		// success
		{
			Send: expect.SendVars(map[string]string{"data": "10"}),
			Want: expect.OK(10),
		},
	}
	tests.Run(t, Simple)
}
