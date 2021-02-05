package logging

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BSick7/go-api/intercept"
)

func LogAllRequests(prefix string) intercept.OnResponseFunc {
	stdNoTime := log.New(os.Stderr, prefix, 0)

	return func(r *http.Request, data intercept.ResponseData, duration time.Duration) {
		stdNoTime.Printf("%s %d %s %s %s", duration, data.StatusCode(), r.Method, r.RequestURI, data.Body())
	}
}
