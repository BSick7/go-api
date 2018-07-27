package expect

import (
	gocontext "context"
	"fmt"
	"testing"

	"github.com/BSick7/go-api"
	"github.com/BSick7/go-api/context"
)

type Test struct {
	CtxWrapper context.Wrapper
	Send       RequestData
	Decoder    func(a interface{}, b interface{})
	Want       Response
}

type Tests []Test

func (tests Tests) Run(t *testing.T, handler api.EndpointHandler) {
	for i, test := range tests {
		ctx := gocontext.Background()
		if test.CtxWrapper != nil {
			ctx = test.CtxWrapper(ctx)
		}
		mreq := NewRequest(test.Send, ctx)
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
