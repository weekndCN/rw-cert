package handler

var (
	// ErrNotFound not founded
	ErrNotFound = NewError("Not Found")
)

// Error a json-encoded api error
type Error struct {
	Message string `json:"message"`
}

// Error error method
func (e *Error) Error() string {
	return e.Message
}

// NewError return a new error message
func NewError(text string) error {
	return &Error{Message: text}
}
