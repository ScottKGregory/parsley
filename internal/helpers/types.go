// Package helpers is a bunch of useful functions
package helpers

import (
	"fmt"
	"strings"
)

// TypesMatch check if the types of the two values is the same
func TypesMatch(a, b any) bool {
	return fmt.Sprintf("%T", a) == fmt.Sprintf("%T", b)
}

func ToFloat64(input any) (float64, error) {
	switch x := input.(type) {
	case int:
		return float64(x), nil
	case int8:
		return float64(x), nil
	case int16:
		return float64(x), nil
	case int32:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case uint:
		return float64(x), nil
	case uint8:
		return float64(x), nil
	case uint16:
		return float64(x), nil
	case uint32:
		return float64(x), nil
	case uint64:
		return float64(x), nil
	case float32:
		return float64(x), nil
	case float64:
		return float64(x), nil
	}

	return 0, fmt.Errorf("cannot convert to float64: %T", input)
}

func ToBool(e any) (bool, error) {
	switch x := e.(type) {
	case bool:
		return x, nil
	case uint8:
		return x > 0, nil
	case uint16:
		return x > 0, nil
	case uint32:
		return x > 0, nil
	case uint64:
		return x > 0, nil
	case int8:
		return x > 0, nil
	case int16:
		return x > 0, nil
	case int32:
		return x > 0, nil
	case int64:
		return x > 0, nil
	case float32:
		return x > 0, nil
	case float64:
		return x > 0, nil
	case string:
		switch strings.ToLower(x) {
		case "yes", "true", "y", "1", "yarp":
			return true, nil
		case "no", "false", "n", "0", "narp":
			return false, nil
		}

		return false, fmt.Errorf("string '%s' could not be parsed as a bool", x)
	}
	return false, fmt.Errorf("unrecognised type produced by matcher: %T", e)
}
