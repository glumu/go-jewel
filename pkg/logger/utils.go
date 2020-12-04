package logger

import (
	"reflect"
	"strings"
)

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		dataKey := t.Field(i).Tag.Get("json")
		if dataKey == "" {
			dataKey = SnakeString(t.Field(i).Name)
		}
		dataValue := t.Field(i).Tag.Get("default")
		if dataValue == "" {
			dataValue = v.Field(i).Interface().(string)
		}
		data[dataKey] = dataValue
	}
	return data
}
