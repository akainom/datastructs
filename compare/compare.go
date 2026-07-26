package compare

import (
	"reflect"
)

func DefaultCompare[T any]() func(a, b T) bool {
	t := reflect.TypeFor[T]()

	switch t.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.Ptr:
		return func(a, b T) bool {
			return any(a) == any(b)
		}

	default:
		return func(a, b T) bool {
			return reflect.DeepEqual(a, b)
		}
	}
}

func BasicComparerAny[T any](a, b T) bool {
	return any(a) == any(b)
}
