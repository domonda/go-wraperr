package sentinel

// Error implements the error interface for a string
// and is meant to be used to declare const sentinel errors.
//
// Example:
//   const ErrUserNotFound = sentinel.Error("user not found")
type Error string

func (s Error) Error() string {
	return string(s)
}
