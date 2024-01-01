package main

import (
	"bytes"
	"fmt"
)

type Operation struct {
}

func (operation *Operation) Handle(buffer []byte) string {
	command := ClearAllZeroBytes(buffer)
	fmt.Println(string(command))
	checkByteHasTerminatorAtTheEndOfTheArray(command)
	splitCommand := bytes.Split(command, SeparatorBytes)
	val, _ := parseResp(splitCommand)
	//not sure of this
	strArr := convertInterfaceToStringArr(val)
	result := CommandFactory(strArr).execute(strArr)
	return result.(string)
}

func checkByteHasTerminatorAtTheEndOfTheArray(command []byte) {
	if command[len(command)-2] != CarriageReturn || command[len(command)-1] != NewLine {
		panic("wrong terminator operator for basic string")
	}
}
