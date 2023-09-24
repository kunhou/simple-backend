package reason

// Reason type represents different types of error codes.
type Reason string

const (
	// Success represents a successful operation.
	Success Reason = "base.success"

	// Unknown represents an unknown error.
	Unknown Reason = "error.unknown"

	// InternalServer represents an internal server error.
	InternalServer Reason = "error.internal_server"

	// InvalidRequest represents an invalid request error.
	InvalidRequest Reason = "error.invalid_request"

	// InvalidValue represents an invalid value error.
	InvalidValue Reason = "error.invalid_value"

	// NotFound represents a not found error.
	NotFound Reason = "error.not_found"

	// Unauthorized represents an unauthorized error.
	Unauthorized Reason = "error.unauthorized"

	// Forbidden represents a forbidden error.
	Forbidden Reason = "error.forbidden"

	// Duplicate represents a duplicate object.
	Duplicate Reason = "error.duplicate"

	// DatabaseError represents a database error.
	DatabaseError Reason = "error.database"

	DatabaseConnectionFailed = "error.database.connection_failed"
)
