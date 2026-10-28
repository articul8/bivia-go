package bivia

import "fmt"

// APIError represents an error returned by the Bivia API.
type APIError struct {
	// Status is the HTTP status code.
	Status int
	// Body is the decoded response body, if available.
	Body any
	// Message is a human-readable error description.
	Message string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("bivia: %d %s", e.Status, e.Message)
	}
	return fmt.Sprintf("bivia: HTTP %d", e.Status)
}
