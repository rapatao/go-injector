package injector

import (
	"errors"
)

var (
	// ErrNilPointer is returned when a non-nil pointer is used when invoking a function.
	ErrNilPointer = errors.New("argument should be a non-nil pointer")

	// ErrNilValue is returned when a nil argument is used when invoking a function.
	ErrNilValue = errors.New("value should be non-nil")

	// ErrNotInitializable is returned when a struct is not present in the injector context,
	// and it did not implement injector.Injectable.
	ErrNotInitializable = errors.New("unable to initialize the type")

	// ErrInitializing is returned when a failure happens during the struct initialization.
	ErrInitializing = errors.New("failed to initialize struct")
)
