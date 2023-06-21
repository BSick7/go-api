package errors

import (
	"fmt"
)

var _ error = ValidationError{}
var _ ResponseErrorer = ValidationError{}

type ValidationError struct {
	Context string
	Message string
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("%s - %s", ve.Context, ve.Message)
}

func (ve ValidationError) ResponseError() error {
	return InvalidRequestError{
		Errors: ValidationErrors{ve},
	}
}

var _ error = ValidationErrors{}
var _ ResponseErrorer = ValidationErrors{}

type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	result := ""
	if len(ve) > 0 {
		result += "validation errors:\n"
	}
	for _, err := range ve {
		result += fmt.Sprintf("\t%s\n", err.Error())
	}
	return result
}

func (ve ValidationErrors) ToMap() map[string][]string {
	result := map[string][]string{}
	for _, err := range ve {
		context := err.Context
		if context == "" {
			context = "Base"
		}
		if _, ok := result[err.Context]; !ok {
			result[err.Context] = make([]string, 0)
		}
		result[err.Context] = append(result[err.Context], err.Message)
	}
	return result
}

func ValidationErrorsFromMap(m map[string][]string) ValidationErrors {
	result := ValidationErrors{}
	for context, messages := range m {
		for _, message := range messages {
			result = append(result, ValidationError{Context: context, Message: message})
		}
	}
	return result
}

func (ve ValidationErrors) ResponseError() error {
	return InvalidRequestError{
		Errors: ve,
	}
}
