package injector_test

import (
	"errors"
	"fmt"

	"github.com/rapatao/go-injector"
)

var ErrDummy = errors.New("awesome dummy error")

// NotAutoConfigurable ...
type NotAutoConfigurable struct {
	Init bool
}

// AutoConfigurable ...
type AutoConfigurable struct {
	Init bool
}

func (a *AutoConfigurable) Initialize(_ *injector.Container) error {
	a.Init = true

	return nil
}

// AutoConfigurableWithDep ...
type AutoConfigurableWithDep struct {
	NotAutoConfigurable NotAutoConfigurable
	AutoConfigurable    AutoConfigurable
}

func (a *AutoConfigurableWithDep) Initialize(container *injector.Container) error {
	err := container.Get(&a.AutoConfigurable)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = container.Get(&a.NotAutoConfigurable)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// FailureAutoConfigurable ...
type FailureAutoConfigurable struct{}

func (a *FailureAutoConfigurable) Initialize(_ *injector.Container) error {
	return ErrDummy
}

var (
	_ injector.Injectable = (*AutoConfigurable)(nil)
	_ injector.Injectable = (*AutoConfigurableWithDep)(nil)
	_ injector.Injectable = (*FailureAutoConfigurable)(nil)
)
