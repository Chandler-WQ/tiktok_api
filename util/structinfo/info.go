package structinfo

import (
	"reflect"

	"github.com/spf13/cast"
)

func ToSliceName(obj interface{}) []string {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v)
	t := v.Type()
	fields := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			fields = walkStructName(field.Type, fields)
		} else {
			fields = append(fields, cast.ToString(field.Name))
		}
	}
	return fields
}

func walkStructName(t reflect.Type, fields []string) []string {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			fields = walkStructName(field.Type, fields)
		} else {
			fields = append(fields, cast.ToString(field.Name))
		}
	}
	return fields
}

func ToSlice(obj interface{}) []string {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v)
	t := v.Type()
	fields := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i)
		if value.Kind() == reflect.Struct {
			fields = walkStruct(value, fields)
		} else {
			fields = append(fields, cast.ToString(value.Interface()))
		}
	}
	return fields
}

func walkStruct(v reflect.Value, fields []string) []string {
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)

		if value.Kind() == reflect.Struct {
			fields = walkStruct(value, fields)
		} else {
			fields = append(fields, cast.ToString(value.Interface()))
		}
	}
	return fields
}
