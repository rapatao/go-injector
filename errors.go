package injector

import (
	"errors"
)

var (
	// ErrNilPointer is returned when a non-nil pointer is used when invoking a function.
	ErrNilPointer = errors.New("argument should be a non-nil pointer")

	// ErrNilValue is returned when a nil argument is used when invoking a function.
	ErrNilValue = errors.New("value should be non-nil")

	// ErrInitializingType is returned when a failure happens during the type initialization.
	ErrInitializingType = errors.New("failed to initialize type")
)
