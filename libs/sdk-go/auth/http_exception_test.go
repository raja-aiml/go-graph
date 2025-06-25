package auth

import "testing"

func TestNewHTTPExceptionDefault(t *testing.T) {
    err := NewHTTPException(401, "", nil)
    if err.Status != 401 {
        t.Errorf("expected status 401, got %d", err.Status)
    }
    if err.Detail != "Unauthorized" {
        t.Errorf("unexpected detail: %s", err.Detail)
    }
    if err.Error() != "401: Unauthorized" {
        t.Errorf("unexpected error string: %s", err.Error())
    }
    if err.Headers == nil || len(err.Headers) != 0 {
        t.Errorf("expected empty headers")
    }
}

func TestNewHTTPExceptionCustom(t *testing.T) {
    headers := map[string]string{"X-Test": "1"}
    err := NewHTTPException(404, "Not found", headers)
    if err.Status != 404 || err.Detail != "Not found" {
        t.Errorf("unexpected err fields: %+v", err)
    }
    if err.Headers["X-Test"] != "1" {
        t.Errorf("unexpected headers: %+v", err.Headers)
    }
}
