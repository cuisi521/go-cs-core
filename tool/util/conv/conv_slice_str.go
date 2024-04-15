// package conv
// @Author cuisi
// @Date 2024/4/8 10:28:00
// @Desc
package conv

import (
	"reflect"

	"github.com/cuisi521/go-cs-core/tool/json"
)

// SliceStr is alias of Strings.
func SliceStr(any interface{}) []string {
	return Strs(any)
}

// Strs converts `any` to []string.
func Strs(any interface{}) []string {
	if any == nil {
		return nil
	}
	var (
		array []string = nil
	)
	switch value := any.(type) {
	case []int:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []int8:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []int16:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []int32:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []int64:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []uint:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []uint8:
		if json.Valid(value) {
			_ = json.UnmarshalUseNumber(value, &array)
		} else {
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = Str(v)
			}
		}
	case []uint16:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []uint32:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []uint64:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []bool:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []float32:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []float64:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []interface{}:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	case []string:
		array = value
	case [][]byte:
		array = make([]string, len(value))
		for k, v := range value {
			array[k] = Str(v)
		}
	}
	if array != nil {
		return array
	}
	if v, ok := any.(iStrings); ok {
		return v.Strings()
	}
	if v, ok := any.(iInterfaces); ok {
		return Strs(v.Interfaces())
	}
	// JSON format string value converting.
	if checkJsonAndUnmarshalUseNumber(any, &array) {
		return array
	}

	var (
		refValue reflect.Value
		refKind  reflect.Kind
	)
	if v, ok := any.(reflect.Value); ok {
		refValue = v
	} else {
		refValue = reflect.ValueOf(any)
	}
	refKind = refValue.Kind()
	for refKind == reflect.Ptr {
		refValue = refValue.Elem()
		refKind = refValue.Kind()
	}

	switch refKind {
	case reflect.Slice, reflect.Array:
		var (
			length = refValue.Len()
			slice  = make([]string, length)
		)
		for i := 0; i < length; i++ {
			slice[i] = Str(refValue.Index(i).Interface())
		}
		return slice

	default:
		if refValue.IsZero() {
			return []string{}
		}
		return []string{Str(any)}
	}
}
