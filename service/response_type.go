package service

// ResponseType provides an enumeration of service response types
type ResponseType int

const (
	// InvalidData response
	InvalidData ResponseType = iota

	// Error response
	Error

	// Conflict response
	Conflict

	// NotFound response
	NotFound

	// Success response
	Success
)

var values = [...]string{
	"invalid-data",
	"error",
	"conflict",
	"not-found",
	"success",
}

// String representation of `ResponseType`
func (a ResponseType) String() string {
	return values[a]
}
