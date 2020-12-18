package common

import (
	"reflect"
)

type Instruction func(m *Memory, stk *Stack)

func EvalConstInstructions(instructions []Instruction) interface{} {
	instructions = append(instructions, func(m *Memory, stk *Stack) {
		stk.Pc = -1
	}) // make sure it will end
	stk := Get()
	defer Put(stk)
	for stk.Pc != -1 {
		instructions[stk.Pc](nil, stk)
	}
	return stk.Top()
}

var ErrorInstruction Instruction = func(m *Memory, stk *Stack) {}

func IsError(i Instruction) bool {
	return reflect.ValueOf(i).Pointer() == reflect.ValueOf(ErrorInstruction).Pointer()
}
