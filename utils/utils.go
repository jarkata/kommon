package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func ReadFields(val interface{}) []any {
	fmt.Printf("ReadFields Val: %v\n", val)
	reflectVal := reflect.Indirect(reflect.ValueOf(val))
	fieldCount := reflectVal.NumField()
	data := make([]interface{}, 0)
	for index := 0; index < fieldCount; index++ {
		fieldVal := reflectVal.Field(index)
		fmt.Println(fieldVal.Interface())
		data = append(data, fieldVal.Interface())
		// fmt.Println("index:", index, fieldVal, fieldVal.CanSet())
		// fmt.Println("fieldVal.CanUint(): ", fieldVal.CanUint(),
		// 	",CanInterface:", fieldVal.CanInterface(),
		// 	",CanInt:", fieldVal.CanInt(),
		// 	",CanComplex:", fieldVal.CanComplex(),
		// 	",CanFloat:", fieldVal.CanFloat(), ",CanSet:", fieldVal.CanSet())
		// if CanString(fieldVal) {
		// 	fieldVal.SetString("fasdfasfsf")
		// }

		// fmt.Println("fieldVal.Type(): ", fieldVal.String(), fieldVal.Kind(), fieldVal.Type())
	}
	fmt.Println(data...)
	// fmt.Printf("reflectVal.NumField(): %v\n", reflectVal.NumField())
	fmt.Printf("reflect.ValueOf(val): %v\n", reflectVal)
	return nil
}

func CanString(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.String:
		return true
	default:
		return false
	}
}

func IsComplexType(fieldType string) bool {
	return strings.EqualFold("int64", fieldType)
}
