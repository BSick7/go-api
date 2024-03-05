package logging

import (
	"github.com/BSick7/go-api/intercept"
	"log"
	"net/http"
	"os"
)

func LogAllRequests(prefix string) intercept.OnResponseFunc {
	stdNoTime := log.New(os.Stderr, prefix, 0)

	return func(r *http.Request, data intercept.ResponseData) {
		stdNoTime.Printf("%s %d %s %s", data.Duration, data.StatusCode, r.Method, r.RequestURI)
	}
}
