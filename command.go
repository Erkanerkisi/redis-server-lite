package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CommandFactory(arr []string) Command {
	head := arr[0]
	switch strings.ToUpper(head) {
	case "PING":
		return Ping{}
	case "SET":
		return Set{}
	case "GET":
		return Get{}
	case "EXISTS":
		return Exists{}
	case "DEL":
		return Del{}
	case "INCR":
		return Incr{}
	case "DECR":
		return Decr{}
	case "LPUSH":
		return LPush{}
	case "LRANGE":
		return LRange{}
	case "CONFIG":
		return Ping{}
	default:
		panic("unrecognized command detected.")
	}
}

type Command interface {
	execute(arr []string) interface{}
}

type Ping struct {
}

func (ping Ping) execute(arr []string) interface{} {
	return Serialize("PONG")
}

type Set struct {
}

func (set Set) execute(arr []string) interface{} {
	GetStorage().set(arr[1], arr[2])
	return Serialize("OK")
}

type Get struct {
}

func (get Get) execute(arr []string) interface{} {
	value := GetStorage().get(arr[1])
	return Serialize(value)
}

type Exists struct {
}

func (exist Exists) execute(arr []string) interface{} {
	if GetStorage().exists(arr[1]) {
		return Serialize(1)
	}
	return Serialize(0)
}

type Del struct {
}

func (del Del) execute(arr []string) interface{} {
	var removedCount = 0
	for _, s := range arr[1:] {
		b := GetStorage().delete(s)
		if b {
			removedCount++
		}
	}
	return Serialize(removedCount)
}

type Incr struct {
}

func (incr Incr) execute(arr []string) interface{} {
	storage := GetStorage()
	key := arr[1]
	value := storage.get(key)
	if value == nil {
		newValue := 1
		storage.set(key, strconv.Itoa(newValue))
		return Serialize(newValue)
	}
	valueAsString := ParseStringFromInterface(value)
	intValue, err := strconv.Atoi(valueAsString)
	if err != nil {
		panic(fmt.Sprintf("key holds not an integer value. key : %s, value : %s ", arr[1], valueAsString))
	}
	newValue := intValue + 1
	storage.set(key, strconv.Itoa(newValue))
	return Serialize(newValue)

}

type Decr struct {
}

func (decr Decr) execute(arr []string) interface{} {
	storage := GetStorage()
	key := arr[1]
	value := storage.get(key)
	if value == nil {
		newValue := -1
		storage.set(key, strconv.Itoa(newValue))
		return Serialize(newValue)
	}
	valueAsString := ParseStringFromInterface(value)

	intValue, err := strconv.Atoi(valueAsString)
	if err != nil {
		panic(fmt.Sprintf("key holds not an integer value. key : %s, value : %s ", arr[1], valueAsString))
	}
	newValue := intValue - 1
	storage.set(key, strconv.Itoa(newValue))
	return Serialize(newValue)

}

type LPush struct {
}

func (lPush LPush) execute(arr []string) interface{} {
	storage := GetStorage()
	key := arr[1]
	val := arr[2]
	value := storage.get(key)
	if value == nil {
		newValue := []interface{}{val}
		storage.setArray(key, newValue)
		return Serialize(len(newValue))
	}
	if !IsArray(value) {
		panic(fmt.Sprintf("key does not hold an array. key: %s, value: %s", key, value))
	}
	array := storage.get(key).([]interface{})
	newArray := make([]interface{}, 0)
	newArray = append(newArray, val)
	for _, v := range array {
		newArray = append(newArray, v)
	}
	storage.setArray(key, newArray)
	return Serialize(len(newArray))
}

type LRange struct {
}

func (lRange LRange) execute(arr []string) interface{} {
	storage := GetStorage()
	key := arr[1]
	start, err := strconv.Atoi(arr[2])
	stop, err2 := strconv.Atoi(arr[3])
	if err != nil || err2 != nil {
		panic(fmt.Sprintf("start and stop values must be integer"))
	}
	value := storage.get(key)
	if value == nil {
		return Serialize(nil)
	}
	if !IsArray(value) {
		panic(fmt.Sprintf("key does not hold an array. key: %s, value: %s", key, value))
	}
	array := storage.get(key).([]interface{})

	if start > len(array)-1 || start > stop {
		return Serialize(make([]interface{}, 0))
	}
	if stop > len(array)-1 {
		stop = len(array) - 1
	}
	if start < 0 && stop < 0 {
		start = len(array) + start
		stop = len(array) + stop
	} else if start < len(array)*-1 {
		start = 0
	}
	return Serialize(array[start : stop+1])
}
