package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/instruction"
	"reflect"
)

var BreakPlaceHolder = func(m *common.Memory, stk *common.Stack) {}
var ContinuePlaceHolder = func(m *common.Memory, stk *common.Stack) {}

type BlockNode struct {
	Node
	Executions []ASTNode
}

type ReturnNode struct {
	Node
	Expr []ASTNode
}

type ProgramRoot struct {
	Node
	TypeDefNode   []ASTNode
	StatementNode []ASTNode
}

type BreakNode struct {
	Node
}

type ContinueNode struct {
	Node
}

type IfNode struct {
	Node
	Condition ASTNode
	Block     ASTNode
	ElseBlock ASTNode
}

type ForMapNode struct {
	Node
	KeyName string
	ValName string
	Map     ASTNode
	Serial  ASTNode
}

type ForSliceNode struct {
	Node
	ItemName string
	Slice    ASTNode
	Serial   ASTNode
}

type SliceIterator struct {
	slice []interface{}
	idx   int
}

type ForNode struct {
	Node
	Init      ASTNode
	Condition ASTNode
	Step      ASTNode
	Block     ASTNode
}

type SwitchNode struct {
	Node
	Expr       ASTNode
	Conditions []interface{}
	Blocks     []ASTNode
}

type RestoreStackNode struct {
	Node
	Line ASTNode
}

func (b *RestoreStackNode) Compile(c *CompileContext) {
	b.Line.Compile(c)
	b.AppendInstruction(b.Line.GetInstructions()...)
	n := b.Line.GetStackIncrement()
	if n > 0 {
		b.AppendInstruction(instruction.StackPopN(n))
	}
}

func (b *BlockNode) Compile(c *CompileContext) {
	c.MakeChildScope()
	defer c.ReturnParentScope()

	for _, r := range b.Executions {
		r.Compile(c)
	}
	for i := 0; i < len(b.Executions); i++ {
		b.AppendInstruction(b.Executions[i].GetInstructions()...)
	}
	num := len(c.Scope.Parameters)
	if num > 0 {
		b.AppendInstruction(instruction.StackPopN(num))
	}
}

func (r *ReturnNode) Compile(c *CompileContext) {
	// todo check validity
	lenRet := len(r.Expr)
	for _, expr := range r.Expr {
		expr.Compile(c)
		r.AppendInstruction(expr.GetInstructions()...)
	}
	if lenRet == 1 {
		r.AppendInstruction(instruction.StackReturn())
	} else if lenRet > 1 {
		r.AppendInstruction(instruction.StackReturnN(lenRet))
	}
	r.AppendInstruction(instruction.StackReturnVoid())
}

func (b *BreakNode) Compile(_ *CompileContext) {
	b.Instructions = []common.Instruction{BreakPlaceHolder}
}

func (c *ContinueNode) Compile(_ *CompileContext) {
	c.Instructions = []common.Instruction{ContinuePlaceHolder}
}

func (n *IfNode) Compile(c *CompileContext) {
	n.Condition.Compile(c)
	n.AppendInstruction(n.Condition.GetInstructions()...)
	n.Block.Compile(c)
	blockInstructions := n.Block.GetInstructions()
	if n.ElseBlock == nil {
		skip := len(blockInstructions) + 1
		n.AppendInstruction(
			func(m *common.Memory, stk *common.Stack) {
				if stk.Top().(bool) {
					stk.Pop()
					stk.Pc++
				} else {
					stk.Pop()
					stk.Pc += skip
				}
			},
		)
		n.AppendInstruction(blockInstructions...)
	} else {
		n.ElseBlock.Compile(c)
		elseBlockInstructions := n.ElseBlock.GetInstructions()
		skip1 := len(blockInstructions) + 2
		skip2 := len(elseBlockInstructions) + 1
		n.AppendInstruction(
			func(m *common.Memory, stk *common.Stack) {
				if stk.Top().(bool) {
					stk.Pop()
					stk.Pc++
				} else {
					stk.Pop()
					stk.Pc += skip1
				}
			},
		)
		n.AppendInstruction(blockInstructions...)
		n.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
			stk.Pc += skip2
		})
		n.AppendInstruction(elseBlockInstructions...)
	}
}

func relocateBreakAndContinue(instructions []common.Instruction) {
	lenSerialInstructions := len(instructions)
	for i := 0; i < lenSerialInstructions; i++ {
		if reflect.ValueOf(instructions[i]).Pointer() == reflect.ValueOf(BreakPlaceHolder).Pointer() {
			instructions[i] = func(m *common.Memory, stk *common.Stack) {
				stk.Pc += lenSerialInstructions - i + 1
			}
		} else if reflect.ValueOf(instructions[i]).Pointer() == reflect.ValueOf(ContinuePlaceHolder).Pointer() {
			instructions[i] = func(m *common.Memory, stk *common.Stack) {
				stk.Pc += lenSerialInstructions - i
			}
		}
	}
}

func (f *ForMapNode) Compile(c *CompileContext) {
	f.Map.Compile(c)
	if f.Map.GetDataType().Kind.Kind != common.Map {
		panic(common.NewTypeErr(f.ErrorWithSource("for loop with slice doesn't get a slice")))
	}
	f.AppendInstruction(f.Map.GetInstructions()...)

	c.MakeChildScope()
	defer c.ReturnParentScope()

	c.Scope.AddLocalVariable(&common.Symbol{
		Symbol:   "$",
		DataType: common.BasicTypeMap[common.AnyType],
	})
	c.Scope.AddLocalVariable(&common.Symbol{
		Symbol:   f.KeyName,
		DataType: f.Map.GetDataType().KeyType,
	})
	c.Scope.AddLocalVariable(&common.Symbol{
		Symbol:   f.ValName,
		DataType: f.Map.GetDataType().ValueType,
	})

	f.Serial.Compile(c)
	serialInstructions := f.Serial.GetInstructions()
	relocateBreakAndContinue(serialInstructions)
	lenSerialInstructions := len(serialInstructions)
	f.AppendInstruction(
		// initialization
		func(m *common.Memory, stk *common.Stack) {
			iter := reflect.ValueOf(stk.Top()).MapRange()
			stk.Set(0, iter)
			stk.Pc++
		},
	)
	f.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
		// condition
		iter := stk.Top().(*reflect.MapIter)
		if iter.Next() {
			stk.Push(iter.Key().Interface())
			stk.Push(iter.Value().Interface())
			stk.Pc++
		} else {
			stk.Pop()
			stk.Pc += lenSerialInstructions + 2
		}
	})
	f.AppendInstruction(serialInstructions...)
	f.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
		// step
		stk.PopN(2)
		stk.Pc -= lenSerialInstructions + 1
	})
}

func InitSliceIterator(s []interface{}) *SliceIterator {
	return &SliceIterator{
		slice: s,
		idx:   -1,
	}
}

func (i *SliceIterator) Next() bool {
	i.idx++
	if i.idx < len(i.slice) {
		return true
	}
	return false
}

func (i *SliceIterator) Item() interface{} {
	return i.slice[i.idx]
}

func (f *ForSliceNode) Compile(c *CompileContext) {
	f.Slice.Compile(c)
	if f.Slice.GetDataType().Kind.Kind != common.Slice {
		panic(common.NewTypeErr(f.ErrorWithSource("for loop with slice doesn't get a slice")))
	}
	f.AppendInstruction(f.Slice.GetInstructions()...)

	c.MakeChildScope()
	defer c.ReturnParentScope()

	c.Scope.AddLocalVariable(&common.Symbol{
		Symbol:   "$",
		DataType: common.BasicTypeMap[common.Int64Type],
	})
	c.Scope.AddLocalVariable(&common.Symbol{
		Symbol:   f.ItemName,
		DataType: f.Slice.GetDataType().ItemType,
	})

	f.Serial.Compile(c)
	serialInstructions := f.Serial.GetInstructions()
	relocateBreakAndContinue(serialInstructions)
	lenSerialInstructions := len(serialInstructions)
	f.AppendInstruction(
		// initialization
		func(m *common.Memory, stk *common.Stack) {
			iter := InitSliceIterator(stk.Top().([]interface{}))
			stk.Set(0, iter)
			stk.Pc++
		},
	)
	f.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
		// condition
		iter := stk.Top().(*SliceIterator)
		if iter.Next() {
			stk.Push(iter.Item())
			stk.Pc++
		} else {
			stk.Pop()
			stk.Pc += lenSerialInstructions + 2
		}
	})
	f.AppendInstruction(serialInstructions...)
	f.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
		// step
		stk.Pop()
		stk.Pc -= lenSerialInstructions + 1
	})

	c.Scope = c.Scope.Outer
}

func (f *ForNode) Compile(c *CompileContext) {
	c.MakeChildScope()
	defer c.ReturnParentScope()

	f.Init.Compile(c)
	f.AppendInstruction(f.Init.GetInstructions()...)
	f.Condition.Compile(c)
	conditionInstructions := f.Condition.GetInstructions()
	lenConditionInstructions := len(conditionInstructions)
	f.AppendInstruction(conditionInstructions...)
	f.Step.Compile(c)
	stepInstructions := f.Step.GetInstructions()
	lenStepInstructions := len(stepInstructions)
	f.Block.Compile(c)
	serialInstructions := f.Block.GetInstructions()
	lenSerialInstructions := len(serialInstructions)
	num := len(c.Scope.Parameters)
	skip := lenSerialInstructions + lenStepInstructions + 2
	f.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
		// condition
		if stk.Top().(bool) {
			stk.Pop()
			stk.Pc++
		} else {
			stk.PopN(num + 1)
			stk.Pc += skip
		}
	})
	f.AppendInstruction(serialInstructions...)
	f.AppendInstruction(stepInstructions...)
	back := lenSerialInstructions + lenStepInstructions + lenConditionInstructions + 1
	f.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
		// step
		stk.Pc -= back
	})
}

func (s *SwitchNode) Compile(c *CompileContext) {
	s.Expr.Compile(c)
	s.AppendInstruction(s.Expr.GetInstructions()...)
	num := len(s.Conditions)
	var blockInstructions []common.Instruction
	var skip []int
	last := 1
	for i := 0; i < num; i++ {
		s.Blocks[i].Compile(c)
		blockInstructions = append(blockInstructions, s.Blocks[i].GetInstructions()...)
		skip = append(skip, last)
		last += len(s.Blocks[i].GetInstructions())
	}
	conditions := s.Conditions
	s.AppendInstruction(func(m *common.Memory, stk *common.Stack) {
		val := stk.Top()
		stk.Pop()
		for i := 0; i < num; i++ {
			if val == conditions[i] {
				stk.Pc += skip[i]
				return
			}
		}
		stk.Pc += last
	})
	s.AppendInstruction(blockInstructions...)
}
