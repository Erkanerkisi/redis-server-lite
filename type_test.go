package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_types(t *testing.T) {
	var str = "Test"
	var bl = false
	var number = 12
	var number2 int64 = 1
	var interfaceAmaString interface{} = "Test"
	var arrStrings = []string{"1", "2"}
	var arrMixed = []interface{}{"1", true, 123}

	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(bl))
	fmt.Println(reflect.TypeOf(number))
	fmt.Println(reflect.TypeOf(number2))
	fmt.Println(reflect.TypeOf(interfaceAmaString))
	fmt.Println(reflect.TypeOf(arrStrings))
	fmt.Println(reflect.TypeOf(arrMixed))

	assert.Equal(t, "OK", "OK")
}
