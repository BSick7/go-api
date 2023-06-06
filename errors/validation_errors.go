package errors

import "fmt"

type ValidationError struct {
	error
	Context string
}

var _ error = ValidationErrors{}

type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	result := ""
	if len(ve) > 0 {
		result += "validation errors:\n"
	}
	for _, err := range ve {
		result += fmt.Sprintf("\t%s %s\n", err.Context, err.Error())
	}
	return result
}

func (ve ValidationErrors) ToJson() map[string][]string {
	result := map[string][]string{}
	for _, err := range ve {
		if _, ok := result[err.Context]; !ok {
			result[err.Context] = make([]string, 0)
		}
		result[err.Context] = append(result[err.Context], err.Error())
	}
	return result
}
