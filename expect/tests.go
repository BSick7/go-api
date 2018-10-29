package expect

import (
	gocontext "context"
	"testing"

	"github.com/BishopFox/go-api"
	"github.com/BishopFox/go-api/context"
)

type Test struct {
	Name       string
	CtxWrapper context.Wrapper
	Send       RequestData
	Decoder    func(a interface{}, b interface{})
	Want       Response
}

type Tests []Test

func (tests Tests) Run(t *testing.T, handler api.EndpointHandler) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctx := gocontext.Background()
			if test.CtxWrapper != nil {
				ctx = test.CtxWrapper(ctx)
			}
			mreq := NewRequest(test.Send, ctx)
			if test.Decoder != nil {
				mreq.CopyPointer = test.Decoder
			} else {
				mreq.CopyPointer = func(a interface{}, b interface{}) {
					t.Fatal("test needs decoder mocked")
				}
			}
			mres := NewResponder(t, "", test.Want)
			handler(mres, mreq)
			mres.VerifyNotFound()
		})
	}
}
