package handler

import (
	"errors"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrParamNotFound = errors.New("missing parameter")
	ErrConversion    = errors.New("failed to convert to requested type")
)

type parameters interface {
	map[string]string | map[string][]string | url.Values
}

type returnType interface {
	string |
		int | int64 | int32 |
		uint | uint64 | uint32 |
		float64 | float32 |
		uuid.UUID |
		bool |
		[]string |
		[]int | []int32 | []int64
}

// GetRequestParam is used to get parameters from known types that gorilla/mux uses.
// Returns defaultValue and err which can be ignored in case it's not important.
// nolint:funlen,gocognit,gocyclo,cyclop // no point in splitting up all the cases
func GetRequestParam[P parameters, T returnType](params P, key, separator string, defaultValue T) (T, error) {
	param, getParamErr := getParameter(params, key)
	if getParamErr != nil {
		return defaultValue, getParamErr
	}

	var ret T
	switch p := any(&ret).(type) {
	case *string:
		*p = param
	case *int64:
		v, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return defaultValue, fmt.Errorf("int64: %v error %v: %w", param, err, ErrConversion)
		}

		*p = v
	case *int:
		v, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return defaultValue, fmt.Errorf("int: %v error %v: %w", param, err, ErrConversion)
		}

		if v > math.MaxInt {
			return defaultValue, fmt.Errorf(
				"%w: cannot convert int64 to int. Out of bounds maxInt: %v, Value: %v",
				ErrConversion,
				math.MaxInt,
				v,
			)
		}

		*p = int(v)
	case *int32:
		v, err := strconv.ParseInt(param, 10, 32)
		if err != nil {
			return defaultValue, fmt.Errorf("int32: %v error %v: %w", param, err, ErrConversion)
		}

		*p = int32(v)
	case *uint:
		v, err := strconv.ParseUint(param, 10, 64)
		if err != nil {
			return defaultValue, fmt.Errorf("uint: %v error %v: %w", param, err, ErrConversion)
		}

		if v > math.MaxUint {
			return defaultValue, fmt.Errorf(
				"%w: cannot convert uint64 to uint. Out of bounds maxUInt: %v, Value: %v",
				ErrConversion,
				math.MaxUint,
				v,
			)
		}

		*p = uint(v)
	case *uint64:
		v, err := strconv.ParseUint(param, 10, 64)
		if err != nil {
			return defaultValue, fmt.Errorf("uint64: %v error %v: %w", param, err, ErrConversion)
		}

		*p = v
	case *uint32:
		v, err := strconv.ParseUint(param, 10, 32)
		if err != nil {
			return defaultValue, fmt.Errorf("uint32: %v error %v: %w", param, err, ErrConversion)
		}

		*p = uint32(v)
	case *float64:
		v, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return defaultValue, fmt.Errorf("float64: %v error %v: %w", param, err, ErrConversion)
		}

		*p = v
	case *float32:
		v, err := strconv.ParseFloat(param, 32)
		if err != nil {
			return defaultValue, fmt.Errorf("float32: %v error %v: %w", param, err, ErrConversion)
		}

		*p = float32(v)
	case *uuid.UUID:
		v, err := uuid.Parse(param)
		if err != nil {
			return defaultValue, fmt.Errorf("uuid: %v error %v: %w", param, err, ErrConversion)
		}

		*p = v
	case *bool:
		switch strings.ToUpper(param) {
		case "TRUE", "1", "T":
			*p = true
		case "FALSE", "0", "F":
			*p = false
		default:
			return defaultValue, fmt.Errorf("bool: %v is not true/t/1/false/f/0 (case insesitive): %w", param, ErrConversion)
		}
	case *[]string:
		*p = strings.Split(param, separator)
	case *[]int:
		values := strings.Split(param, separator)
		valuesInt := make([]int, len(values))

		for i := range values {
			i64, err := strconv.ParseInt(values[i], 10, 64)
			if err != nil {
				return defaultValue, fmt.Errorf("%w: cannot convert value: %v to []int", ErrConversion, param)
			}

			if i64 > math.MaxInt {
				return defaultValue, fmt.Errorf(
					"%w: cannot convert int64 to int. Out of bounds maxInt: %v, Value: %v",
					ErrConversion,
					math.MaxInt,
					i64,
				)
			}

			valuesInt[i] = int(i64)
		}

		*p = valuesInt
	case *[]int32:
		values := strings.Split(param, separator)
		valuesInt := make([]int32, len(values))

		for i := range values {
			i32, err := strconv.ParseInt(values[i], 10, 32)
			if err != nil {
				return defaultValue, fmt.Errorf("%w: cannot convert value: %v to []int", ErrConversion, param)
			}

			valuesInt[i] = int32(i32)
		}

		*p = valuesInt
	case *[]int64:
		values := strings.Split(param, separator)
		valuesInt := make([]int64, len(values))

		for i := range values {
			valuesInt[i], getParamErr = strconv.ParseInt(values[i], 10, 64)
			if getParamErr != nil {
				return defaultValue, fmt.Errorf("%w: cannot convert value: %v to []int", ErrConversion, param)
			}
		}

		*p = valuesInt
	}

	return ret, nil
}

func getParameter[P parameters](params P, key string) (param string, err error) {
	var ok bool
	switch a := any(params).(type) {
	case map[string]string:
		param, ok = a[key]
	case map[string][]string:
		if _, ok = a[key]; !ok || len(a[key]) == 0 {
			ok = false
			break
		}

		param = a[key][0]
	case url.Values:
		if _, ok = a[key]; !ok || len(a[key]) == 0 {
			ok = false
			break
		}

		param = a[key][0]
	}

	if !ok {
		err = fmt.Errorf("%w: cannot find key: %v", ErrParamNotFound, key)
	}

	return param, err
}
