package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/instruction"
)

type SymbolNode struct {
	Node
	SymbolName    string
	SymbolType    common.SymbolType
	DataTypeNode  ASTNode
	ValueToAssign ASTNode
}

func (t *SymbolNode) Compile(c *CompileContext) {
	if t.DataTypeNode != nil {
		t.DataTypeNode.Compile(c)
		t.DataType = t.DataTypeNode.GetDataType()
	}
	switch t.SymbolType {
	case common.Global:
		if t.ValueToAssign != nil {
			if t.DataType != nil {
				t.ValueToAssign.SetRequiredType(t.DataType)
				t.ValueToAssign.Compile(c)
			} else {
				t.ValueToAssign.Compile(c)
				t.DataType = t.ValueToAssign.GetDataType()
			}

			t.AppendInstruction(t.ValueToAssign.GetInstructions()...)
			sym := &common.Symbol{
				Symbol:   t.SymbolName,
				DataType: t.DataType,
			}
			c.Scope.AddGlobalVariable(sym)
			t.AppendInstruction(instruction.GlobalSymbolAssign(sym))
		} else {
			c.Scope.AddGlobalVariable(&common.Symbol{
				Symbol:   t.SymbolName,
				DataType: t.DataType,
			})
		}
	case common.Local:
		if t.ValueToAssign != nil {
			if t.DataType != nil {
				t.ValueToAssign.SetRequiredType(t.DataType)
				t.ValueToAssign.Compile(c)
			} else {
				t.ValueToAssign.Compile(c)
				t.DataType = t.ValueToAssign.GetDataType()
			}
			t.AppendInstruction(t.ValueToAssign.GetInstructions()...)
		} else {
			t.AppendInstruction(instruction.LocalSymbolNil())
		}
		c.Scope.AddLocalVariable(&common.Symbol{
			Symbol:   t.SymbolName,
			DataType: t.DataType,
		})
	case common.Const:
		if t.DataType != nil {
			t.ValueToAssign.SetRequiredType(t.DataType)
			t.ValueToAssign.Compile(c)
		} else {
			t.ValueToAssign.Compile(c)
			t.DataType = t.ValueToAssign.GetDataType()
		}
		defer func() {
			if r := recover(); r != nil {
				panic(common.NewCompileErr(t.ErrorWithSource("const symbol should be initialized using const expression")))
			}
		}()
		constValue := common.EvalConstInstructions(t.ValueToAssign.GetInstructions())

		c.Scope.AddConstant(t.SymbolName, &common.Symbol{
			Symbol:   t.SymbolName,
			DataType: t.ValueToAssign.GetDataType(),
			Data:     constValue,
		})
	}
}
