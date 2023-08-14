package injector

import (
	"fmt"
	"reflect"
)

func typeName(val reflect.Value) string {
	t := val.Type()

	return fmt.Sprintf("%s/%s", t.PkgPath(), t.Name())
}

func concreteValueFrom(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr && !value.IsNil() {
		return concreteValueFrom(value.Elem())
	}

	return value
}
