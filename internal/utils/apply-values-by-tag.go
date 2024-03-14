package utils

import (
	"reflect"
	"strconv"
)

func FindAndApplyValuesByTag(values map[string]string, obj reflect.Value, tagName string, depth int) {
	if depth > 10 {
		return
	}
	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)

		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem()))
			}
			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			FindAndApplyValuesByTag(values, field, tagName, depth+1)
			continue
		}

		tagValue := obj.Type().Field(i).Tag.Get(tagName)
		if tagValue == "" || tagValue == "-" {
			continue
		}

		if value, ok := values[tagValue]; ok {
			if !field.IsValid() || !field.CanSet() {
				continue
			}

			parseAndSetFieldValue(field, value)
		}
	}
}

func parseAndSetFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int32, reflect.Int64:
		valueAsInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return
		}
		field.SetInt(valueAsInt)
	case reflect.Float32, reflect.Float64:
		valueAsFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		field.SetFloat(valueAsFloat)
	case reflect.Bool:
		valueAsBool, err := strconv.ParseBool(value)
		if err != nil {
			return
		}
		field.SetBool(valueAsBool)
	}
}
