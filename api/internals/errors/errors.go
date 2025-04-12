package errors

import (
	"net/http"
)

type ErrorType string

const (
	ErrorTypeValidation     ErrorType = "VALIDATION_ERROR"
	ErrorTypeDatabase       ErrorType = "DATABASE_ERROR"
	ErrorTypeNotFound       ErrorType = "NOT_FOUND"
	ErrorTypeUnauthorized   ErrorType = "UNAUTHORIZED"
	ErrorTypeBadRequest     ErrorType = "BAD_REQUEST"
	ErrorTypeInternalServer ErrorType = "INTERNAL_SERVER_ERROR"
	ErrorTypeRateLimit      ErrorType = "RATE_LIMIT_ERROR"
	ErrorTypeForbidden      ErrorType = "FORBIDDEN"
)

type APIError struct {
	Type       ErrorType `json:"type"`
	Message    string    `json:"message"`
	StatusCode int       `json:"status_code"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewValidationError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeValidation,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewDatabaseError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeDatabase,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeNotFound,
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func NewBadRequestError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeBadRequest,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeInternalServer,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewRateLimitError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeRateLimit,
		Message:    message,
		StatusCode: http.StatusTooManyRequests,
	}
}

func NewUnauthorizedError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeUnauthorized,
		Message:    message,
		StatusCode: http.StatusUnauthorized}
}

func NewAuthenticationError(message string, args ...any) *APIError {
	return &APIError{
		Type:       ErrorTypeUnauthorized,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func NewConflictError(message string, args ...any) *APIError {
	return &APIError{
		Type:       ErrorTypeBadRequest,
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}

func NewForbiddenError(message string, args ...any) *APIError {
	return &APIError{
		Type:       ErrorTypeForbidden,
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}
