package terminal

import (
	"encoding/json"
	"fmt"
)

// APIResponse represents the standard API response structure returned by Terminal Africa.
type APIResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

// APIError represents an error returned by the Terminal Africa API or HTTP transport.
type APIError struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message"`
	Status     bool   `json:"status"`
	Raw        []byte `json:"-"`
}

func (e *APIError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("terminal africa api error (status %d): %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("terminal africa api error: %s", e.Message)
}

// PaginationMeta contains pagination details returned in list endpoints.
type PaginationMeta struct {
	Page      int `json:"page,omitempty"`
	PerPage   int `json:"perPage,omitempty"`
	Total     int `json:"total,omitempty"`
	PageCount int `json:"pageCount,omitempty"`
}

// RawResponse is a helper type for unmarshaling dynamic json responses.
type RawResponse = APIResponse[json.RawMessage]
