package wraperr

import (
	"errors"
	"strings"
)

// MultiError combines multiple errors into one.
// The Error method returns the strings from the individual Error methods
// joined by the new line character '\n'.
// The motivation behind Combine and MultiError is to combine different
// logical errors into one, as compared to error wrapping,
// which adds more information to one logical error.
type MultiError interface {
	Error() string
	Errors() []error
}

type multiError []error

// Combine returns a MultiError error for 2 or more errors which are not nil,
// or the the single error if only one error was passed,
// or nil if zero arguments are passed or all passed errors are nil.
// The MultiError type's Error method returns the strings from the
// individual Error methods joined by the new line character '\n'.
// In case of a MultiError, errors.Is and errors.As will return true
// for the first matched error.
// Combine does not wrap the passed errors with a text or call stack.
// The motivation behind Combine and MultiError is to combine different
// logical errors into one, as compared to error wrapping,
// which adds more information to one logical error.
func Combine(errs ...error) error {
	var combined multiError
	for _, err := range errs {
		if err != nil {
			var m multiError
			if errors.As(err, &m) {
				combined = append(combined, m.Errors()...)
			} else {
				combined = append(combined, err)
			}
		}
	}

	switch len(combined) {
	case 0:
		return nil
	case 1:
		return combined[0]
	default:
		return combined
	}
}

// Uncombine returns multiple errors if err is a MultiError,
// else it will return a single element slice containing err
// or nil if err is nil.
func Uncombine(err error) []error {
	if err == nil {
		return nil
	}
	var multi MultiError
	if errors.As(err, &multi) {
		return multi.Errors()
	}
	return []error{err}
}

func (m multiError) Error() string {
	var b strings.Builder
	for i, err := range m {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(err.Error())
	}
	return b.String()
}

func (m multiError) Errors() []error {
	return m
}

func (m multiError) Is(target error) bool {
	for _, err := range m {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func (m multiError) As(target interface{}) bool {
	for _, err := range m {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}
