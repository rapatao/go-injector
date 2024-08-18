package injector_test

import (
	"testing"

	"github.com/rapatao/go-injector"

	"github.com/stretchr/testify/assert"
)

func TestRegister_failures(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Register(nil)
	assert.ErrorIs(t, err, injector.ErrNilValue)

	var nilInstance *NotAutoConfigurable
	err = container.Register(nilInstance)
	assert.ErrorIs(t, err, injector.ErrNilPointer)
}

func TestRegister_manual(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Register(NotAutoConfigurable{})
	assert.NoError(t, err)
}

func TestRegister_auto(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Register(&NotAutoConfigurable{})
	assert.NoError(t, err)
}

func TestRegister_auto_ptr(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Register(AutoConfigurable{})
	assert.NoError(t, err)
}

func TestRegister_auto_deps(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Register(AutoConfigurableWithDep{})
	assert.NoError(t, err)
}

func TestGet_failures(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Get(NotAutoConfigurable{})
	assert.ErrorIs(t, err, injector.ErrNilPointer)

	var nilPtr *NotAutoConfigurable
	err = container.Get(&nilPtr)
	assert.ErrorIs(t, err, injector.ErrNilPointer)

	var notAc NotAutoConfigurable
	err = container.Get(&notAc)
	assert.ErrorIs(t, err, injector.ErrInitializingType)

	var failAc FailureAutoConfigurable
	err = container.Get(&failAc)
	assert.Error(t, err)

	var dep AutoConfigurableWithDep
	err = container.Get(&dep)
	assert.ErrorIs(t, err, injector.ErrInitializingType)
}

func TestGet_auto(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	var dep AutoConfigurable
	err := container.Get(&dep)
	assert.NoError(t, err)
	assert.True(t, dep.Init)
}

func TestGet_manual(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Register(&NotAutoConfigurable{Init: true})
	assert.NoError(t, err)

	var dep NotAutoConfigurable
	err = container.Get(&dep)
	assert.NoError(t, err)
	assert.True(t, dep.Init)
}

func TestGet_auto_deps(t *testing.T) {
	t.Parallel()

	container := injector.Create()

	err := container.Register(&NotAutoConfigurable{Init: true})
	assert.NoError(t, err)

	var dep AutoConfigurableWithDep
	err = container.Get(&dep)
	assert.NoError(t, err)
	assert.True(t, dep.AutoConfigurable.Init)
}
