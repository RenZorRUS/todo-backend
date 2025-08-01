package utils

import (
	"math"

	"github.com/RenZorRUS/todo-backend/src/internal/errs"
)

type (
	Uint interface {
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
	}

	Int interface {
		~int8 | ~int16 | ~int32 | ~int64 | ~int
	}
)

func NilableUintToInt[I Uint, O Int](input *I) (*O, error) {
	if input == nil {
		return nil, nil //nolint:nilnil
	}

	output, err := UintToInt[I, O](*input)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func UintToInt[I Uint, O Int](value I) (O, error) { //nolint:ireturn
	var ouputType O

	threshold, err := getMaxSafeUintToIntConversionThreashold(ouputType)
	if err != nil {
		return 0, err
	}

	if uint64(value) > threshold {
		return 0, errs.ErrUintToIntOverflow
	}

	return O(value), nil
}

func getMaxSafeUintToIntConversionThreashold[T Int](outputType T) (uint64, error) {
	switch any(outputType).(type) {
	case int8:
		return uint64(math.MaxInt8), nil
	case int16:
		return uint64(math.MaxInt16), nil
	case int32:
		return uint64(math.MaxInt32), nil
	case int64:
		return uint64(math.MaxInt64), nil
	case int:
		return uint64(math.MaxInt), nil
	default:
		return 0, errs.ErrUnknownConversionOutputType
	}
}
