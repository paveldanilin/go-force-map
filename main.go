package main

import (
	"fmt"
	"reflect"
)

func ForceMap(v interface{}) interface{} {
	arrIter := func (t interface{}) interface{} {
		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice, reflect.Array:
			out := make(map[int]interface{})
			s := reflect.ValueOf(t)
			for i := 0 ; i < s.Len(); i++ {
				out[i] = ForceMap(s.Index(i).Interface())
			}
			return out
		default:
			return t
		}
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		return arrIter(v)
	case reflect.Map:
		m := reflect.ValueOf(v)
		converted := make(map[string]interface{}, len(m.MapKeys()))
		for _, key := range m.MapKeys() {
			converted[fmt.Sprintf("%v", key.Interface())] = ForceMap(m.MapIndex(key).Interface())
		}
		return converted
	case reflect.Struct:
		s := reflect.ValueOf(v)
		converted := make(map[string]interface{}, s.NumField())
		for i := 0 ; i < s.NumField(); i++ {
			converted[s.Type().Field(i).Name] = ForceMap(s.Field(i).Interface())
		}
		return converted
	default:
		return v
	}
}

func main() {
}
