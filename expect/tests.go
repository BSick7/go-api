package expect

import (
	gocontext "context"
	"fmt"
	"testing"

	"github.com/BSick7/go-api"
	"github.com/BSick7/go-api/context"
)

type Test struct {
	CtxValues context.Values
	Send      RequestData
	Decoder   func(a interface{}, b interface{})
	Want      Response
}

type Tests []Test

func (tests Tests) Run(t *testing.T, handler api.EndpointHandler) {
	for i, test := range tests {
		mreq := NewRequest(test.Send, context.WithValues(gocontext.Background(), test.CtxValues))
		if test.Decoder != nil {
			mreq.CopyPointer = test.Decoder
		} else {
			mreq.CopyPointer = func(a interface{}, b interface{}) {
				t.Fatalf("[%d] test needs decoder mocked", i)
			}
		}
		mres := NewResponder(t, fmt.Sprintf("[%d] ", i), test.Want)
		handler(mres, mreq)
		mres.VerifyNotFound()
	}
}
