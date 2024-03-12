package strseg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	Parser[T any] func(string, string) (*T, error)
)

func CreateParser[T any]() Parser[T] {
	value := new(T)
	typeOfT := reflect.TypeOf(*value)
	fieldLen := typeOfT.NumField()
	return func(str string, sep string) (t *T, err error) {
		value := new(T)
		segments := strings.Split(str, sep)
		valueOfT := reflect.ValueOf(value)
		for i := 0; i < fieldLen; i++ {
			field := typeOfT.Field(i)
			tag := field.Tag.Get("index")
			if len(tag) == 0 {
				continue
			}
			if field.Type.Kind() != reflect.String {
				return nil, fmt.Errorf("expected string but found %s", field.Type.Kind())
			}
			index, err := strconv.Atoi(tag)
			if err != nil {
				return nil, err
			}
			if index >= len(segments) {
				return nil, fmt.Errorf("index out of range")
			}
			valueOfT.Elem().Field(i).Set(reflect.ValueOf(segments[index]))

		}
		return value, nil
	}
}
