package client

import (
	"context"
	"errors"
	"fmt"
	"net"
)

// APIError is returned when the API responds with an error or the response can not be processed.
// A transport error (timeout, DNS, TLS, canceled context, ...) is kept in Err, so errors.Is and errors.As
// see it: errors.Is(err, context.DeadlineExceeded), errors.As(err, &netErr).
type APIError struct {
	// Message of the error.
	Message string

	// HTTPStatus of the response, 0 when the request was not sent.
	HTTPStatus int

	// Code is the API error code (the "code" field of the error response), nil when absent.
	Code *int

	// Body is the raw response body.
	Body []byte

	// Err is the underlying error (transport, encoding, decoding), nil for an error response of the API.
	Err error
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Code != nil {
		return fmt.Sprintf("%s (HTTP %d, code %d)", e.Message, e.HTTPStatus, *e.Code)
	}
	return fmt.Sprintf("%s (HTTP %d)", e.Message, e.HTTPStatus)
}

// ErrorCode returns the API error code, 0 when the response has none (Code is nil).
func (e *APIError) ErrorCode() int {
	if e.Code == nil {
		return 0
	}
	return *e.Code
}

// Unwrap returns the underlying error, for errors.Is and errors.As.
func (e *APIError) Unwrap() error {
	return e.Err
}

// Is reports a transport timeout as context.DeadlineExceeded also when it comes from http.Client.Timeout
// (net/http does not wrap context.DeadlineExceeded then): errors.Is(err, context.DeadlineExceeded) catches
// every timeout.
func (e *APIError) Is(target error) bool {
	if target != context.DeadlineExceeded {
		return false
	}
	var netErr net.Error
	return errors.As(e.Err, &netErr) && netErr.Timeout()
}
