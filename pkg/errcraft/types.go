package errcraft

import (
	"github/kunhou/simple-backend/pkg/reason"
	"net/http"
)

// InternalServer new InternalServer error
func InternalServer(reason reason.Reason) *Error {
	return New(500, reason)
}

// IsInternalServer determines if err is InternalServer error.
func IsInternalServer(err *Error) bool {
	return err.Code == 500
}

// NotFound new NotFound error
func NotFound(reason reason.Reason) *Error {
	return New(404, reason)
}

// IsNotFound determines if err is NotFound error.
func IsNotFound(err *Error) bool {
	return err.Code == 404
}

// DuplicateKey new DuplicateKey error
func DuplicateKey(reason reason.Reason) *Error {
	return New(http.StatusConflict, reason)
}

// IsDuplicateKey determines if err is DuplicateKey error.
func IsDuplicateKey(err *Error) bool {
	return err.Code == 409
}
