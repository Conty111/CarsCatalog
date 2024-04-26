package external_api

import "fmt"

type ExternalAPIError struct {
	Description string
}

func (e *ExternalAPIError) Error() string {
	return e.Description
}
func NewExternalAPIError(err error) error {
	return &ExternalAPIError{
		Description: fmt.Sprintf("external API error: %s", err.Error()),
	}
}
