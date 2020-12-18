package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/instruction"
)

type ASTNode interface {
	Compile(c *Compiler)
	SetLhs()
	SetRequiredType(dt *common.DataType)

	GetDataType() *common.DataType
	SetDataType(*common.DataType)

	IsVariadic() bool

	GetInstructions() []common.Instruction

	AppendInstruction(instruction ...common.Instruction)

	GetConstantValue() interface{}

	CheckIsConstant()
	PostProcess()

	SetSource(ctx antlr.ParserRuleContext)
	GetText() string
	ErrorWithSource(string) string

	SetStackPtr()
	IsStackPtr() bool
	GetStackIncrement() int
}

type Node struct {
	LineNumber     int
	Text           string
	Instructions   []common.Instruction
	DataType       *common.DataType
	RequiredType   *common.DataType
	StackPtr       bool
	LhsFlag        bool
	Variadic       bool
	StackIncrement int
	ConstantValue  interface{}
}

func (n *Node) GetStackIncrement() int {
	return n.StackIncrement
}

func (n *Node) SetStackPtr() {
	n.StackPtr = true
}

func (n *Node) IsStackPtr() bool {
	return n.StackPtr
}

func (n *Node) CheckIsConstant() {
	panic("implement me")
}

func (n *Node) PostProcess() {
	if n.RequiredType != nil {
		covertInstruction := instruction.TypeConvert(n.DataType, n.RequiredType)
		if common.IsError(covertInstruction) {
			panic(common.NewTypeErr(n.ErrorWithSource("type convert cannot be implemented")))
		}
		if covertInstruction != nil {
			n.AppendInstruction(covertInstruction)
		}
		n.DataType = n.RequiredType
	}
	if !n.Variadic {
		val := common.EvalConstInstructions(n.Instructions)
		n.ConstantValue = val
		n.Instructions = []common.Instruction{instruction.PushConstantToStack(val)}
	}
}

func (n *Node) GetText() string {
	return n.Text
}

func (n *Node) GetConstantValue() interface{} {
	return n.ConstantValue
}

func (n *Node) AppendInstruction(instruction ...common.Instruction) {
	n.Instructions = append(n.Instructions, instruction...)
}

func (n *Node) Compile(_ *Compiler) {
	panic("this won't compile")
}

func (n *Node) SetSource(ctx antlr.ParserRuleContext) {
	n.LineNumber = ctx.GetStart().GetLine()
	n.Text = ctx.GetText()
}

func (n *Node) ErrorWithSource(message string) string {
	return fmt.Sprintf("%s, at line %d, %s", message, n.LineNumber, n.Text)
}

func (n *Node) GetInstructions() []common.Instruction {
	return n.Instructions
}

func (n *Node) IsVariadic() bool {
	return n.Variadic
}

func (n *Node) GetDataType() *common.DataType {
	return n.DataType
}

func (n *Node) SetRequiredType(dt *common.DataType) {
	n.RequiredType = dt
}

func (n *Node) SetDataType(dt *common.DataType) {
	n.DataType = dt
}

func (n *Node) SetLhs() {
	n.LhsFlag = true
}
