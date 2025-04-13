package helpers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestTypesMatch(t *testing.T) {
	testCases := []struct {
		a, b   any
		result bool
	}{
		{int(1), int(2), true},
		{"string", "int(2)", true},
		{float64(1), float64(2), true},
		{float32(1), float32(2), true},
		{int(1), float32(2), false},
		{float64(1), float32(2), false},
		{float64(1), "string", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%T_%T", tc.a, tc.b), func(t *testing.T) {
			assert.Equal(t, tc.result, TypesMatch(tc.a, tc.b))
		})
	}
}

func TestToFloat64(t *testing.T) {
	testCases := []struct {
		a      any
		result float64
		err    error
	}{
		{int(12), float64(12), nil},
		{int8(12), float64(12), nil},
		{int16(12), float64(12), nil},
		{int32(12), float64(12), nil},
		{int64(12), float64(12), nil},
		{uint(12), float64(12), nil},
		{uint8(12), float64(12), nil},
		{uint16(12), float64(12), nil},
		{uint32(12), float64(12), nil},
		{uint64(12), float64(12), nil},
		{float32(12), float64(12), nil},
		{float64(12), float64(12), nil},
		{"12", float64(12), nil},

		{int(33), float64(33), nil},
		{int8(33), float64(33), nil},
		{int16(33), float64(33), nil},
		{int32(33), float64(33), nil},
		{int64(33), float64(33), nil},
		{uint(33), float64(33), nil},
		{uint8(33), float64(33), nil},
		{uint16(33), float64(33), nil},
		{uint32(33), float64(33), nil},
		{uint64(33), float64(33), nil},
		{float32(33.33333), float64(33.33332824707031), nil},
		{float64(33.33333), float64(33.33333), nil},
		{"33", float64(33), nil},

		{"NOT_A_NUMBER", 0, errors.New("error parsing value as float, could not parse string 'NOT_A_NUMBER'")},
		{true, 0, errors.New("error parsing value as float, invalid type: bool")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%T_%v", tc.a, tc.a), func(t *testing.T) {
			result, err := ToFloat64(tc.a)
			assert.Equal(t, tc.result, result)
			assert.ErrorEqual(t, tc.err, err)
			if tc.err != nil {
				assert.ErrorIs(t, ErrInvalidFloat, err)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	testCases := []struct {
		a      any
		result bool
		err    error
	}{
		{int(12), true, nil},
		{int8(12), true, nil},
		{int16(12), true, nil},
		{int32(12), true, nil},
		{int64(12), true, nil},
		{uint(12), true, nil},
		{uint8(12), true, nil},
		{uint16(12), true, nil},
		{uint32(12), true, nil},
		{uint64(12), true, nil},
		{float32(12), true, nil},
		{float64(12), true, nil},
		{true, true, nil},

		{int(0), false, nil},
		{int8(0), false, nil},
		{int16(0), false, nil},
		{int32(0), false, nil},
		{int64(0), false, nil},
		{uint(0), false, nil},
		{uint8(0), false, nil},
		{uint16(0), false, nil},
		{uint32(0), false, nil},
		{uint64(0), false, nil},
		{float32(0), false, nil},
		{float64(0), false, nil},
		{false, false, nil},

		{"yes", true, nil},
		{"true", true, nil},
		{"y", true, nil},
		{"1", true, nil},
		{"yarp", true, nil},

		{"no", false, nil},
		{"false", false, nil},
		{"n", false, nil},
		{"0", false, nil},
		{"narp", false, nil},

		{"0", false, nil},
		{"0.00", false, nil},
		{"12", true, nil},
		{"1.2", true, nil},
		{"0.4", true, nil},

		{"asdfouhasdf", false, errors.New("error parsing value as bool, could not parse string 'asdfouhasdf'")},
		{[]string{"foo", "bar"}, false, errors.New("error parsing value as bool, invalid type: []string")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%T_%v", tc.a, tc.a), func(t *testing.T) {
			result, err := ToBool(tc.a)
			assert.Equal(t, tc.result, result)
			assert.ErrorEqual(t, tc.err, err)
			if tc.err != nil {
				assert.ErrorIs(t, ErrInvalidBool, err)
			}
		})
	}
}
