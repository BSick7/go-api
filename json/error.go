package json

import (
	"fmt"
)

type Error struct {
	StatusCode   int    `json:"-"`
	Status       string `json:"-"`
	ErrorMessage string `json:"error"`
}

func (e Error) Error() string {
	if e.ErrorMessage != "" {
		return fmt.Sprintf("(%d) %s: %s", e.StatusCode, e.Status, e.ErrorMessage)
	} else {
		return fmt.Sprintf("(%d) %s", e.StatusCode, e.Status)
	}
}
