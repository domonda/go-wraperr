package wraperr

// Sentinel implements the error interface for a string
// and is meant to be used to declare const sentinel errors.
//
// Example:
//   const ErrUserNotFound = wraperr.Sentinel("user not found")
type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}
