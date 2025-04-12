// Package helpers is a bunch of useful functions
package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// TypesMatch check if the types of the two values are the same
func TypesMatch(a, b any) bool {
	return fmt.Sprintf("%T", a) == fmt.Sprintf("%T", b)
}

// ToFloat64 attempts to convert the input value in to a float64. It will cast int/uint/flaot types, and attempt to parse strings as floats
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
	case string:
		return strconv.ParseFloat(x, 64)
	}

	return 0, fmt.Errorf("cannot convert to float64: %T", input)
}

// ToBool converts the input value in to a bool.
//
// - If the type is a number then any value over 0 will return true
// - Strings will be checked against known values
// - Strings not matching a known value will attempt to parse as a float
func ToBool(e any) (bool, error) {
	switch x := e.(type) {
	case bool:
		return x, nil
	case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		i, err := ToFloat64(x)
		if err != nil {
			return false, err
		}

		return i > 0, nil
	case string:
		switch strings.ToLower(x) {
		case "yes", "true", "y", "1", "yarp":
			return true, nil
		case "no", "false", "n", "0", "narp":
			return false, nil
		}

		i, err := ToFloat64(x)
		if err == nil {
			return i > 0, nil
		}

		return false, fmt.Errorf("string '%s' could not be parsed as a bool", x)
	}
	return false, fmt.Errorf("unrecognised type provided to bool: %T", e)
}

// ToString converts the input to a string
func ToString(e any) (string, error) {
	return fmt.Sprint(e), nil
}
