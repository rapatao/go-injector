# go-injector

Provides a dependency injection mechanism for Golang.

## Documentation

### Installing

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```go
import "github.com/rapatao/go-injector"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the injector package:

```shell
go get -u github.com/rapatao/go-injector
```

### Using Injector

First you need to import Injector package for using it, one simplest example likes the follow `example.go`:

```go
package main

import (
	"fmt"
	"github.com/rapatao/go-injector"
)

type A struct {
	name string
	b    B
}
type B struct {
	name string
}

func (a *A) Run() {
	fmt.Printf("%s, %s", a.name, a.b.name)
}

func (a *A) Initialize(container *injector.Container) error {
	a.name = "hello"

	var b B
	err := container.Get(&b)
	a.b = b

	return err
}

func (b *B) Initialize(_ *injector.Container) error {
	b.name = "world"
	return nil
}

func main() {
	container := injector.Create()

	var a A
	err := container.Get(&a)
	if err != nil {
		panic(err)
	}

	a.Run()
}
```

And use the Go command to run the demo:

```shell
go run example.go
```

# Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.
