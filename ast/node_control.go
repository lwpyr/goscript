package ast

import (
	"github.com/lwpyr/goscript/common"
	"reflect"
)

var BreakPlaceHolder = func(m *common.Memory, stk *common.Stack) {}
var ContinuePlaceHolder = func(m *common.Memory, stk *common.Stack) {}

type LineNode struct {
	Node
	Line ASTNode
}

type RestoreStackNode struct {
	Node
	Line ASTNode
}

type BlockNode struct {
	Node
	Executions []ASTNode
	Scope      *common.Scope
}

type ReturnNode struct {
	Node
	Expr []ASTNode
}

type ProgramRoot struct {
	Node
	EnumNode        []ASTNode
	TypeDefNode     []ASTNode
	VarDefNode      []ASTNode
	FunctionDefNode []ASTNode
	RunnableNode    []ASTNode
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
	Key     *common.Variable
	Val     *common.Variable
	Serial  ASTNode
}

type ForSliceNode struct {
	Node
	ItemName string
	Slice    ASTNode
	Item     *common.Variable
	Serial   ASTNode
}

type SliceIterator struct {
	slice []interface{}
	idx   int
}

type ForNode struct {
	Node
	Scope     *common.Scope
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

func (b *LineNode) Compile(c *Compiler) {
	b.Line.Compile(c)
	b.Instructions = b.Line.GetInstructions()
	b.Instructions = append(b.Instructions,
		func(m *common.Memory, stk *common.Stack) {
			stk.Pc++
		},
	)
}

func (b *RestoreStackNode) Compile(c *Compiler) {
	b.Line.Compile(c)
	lineInst := c.InstructionPop()
	b.Instructions = []common.Instruction{
		func(m *common.Memory, stk *common.Stack) {
			sp := stk.Sp
			lineInst(m, stk)
			stk.Data = stk.Data[:sp+1]
			stk.Sp = sp
			stk.Pc++
		},
	}
}

func (b *BlockNode) Compile(c *Compiler) {
	for _, r := range b.Executions {
		r.Compile(c)
	}
	for i := 0; i < len(b.Executions); i++ {
		b.Instructions = append(b.Instructions, b.Executions[i].GetInstructions()...)
	}
	num := len(b.Scope.Parameters)
	b.Instructions = append(b.Instructions, func(m *common.Memory, stk *common.Stack) {
		stk.PopN(num)
		stk.Pc++
	})
}

// todo: process type
func (p *ProgramRoot) Compile(c *Compiler) {
	num := len(p.RunnableNode)
	var runnableInstructions []common.Instruction
	for i := 0; i < num; i++ {
		p.RunnableNode[i].Compile(c)
		runnableInstructions = append(runnableInstructions, p.RunnableNode[i].GetInstructions()...)
	}
	runnableInstructions = append(runnableInstructions, func(m *common.Memory, stk *common.Stack) {
		stk.Pc = -1
	})
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		for stk.Pc != -1 {
			runnableInstructions[stk.Pc](m, stk)
		}
	})
}

func (r *ReturnNode) Compile(c *Compiler) {
	lenRet := len(r.Expr)
	for _, expr := range r.Expr {
		expr.Compile(c)
	}
	var exprInstructions []common.Instruction
	for i := 0; i < lenRet; i++ {
		exprInstructions = append(exprInstructions, c.InstructionPop())
	}
	if r.Expr == nil {
		r.Instructions = []common.Instruction{
			func(m *common.Memory, stk *common.Stack) {
				stk.Pc = -1
			},
		}
	} else {
		r.Instructions = []common.Instruction{
			func(m *common.Memory, stk *common.Stack) {
				for _, exprInstruction := range exprInstructions {
					exprInstruction(m, stk)
				}
				stk.ReturnN(lenRet)
			},
		}
	}
}

func (b *BreakNode) Compile(_ *Compiler) {
	b.Instructions = []common.Instruction{BreakPlaceHolder}
}

func (c *ContinueNode) Compile(_ *Compiler) {
	c.Instructions = []common.Instruction{ContinuePlaceHolder}
}

func (n *IfNode) Compile(c *Compiler) {
	n.Condition.Compile(c)
	conditionInstruction := c.InstructionPop()
	n.Block.Compile(c)
	blockInstructions := n.Block.GetInstructions()
	if n.ElseBlock == nil {
		skip := len(blockInstructions) + 1
		n.Instructions = []common.Instruction{
			func(m *common.Memory, stk *common.Stack) {
				conditionInstruction(m, stk)
				if stk.Top().(bool) {
					stk.Pop()
					stk.Pc++
				} else {
					stk.Pop()
					stk.Pc += skip
				}
			},
		}
		n.Instructions = append(n.Instructions, blockInstructions...)
	} else {
		n.ElseBlock.Compile(c)
		elseBlockInstructions := n.ElseBlock.GetInstructions()
		skip1 := len(blockInstructions) + 2
		skip2 := len(elseBlockInstructions) + 1
		n.Instructions = []common.Instruction{
			func(m *common.Memory, stk *common.Stack) {
				conditionInstruction(m, stk)
				if stk.Top().(bool) {
					stk.Pop()
					stk.Pc++
				} else {
					stk.Pop()
					stk.Pc += skip1
				}
			},
		}
		n.Instructions = append(n.Instructions, blockInstructions...)
		n.Instructions = append(n.Instructions, func(m *common.Memory, stk *common.Stack) {
			stk.Pc += skip2
		})
		n.Instructions = append(n.Instructions, elseBlockInstructions...)
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

func (f *ForMapNode) Compile(c *Compiler) {
	f.Map.Compile(c)
	mapInstruction := c.InstructionPop()
	f.Serial.Compile(c)
	serialInstructions := f.Serial.GetInstructions()
	relocateBreakAndContinue(serialInstructions)
	lenSerialInstructions := len(serialInstructions)
	f.Instructions = []common.Instruction{
		// initialization
		func(m *common.Memory, stk *common.Stack) {
			mapInstruction(m, stk)
			iter := reflect.ValueOf(stk.Top()).MapRange()
			stk.Set(0, iter)
			stk.Pc++
		},
	}
	f.Instructions = append(f.Instructions, func(m *common.Memory, stk *common.Stack) {
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
	f.Instructions = append(f.Instructions, serialInstructions...)
	f.Instructions = append(f.Instructions, func(m *common.Memory, stk *common.Stack) {
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

func (f *ForSliceNode) Compile(c *Compiler) {
	f.Slice.Compile(c)
	arrInstruction := c.InstructionPop()
	f.Serial.Compile(c)
	serialInstructions := f.Serial.GetInstructions()
	relocateBreakAndContinue(serialInstructions)
	lenSerialInstructions := len(serialInstructions)
	f.Instructions = []common.Instruction{
		// initialization
		func(m *common.Memory, stk *common.Stack) {
			arrInstruction(m, stk)
			iter := InitSliceIterator(stk.Top().([]interface{}))
			stk.Set(0, iter)
			stk.Pc++
		},
	}
	f.Instructions = append(f.Instructions, func(m *common.Memory, stk *common.Stack) {
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
	f.Instructions = append(f.Instructions, serialInstructions...)
	f.Instructions = append(f.Instructions, func(m *common.Memory, stk *common.Stack) {
		// step
		stk.Pop()
		stk.Pc -= lenSerialInstructions + 1
	})
}

func (f *ForNode) Compile(c *Compiler) {
	f.Init.Compile(c)
	initInstruction := f.Init.GetInstructions()
	f.Condition.Compile(c)
	conditionInstruction := c.InstructionPop()
	f.Step.Compile(c)
	stepInstruction := f.Step.GetInstructions()[0]
	f.Block.Compile(c)
	serialInstructions := f.Block.GetInstructions()
	lenSerialInstructions := len(serialInstructions)
	num := len(f.Scope.Parameters)
	f.Instructions = initInstruction
	f.Instructions = append(f.Instructions, func(m *common.Memory, stk *common.Stack) {
		// condition
		conditionInstruction(m, stk)
		if stk.Top().(bool) {
			stk.Pop()
			stk.Pc++
		} else {
			stk.PopN(num + 1)
			stk.Pc += lenSerialInstructions + 2
		}
	})
	f.Instructions = append(f.Instructions, serialInstructions...)
	f.Instructions = append(f.Instructions, func(m *common.Memory, stk *common.Stack) {
		// step
		stepInstruction(m, stk)
		stk.Pc -= lenSerialInstructions + 2
	})
}

func (s *SwitchNode) Compile(c *Compiler) {
	s.Expr.Compile(c)
	exprInst := c.InstructionPop()
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
	s.Instructions = append(s.Instructions, func(m *common.Memory, stk *common.Stack) {
		exprInst(m, stk)
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
	s.Instructions = append(s.Instructions, blockInstructions...)
}
