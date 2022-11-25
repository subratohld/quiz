package errors

import "net/http"

// Bad request error
type BadRequest struct {
	msg string
}

func NewBadRequest(msg string) *BadRequest {
	return &BadRequest{
		msg: msg,
	}
}

func (b *BadRequest) Error() string {
	return b.msg
}

func (b *BadRequest) StatusCode() int {
	return http.StatusBadRequest
}

// Service error
type ServiceError struct {
	msg string
}

func NewServiceError(msg string) *ServiceError {
	return &ServiceError{
		msg: msg,
	}
}

func (b *ServiceError) Error() string {
	return b.msg
}

func (b *ServiceError) StatusCode() int {
	return http.StatusInternalServerError
}

// DB Error
type DBError struct {
	msg string
}

func NewDBError(msg string) *DBError {
	return &DBError{
		msg: msg,
	}
}

func (b *DBError) Error() string {
	return b.msg
}

func (b *DBError) StatusCode() int {
	return http.StatusInternalServerError
}

// Not found Error
type NotFoundError struct {
	msg string
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		msg: msg,
	}
}

func (b *NotFoundError) Error() string {
	return b.msg
}

func (b *NotFoundError) StatusCode() int {
	return http.StatusNotFound
}
