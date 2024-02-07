package customerror

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	errors map[string]string
}

func NewValidationError() *ValidationError {
	return &ValidationError{errors: make(map[string]string)}
}

func NewValidationErrorWithMessage(key, message string) *ValidationError {
	return &ValidationError{
		errors: map[string]string{
			key: message,
		},
	}
}

func (ve *ValidationError) AddError(key, message string) {
	if val, ok := ve.errors[key]; ok {
		ve.errors[key] = val + "," + message
	} else {
		ve.errors[key] = message
	}
}

func (ve ValidationError) Error() string {
	var errorMessages []string
	for key, msg := range ve.errors {
		errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", key, msg))
	}
	return strings.Join(errorMessages, ";")
}

func (ve ValidationError) HasError() bool {
	return len(ve.errors) > 0
}

func (ve *ValidationError) Merge(otherErr error) {
	if otherErr != nil {
		if otherValidationError, ok := otherErr.(*ValidationError); ok {
			for key, msg := range otherValidationError.errors {
				ve.AddError(key, msg)
			}
		}
	}
}
