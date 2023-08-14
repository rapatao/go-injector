package injector

// Injectable defines expected method when using the autoconfiguration.
type Injectable interface {
	// Initialize defines the values for a struct, or returns a non nil error in case of failure.
	//
	// When a struct needs other structs to be created, it can access them by using injector.Get function.
	Initialize(container *Container) error
}
