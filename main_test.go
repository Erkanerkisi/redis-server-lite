package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"redis-lite/resp"
	"redis-lite/util"
	"testing"
)

func Test_deserializationOfString(t *testing.T) {
	command := "+OK\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)
	parsedString, _ := resp.parseSimple(splittedCommand)
	assert.Equal(t, "OK", parsedString)
}

func Test_DeserializationOfError(t *testing.T) {
	command := "-Error message\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)
	parsedString, _ := resp.parseSimple(splittedCommand)
	assert.Equal(t, "Error message", parsedString)
}

func Test_DeserializationOfInteger(t *testing.T) {
	command := ":1000\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)

	parsedString, _ := resp.parseNumber(splittedCommand)
	assert.Equal(t, "1000", parsedString)
}

func Test_DeserializationOfNilBulkString(t *testing.T) {
	command := "$-1\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)
	parsedString, _ := resp.parseBulkString(splittedCommand)
	assert.Empty(t, parsedString)
}

func Test_DeserializationOfBulkString(t *testing.T) {
	command := "$12\r\nAmit Shekhar\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)
	parsedString, _ := resp.parseBulkString(splittedCommand)
	assert.Equal(t, "Amit Shekhar", parsedString)
}

func Test_DeserializationOfEmptyArray(t *testing.T) {
	command := "*0\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)
	result, _ := resp.parseArray(splittedCommand)
	assert.Empty(t, result)
}

func Test_DeserializationOfArray(t *testing.T) {
	command := "*3\r\n$5\r\nhello\r\n:1\r\n$5\r\nworld\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)
	arr, _ := resp.parseArray(splittedCommand)
	assert.Equal(t, "hello", arr[0])
	assert.Equal(t, "1", arr[1])
	assert.Equal(t, "world", arr[2])

}

func Test_DeserializationOfArrayInteger(t *testing.T) {
	command := "*3\r\n:1\r\n:2\r\n:3\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)
	arr, _ := resp.parseArray(splittedCommand)
	assert.Equal(t, "1", arr[0])
	assert.Equal(t, "2", arr[1])
	assert.Equal(t, "3", arr[2])
}

func Test_DeserializationOfArrayMixedType(t *testing.T) {
	command := "*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$5\r\nhello\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)

	result, _ := resp.parseArray(splittedCommand)
	assert.Equal(t, "1", result[0])
	assert.Equal(t, "2", result[1])
	assert.Equal(t, "3", result[2])
	assert.Equal(t, "4", result[3])
	assert.Equal(t, "hello", result[4])

}

func Test_DeserializationOfNilArray(t *testing.T) {
	command := "*-1\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)

	result, _ := resp.parseArray(splittedCommand)
	assert.Empty(t, result)
}

func Test_Split(t *testing.T) {
	command := "*3\r\n:1\r\n:2\r\n:3\r\n"
	arr := bytes.Split(util.Byte(command), []byte{util.CarriageReturn, util.NewLine})
	fmt.Println(arr)
}

func Test_DeserializationOfArray_TestCase1(t *testing.T) {
	command := "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n"
	splittedCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)

	result, _ := resp.parseArray(splittedCommand)
	assert.Equal(t, "echo", result[0])
	assert.Equal(t, "hello world", result[1])
}

func Test_DeserializationOfNestedArray(t *testing.T) {
	command := "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n"
	splitCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)

	result, _ := resp.parseArray(splitCommand)
	//fmt.Println(result)
	first := result[0].([]interface{})
	second := result[1].([]interface{})
	assert.Equal(t, "1", first[0])
	assert.Equal(t, "Hello", second[0])
	assert.Equal(t, "World", second[1])
}

func Test_DeserializationOfNestedOfNestedArray(t *testing.T) {
	command := "*2\r\n*3\r\n*5\r\n:1\r\n:200\r\n:500\r\n+patates\r\n-haha\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n"
	splitCommand := bytes.Split(util.Byte(command), util.SeparatorBytes)

	result, _ := resp.parseArray(splitCommand)
	first := result[0].([]interface{})
	second := result[1].([]interface{})
	firstOfFirst := first[0].([]interface{})
	assert.Equal(t, "200", firstOfFirst[1])
	assert.Equal(t, "World", second[1])
}
