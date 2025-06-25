package auth

import (
    "fmt"
    "net/http"
)

// HTTPException represents an HTTP error with a status code, message and optional headers.
type HTTPException struct {
    Status  int
    Detail  string
    Headers map[string]string
}

// NewHTTPException creates a new HTTPException. If detail is empty, it uses the
// standard HTTP status text for the status code. If headers is nil, an empty map
// is created.
func NewHTTPException(status int, detail string, headers map[string]string) *HTTPException {
    if detail == "" {
        detail = http.StatusText(status)
        if detail == "" {
            detail = "Unknown error"
        }
    }
    if headers == nil {
        headers = map[string]string{}
    }
    return &HTTPException{Status: status, Detail: detail, Headers: headers}
}

// Error implements the error interface.
func (e *HTTPException) Error() string {
    return fmt.Sprintf("%d: %s", e.Status, e.Detail)
}
