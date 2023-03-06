package utils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jarkata/kommon/utils"
)

func TestIsComplex(t *testing.T) {

	fmt.Printf("utils.IsComplexType(\"int64\"): %v\n", utils.IsComplexType("int64"))
}

func TestReadFields(t *testing.T) {
	val := Users{Username: "test", Password: "123", CreateTime: time.Now()}
	utils.ReadFields(&val)
	// vals := reflect.ValueOf(&val)
	// fmt.Printf("vals: %v\n", vals)
	// fmt.Println("val:", val, reflect.ValueOf(&val))
	// reflectVal := reflect.Indirect(reflect.ValueOf(&val))
	// fmt.Printf("reflectVal.NumField(): %v\n", reflectVal.NumField())
	// fmt.Printf("reflectVal.Len(): %v\n", reflectVal)
	// fieldCount := reflectVal.NumField()
	// for index := 0; index < fieldCount; index++ {
	// 	fieldVal := reflectVal.Field(index)
	// 	fmt.Println("index:", index, fieldVal, fieldVal.CanSet())
	// 	fmt.Println("fieldVal.CanUint(): ", fieldVal.CanUint(),
	// 		",CanInterface:", fieldVal.CanInterface(),
	// 		",CanInt:", fieldVal.CanInt(),
	// 		",CanComplex:", fieldVal.CanComplex(),
	// 		",CanFloat:", fieldVal.CanFloat(), ",CanSet:", fieldVal.CanSet())
	// 	fmt.Println("fieldVal.Type(): ", fieldVal.String(), fieldVal.Kind(), fieldVal.Type())
	// }
	// fmt.Printf("reflectVal.NumField(): %v\n", reflectVal.NumField())
	// fmt.Printf("reflect.ValueOf(val): %v\n", reflect.ValueOf(val))
}

type Users struct {
	Username   string
	Password   string
	CreateTime time.Time
	Times      int64
	Status     Status
	Blanace    float64
	TestField  uint64
}

type Status int
