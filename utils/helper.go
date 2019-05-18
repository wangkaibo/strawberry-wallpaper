package utils

import (
	"fmt"
	"reflect"
)

func StructToMap(i interface{}) map[string]interface{} {
	objT := reflect.TypeOf(i)
	objV := reflect.ValueOf(i)
	fmt.Println(objT.Kind())
	if objT.Kind() != reflect.Struct {
		return map[string]interface{}{}
	}
	resMap := map[string]interface{}{}
	for i := 0; i < objT.NumField(); i++ {
		fieldType := objT.Field(i)
		fieldValue := objV.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		if len(jsonTag) != 0 {
			resMap[jsonTag] = fieldValue.Interface()
		} else {
			resMap[fieldType.Name] = fieldValue.Interface()
		}
	}

	return resMap
}

func SliceKeyBy(s []map[string]string, key string) map[string]interface{} {
	res := make(map[string]interface{}, len(s))
	for _, item := range(s) {
		if _, ok := item[key]; ok {
			res[item[key]] = item
		}
	}
	return res
}