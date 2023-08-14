package injector

import (
	"fmt"
	"reflect"
	"sync"
)

// Container stores all automatic and manually registered struct, using them when new instances are required.
type Container struct {
	instances sync.Map
}

// Create return an injector.Container to be used to register and get structs.
func Create() *Container {
	return &Container{instances: sync.Map{}}
}

// Register is used to manually add structs to the injection context.
// It is not needed when the struct implements injector.Injectable.
func (c *Container) Register(instance interface{}) error {
	if instance == nil {
		return ErrNilValue
	}

	value := concreteValueFrom(reflect.ValueOf(instance))

	if value.Kind() == reflect.Ptr && value.IsNil() {
		return ErrNilPointer
	}

	name := typeName(value)

	c.instances.Store(name, instance)

	return nil
}

// Get is used to get a struct from the injector context.
//
// If it isn't present in the context and it implements injector.Injectable,
// the Initialize method will be invoked, returning the given error in case of failure.
//
// If it isn't present and do not implement the injector.Injectable,
// the error injector.ErrNotInitializable will be returned.
//
// The argument must be a non-nil pointer, or an error injector.ErrNilPointer is returned.
func (c *Container) Get(desired interface{}) error {
	value := reflect.ValueOf(desired)

	val := concreteValueFrom(value)
	if value.Kind() != reflect.Ptr || (val.Kind() == reflect.Ptr && val.IsNil()) {
		return ErrNilPointer
	}

	instance, err := c.loadInstance(typeName(val), desired)
	if err != nil {
		return err
	}

	val.Set(concreteValueFrom(reflect.ValueOf(instance)))

	return nil
}

func (c *Container) loadInstance(name string, desired interface{}) (interface{}, error) {
	if instance, ok := c.instances.Load(name); ok {
		return instance, nil
	}

	err := c.autoConfigure(name, desired)
	if err != nil {
		return nil, err
	}

	_ = c.Register(desired)

	return c.loadInstance(name, desired)
}

func (c *Container) autoConfigure(name string, desired interface{}) error {
	if v, ok := desired.(Injectable); ok {
		err := v.Initialize(c)
		if err != nil {
			return fmt.Errorf("%w, type: %s", err, name)
		}

		return nil
	}

	return fmt.Errorf("%w, type: \"%s\"", ErrNotInitializable, name)
}
