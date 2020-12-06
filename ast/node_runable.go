package ast

import (
	"github.com/lw396285v/goscript/common"
	"github.com/lw396285v/goscript/lambda_chains"
)

type IConstantNode interface {
	ASTNode
	GetConstantKind() int
	GetConstantValue() interface{}
}

type VarNode struct {
	Node
	Variable *common.Variable
}

type SelectorNode struct {
	Node
	Data      ASTNode // selector or var
	FieldName string
}

type SliceFilterNode struct {
	Node
	Slice ASTNode
	Expr  ASTNode
}

type SliceMultiIndexNode struct {
	Node
	Slice   ASTNode
	Indices []ASTNode
}

type MapMultiIndexNode struct {
	Node
	Data   ASTNode
	Fields []ASTNode
}

type IndexNode struct {
	Node
	ToIndex   ASTNode
	Index     ASTNode
	IndexType string
}

type IndicesNode struct {
	Node
	From ASTNode
	To   ASTNode
	Step ASTNode
}

type BinaryNode struct {
	Node
	Lhs ASTNode
	Rhs ASTNode
	Op  string
}

type LeftUnaryNode struct {
	Node
	Operand ASTNode
	Op      string
}

type RightUnaryNode struct {
	Node
	Operand ASTNode
	Op      string
}

type AssignNode struct {
	Node
	Lhs []ASTNode
	Rhs []ASTNode
}

type ConstructorNode struct {
	Node
	Params []ASTNode
}

type FunctionAssignNode struct {
	Node
	Lhs      []ASTNode
	Function IFunctionNode
}

type InitializationSliceNode struct {
	Node
	Items []ASTNode
}

type InitializationMapNode struct {
	Node
	Keys   []ASTNode
	Values []ASTNode
}

type InitializationMessageNode struct {
	Node
	Keys   []string
	Values []ASTNode
}

type InitializationConstantNode struct {
	Node
	Constant ASTNode
}

type ValueNode struct {
	Node
	Val interface{}
}

func (i *InitializationSliceNode) Compile(c *Compiler) {
	num := len(i.Items)
	itemInstructions := make([]common.Instruction, 0, num)
	for _, item := range i.Items {
		item.Compile(c)
	}
	for j := 0; j < num; j++ {
		itemInstructions = append(itemInstructions, c.InstructionPop())
	}
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		items := make([]interface{}, 0, num)
		for j := 0; j < num; j++ {
			itemInstructions[j](m, stk)
			items = append(items, stk.Top())
			stk.Pop()
		}
		stk.Push(items)
	})
}

func (i *InitializationMapNode) Compile(c *Compiler) {
	num := len(i.Keys)
	keyInstructions := make([]common.Instruction, 0, num)
	valInstructions := make([]common.Instruction, 0, num)

	for _, key := range i.Keys {
		key.Compile(c)
	}
	for _, val := range i.Values {
		val.Compile(c)
	}
	for j := 0; j < num; j++ {
		valInstructions = append(valInstructions, c.InstructionPop())
	}
	for j := 0; j < num; j++ {
		keyInstructions = append(keyInstructions, c.InstructionPop())
	}

	constructor := common.MapConstructorMap[i.DataType.KeyType.Kind.Kind]
	mapSet := lambda_chains.GetMapSetFunc(i.DataType.KeyType)
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		ret := constructor()
		for j := 0; j < num; j++ {
			keyInstructions[j](m, stk)
			valInstructions[j](m, stk)
			mapSet(ret, stk.TopIndex(1), stk.Top())
			stk.PopN(2)
		}
		stk.Push(ret)
	})
}

func (i *InitializationMessageNode) Compile(c *Compiler) {
	num := len(i.Keys)
	valInstructions := make([]common.Instruction, 0, num)
	for _, val := range i.Values {
		val.Compile(c)
	}
	for j := 0; j < num; j++ {
		valInstructions = append(valInstructions, c.InstructionPop())
	}
	keys := i.Keys
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		ret := common.NewMessageValue()
		for j := 0; j < num; j++ {
			valInstructions[j](m, stk)
			lambda_chains.MessageSetField(ret, keys[j], stk.Top())
			stk.Pop()
			stk.Pc++
		}
		stk.Push(ret)
	})
}

func (i *InitializationConstantNode) Compile(c *Compiler) {
	i.Constant.Compile(c)
	constInstruction := c.InstructionPop()
	convertFunc := lambda_chains.GetConvertFunc(i.Constant.GetDataType(), i.GetDataType())
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		constInstruction(m, stk)
		stk.Set(0, convertFunc(stk.Top()))
	})
}

func (v *VarNode) Compile(c *Compiler) {
	if v.Variable != nil {
		variable := v.Variable
		if v.Lhs {
			if variable.IsParameter {
				c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
					stk.Push(stk.MustGet(variable))
				})
			} else {
				c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
					stk.Push(m.MustGet(variable))
				})
			}
		} else {
			if variable.IsParameter {
				c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
					stk.Push(stk.Get(variable))
				})
			} else {
				c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
					stk.Push(m.Get(variable))
				})
			}
		}
	} else {
		panic("Fatal error: variable information missing")
	}
}

func (s *SelectorNode) Compile(c *Compiler) {
	if s.Lhs {
		s.Data.SetLhs()
	}
	s.Data.Compile(c)
	dataInstruction := c.InstructionPop()
	fieldName := s.FieldName
	constructor := s.Data.GetDataType().Constructor
	var oneOfs []string
	if OneOfGroupName, ok := s.Data.GetDataType().OneOf[fieldName]; ok {
		tempOneOfs := s.Data.GetDataType().OneOfGroup[OneOfGroupName]
		for _, fn := range tempOneOfs {
			if fn != fieldName {
				oneOfs = append(oneOfs, fn)
			}
		}
	}
	if s.Lhs {
		if len(oneOfs) > 0 {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				dataInstruction(m, stk)
				ptr := stk.Top().(*interface{})
				if *ptr == nil {
					*ptr = constructor()
				}
				for _, name := range oneOfs {
					lambda_chains.MessageResetField(*ptr, name)
				}
				ptr = lambda_chains.MessageMustGetField(*ptr, fieldName)
				stk.Set(0, ptr)
			})
		} else {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				dataInstruction(m, stk)
				ptr := stk.Top().(*interface{})
				if *ptr == nil {
					*ptr = constructor()
				}
				ptr = lambda_chains.MessageMustGetField(*ptr, fieldName)
				stk.Set(0, ptr)
			})
		}

	} else {
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			dataInstruction(m, stk)
			if stk.Top() != nil {
				stk.Set(0, lambda_chains.MessageGetField(stk.Top(), fieldName))
			}
		})
	}
}

func (s *SliceFilterNode) Compile(c *Compiler) {
	if s.Lhs {
		panic("cannot use slice-filter an array on the left hand side")
	}
	s.Slice.Compile(c)
	s.Expr.Compile(c)
	exprInstruction := c.InstructionPop()
	filter := func(m *common.Memory, stk *common.Stack) {
		exprInstruction(m, stk)
		stk.Return()
	}
	sliceInstruction := c.InstructionPop()
	lengthFunc := lambda_chains.GetLengthFunc(s.DataType)
	boolConvertFunc := lambda_chains.GetConvertFunc(s.Expr.GetDataType(), common.BasicTypeMap["bool"])
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		sliceInstruction(m, stk)
		lenSlice := lengthFunc(stk.Top())
		val := make([]interface{}, 0, lenSlice)
		for i := int64(0); i < lenSlice; i++ {
			item := *lambda_chains.GetSliceField(stk.Top(), i)
			stk.Push(nil)
			stk.Push(item)
			stk.Call(filter, m, stk, 1)
			if boolConvertFunc(stk.Top()) == true {
				val = append(val, item)
			}
			stk.Pop()
		}
		stk.Set(0, val)
	})
}

func (s *SliceMultiIndexNode) Compile(c *Compiler) {
	if s.Slice.GetDataType().Kind.Kind == common.Slice {
		if s.Lhs {
			panic("cannot use slice-index an array on the left hand side")
		}
		s.Slice.Compile(c)
		for _, index := range s.Indices {
			index.Compile(c)
		}
		numIndex := len(s.Indices)
		indexInstructions := make([]common.Instruction, numIndex)
		for i := 0; i < numIndex; i++ {
			indexInstructions[i] = c.InstructionPop()
		}
		sliceInstruction := c.InstructionPop()
		lengthFunc := lambda_chains.GetLengthFunc(s.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			sliceInstruction(m, stk)
			slice := stk.Top()
			lenSlice := lengthFunc(slice)
			lenVal := int64(0)
			indices := make([][]int64, 0, len(indexInstructions))
			for _, indexInstruction := range indexInstructions {
				indexInstruction(m, stk)
				index := (stk.Top()).([]int64)
				if index[1] == -1 {
					index[1] = lenSlice
				}
				lenIndex := (index[1]-index[0]-1)/index[2] + 1
				if lenIndex < 0 {
					lenIndex = 0
				}
				lenVal += lenIndex
				indices = append(indices, index)
				stk.Pop()
			}
			val := make([]interface{}, 0, lenVal)
			for _, index := range indices {
				for i := index[0]; i < index[1]; i += index[2] {
					val = append(val, *lambda_chains.GetSliceField(slice, i))
				}
			}
			stk.Set(0, val)
		})
	} else {
		if s.Lhs {
			panic("cannot use slice-index an array on the left hand side")
		}
		s.Slice.Compile(c)
		s.Indices[0].Compile(c)
		indexInstruction := c.InstructionPop()
		sliceInstruction := c.InstructionPop()
		lengthFunc := lambda_chains.GetLengthFunc(s.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			sliceInstruction(m, stk)
			slice := stk.Top()
			indexInstruction(m, stk)
			index := (stk.Top()).([]int64)
			if index[1] == -1 {
				index[1] = lengthFunc(slice)
			}
			stk.Set(0, (slice.(string))[index[0]:index[1]])
		})
	}
}

func (i *IndexNode) Compile(c *Compiler) {
	if i.Lhs {
		i.ToIndex.SetLhs()
	}
	i.ToIndex.Compile(c)
	i.Index.Compile(c)
	switch i.IndexType {
	case "slice":
		indexInstruction := c.InstructionPop()
		sliceInstruction := c.InstructionPop()
		i64ConvertFunc := lambda_chains.GetConvertFunc(i.Index.GetDataType(), common.BasicTypeMap["int64"])
		if i.Lhs {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				sliceInstruction(m, stk)
				indexInstruction(m, stk)
				stk.Set(1, lambda_chains.GetSliceField(*stk.TopIndex(1).(*interface{}), i64ConvertFunc(stk.Top()).(int64)))
				stk.Pop()
			})
		} else {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				sliceInstruction(m, stk)
				if stk.Top() != nil {
					indexInstruction(m, stk)
					stk.Set(1, *lambda_chains.GetSliceField(stk.TopIndex(1), i64ConvertFunc(stk.Top()).(int64)))
					stk.Pop()
				}
			})
		}
	case "string":
		indexInstruction := c.InstructionPop()
		sliceInstruction := c.InstructionPop()
		i64ConvertFunc := lambda_chains.GetConvertFunc(i.Index.GetDataType(), common.BasicTypeMap["int64"])
		if i.Lhs {
			panic("string is immutable")
		} else {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				sliceInstruction(m, stk)
				if stk.Top() != nil {
					indexInstruction(m, stk)
					stk.Set(1, string((stk.TopIndex(1).(string))[i64ConvertFunc(stk.Top()).(int64)]))
					stk.Pop()
				}
			})
		}
	case "map":
		keyInstruction := c.InstructionPop()
		mapInstruction := c.InstructionPop()
		constructor := i.ToIndex.GetDataType().Constructor
		keyConvertFunc := lambda_chains.GetConvertFunc(i.Index.GetDataType(), i.ToIndex.GetDataType().KeyType)
		mapMustGet := lambda_chains.GetMapMustGetFunc(i.ToIndex.GetDataType().KeyType)
		mapGet := lambda_chains.GetMapGetFunc(i.ToIndex.GetDataType().KeyType)
		if i.Lhs {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				mapInstruction(m, stk)
				keyInstruction(m, stk)
				ptr := stk.TopIndex(1).(*interface{})
				if *ptr == nil {
					*ptr = constructor()
				}
				stk.Set(1, mapMustGet(*ptr, keyConvertFunc(stk.Top())))
				stk.Pop()
			})
		} else {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				mapInstruction(m, stk)
				if stk.Top() != nil {
					keyInstruction(m, stk)
					stk.Set(1, mapGet(stk.TopIndex(1), keyConvertFunc(stk.Top())))
					stk.Pop()
				}
			})
		}
	default:
		panic("wtf")
	}
}

func (i *IndicesNode) Compile(c *Compiler) {
	if i.Lhs {
		panic("cannot use slice-index an array on left hand side")
	}
	switch i.NodeType {
	case "IndexNodeType1":
		i.From.Compile(c)
		i.To.Compile(c)
		i.Step.Compile(c)
		stepInstruction := c.InstructionPop()
		toInstruction := c.InstructionPop()
		fromInstruction := c.InstructionPop()
		fromConvertFunc := lambda_chains.GetConvertFunc(i.From.GetDataType(), common.BasicTypeMap["int64"])
		toConvertFunc := lambda_chains.GetConvertFunc(i.To.GetDataType(), common.BasicTypeMap["int64"])
		stepConvertFunc := lambda_chains.GetConvertFunc(i.Step.GetDataType(), common.BasicTypeMap["int64"])
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			val := make([]int64, 3)
			fromInstruction(m, stk)
			val[0] = fromConvertFunc(stk.Top()).(int64)
			stk.Pop()
			toInstruction(m, stk)
			val[1] = toConvertFunc(stk.Top()).(int64)
			stk.Pop()
			stepInstruction(m, stk)
			val[2] = stepConvertFunc(stk.Top()).(int64)
			stk.Set(0, val)
		})
	case "IndexNodeType2":
		i.From.Compile(c)
		i.To.Compile(c)
		toInstruction := c.InstructionPop()
		fromInstruction := c.InstructionPop()
		fromConvertFunc := lambda_chains.GetConvertFunc(i.From.GetDataType(), common.BasicTypeMap["int64"])
		toConvertFunc := lambda_chains.GetConvertFunc(i.To.GetDataType(), common.BasicTypeMap["int64"])
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			val := make([]int64, 3)
			fromInstruction(m, stk)
			val[0] = fromConvertFunc(stk.Top()).(int64)
			stk.Pop()
			toInstruction(m, stk)
			val[1] = toConvertFunc(stk.Top()).(int64)
			val[2] = 1
			stk.Set(0, val)
		})
	case "IndexNodeType3":
		i.From.Compile(c)
		fromInstruction := c.InstructionPop()
		fromConvertFunc := lambda_chains.GetConvertFunc(i.From.GetDataType(), common.BasicTypeMap["int64"])
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			val := make([]int64, 3)
			fromInstruction(m, stk)
			val[0] = fromConvertFunc(stk.Top()).(int64)
			val[1] = -1
			val[2] = 1
			stk.Set(0, val)
		})
	case "IndexNodeType4":
		i.Step.Compile(c)
		stepInstruction := c.InstructionPop()
		stepConvertFunc := lambda_chains.GetConvertFunc(i.Step.GetDataType(), common.BasicTypeMap["int64"])
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			val := make([]int64, 3)
			stepInstruction(m, stk)
			val[0] = stepConvertFunc(stk.Top()).(int64)
			val[1] = val[0] + 1
			val[2] = 1
			stk.Set(0, val)
		})
	case "IndexNodeType5":
		i.To.Compile(c)
		toInstruction := c.InstructionPop()
		toConvertFunc := lambda_chains.GetConvertFunc(i.To.GetDataType(), common.BasicTypeMap["int64"])
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			val := make([]int64, 3)
			val[0] = 0
			toInstruction(m, stk)
			val[1] = toConvertFunc(stk.Top()).(int64)
			val[2] = 1
			stk.Set(0, val)
		})
	default:
		panic("wtf?")
	}
}

func (m *MapMultiIndexNode) Compile(c *Compiler) {
	if m.Lhs {
		panic("cannot use slice-index a map on the left hand side")
	}
	m.Data.Compile(c)
	for _, field := range m.Fields {
		field.Compile(c)
	}
	numFields := len(m.Fields)
	fieldInstructions := make([]common.Instruction, numFields)
	for i := 0; i < numFields; i++ {
		fieldInstructions[i] = c.InstructionPop()
	}
	mapInstruction := c.InstructionPop()
	//constructor := m.DataType.ItemType.Kind.Constructor
	keyConvertFuncs := make([]lambda_chains.TypeConvertFunc, numFields)
	for i := 0; i < numFields; i++ {
		keyConvertFuncs[i] = lambda_chains.GetConvertFunc(m.Fields[numFields-1-i].GetDataType(), m.Data.GetDataType().KeyType)
	}
	mapGet := lambda_chains.GetMapGetFunc(m.Data.GetDataType().KeyType)
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		mapInstruction(m, stk)
		if stk.Top() != nil {
			val := make([]interface{}, numFields)
			for i := 0; i < numFields; i++ {
				fieldInstructions[i](m, stk)
				if stk.Top() != nil {
					stk.Set(0, keyConvertFuncs[i](stk.Top()))
					val[i] = mapGet(stk.TopIndex(1), keyConvertFuncs[i](stk.Top()))
				} else {
					val[i] = nil
				}
				stk.Pop()
			}
			stk.Set(0, val)
		}
	})
}

func (n *BinaryNode) Compile(c *Compiler) {
	if n.Op == "=" {
		n.Lhs.SetLhs()
	}
	n.Lhs.Compile(c)
	n.Rhs.Compile(c)
	rhsInstruction := c.InstructionPop()
	lhsInstruction := c.InstructionPop()
	switch n.Op {
	case "+", "-", "*", "**", "/", "%", "//", "==", "!=", ">", "<", ">=", "<=":
		opFunc := lambda_chains.GetBinaryOpFunc(n.Op, n.Lhs.GetDataType(), n.Rhs.GetDataType())
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			lhsInstruction(m, stk)
			rhsInstruction(m, stk)
			opFunc(m, stk)
		})
	case "&&":
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			lhsInstruction(m, stk)
			if !(stk.Top()).(bool) {
				return
			}
			stk.Pop()
			rhsInstruction(m, stk)
		})
	case "||":
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			lhsInstruction(m, stk)
			if (stk.Top()).(bool) {
				return
			}
			stk.Pop()
			rhsInstruction(m, stk)
		})
	case "=":
		rhsConvertFunc := lambda_chains.GetConvertFunc(n.Rhs.GetDataType(), n.Lhs.GetDataType())
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			rhsInstruction(m, stk)
			if stk.Top() != nil {
				stk.Set(0, rhsConvertFunc(stk.Top()))
				lhsInstruction(m, stk)
				*stk.Top().(*interface{}) = stk.TopIndex(1)
				stk.Pop()
			}
		})
	case "~=":
		// todo
	default:
		panic("unknown binary op " + n.Op)
	}
}

func (u *LeftUnaryNode) Compile(c *Compiler) {
	opFunc := lambda_chains.GetLeftUnaryOpFunc(u.Op, u.Operand.GetDataType())
	if u.Op == "++" || u.Op == "--" {
		u.Operand.SetLhs()
	}
	u.Operand.Compile(c)
	operandInstruction := c.InstructionPop()
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		operandInstruction(m, stk)
		opFunc(m, stk)
	})
}

func (u *RightUnaryNode) Compile(c *Compiler) {
	opFunc := lambda_chains.GetRightUnaryOpFunc(u.Op, u.Operand.GetDataType())
	if u.Op == "++" || u.Op == "--" {
		u.Operand.SetLhs()
	}
	u.Operand.Compile(c)
	operandInstruction := c.InstructionPop()
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		operandInstruction(m, stk)
		opFunc(m, stk)
	})
}

func (f *FunctionAssignNode) Compile(c *Compiler) {
	num := len(f.Lhs)
	for _, lhs := range f.Lhs {
		lhs.SetLhs()
		lhs.Compile(c)
	}
	f.Function.Compile(c)
	functionInstruction := c.InstructionPop()
	lhsInstructions := make([]common.Instruction, 0, num)
	for i := 0; i < num; i++ {
		lhsInstructions = append(lhsInstructions, c.InstructionPop())
	}
	convertFuncs := make([]lambda_chains.TypeConvertFunc, 0, num)
	for i := 0; i < num; i++ {
		convertFuncs = append(convertFuncs, lambda_chains.GetConvertFunc(f.Lhs[num-1-i].GetDataType(), f.Function.GetReturnType()[i]))
	}
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		functionInstruction(m, stk)
		outs := (stk.Top()).([]interface{})
		for i := 0; i < num; i++ {
			if outs[i] != nil {
				lhsInstructions[i](m, stk)
				*stk.Top().(*interface{}) = outs[i]
				stk.Pop()
			}
		}
		stk.Set(0, nil)
	})
}

func (n *ConstructorNode) Compile(c *Compiler) {
	num := len(n.Params)
	constructor := n.DataType.Constructor
	if num == 0 {
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			stk.Push(constructor())
		})
	} else {
		for _, p := range n.Params {
			p.Compile(c)
		}
		paramInstructions := make([]common.Instruction, 0, num)
		for i := 0; i < num; i++ {
			paramInstructions = append(paramInstructions, c.InstructionPop())
		}
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			params := make([]interface{}, 0, num)
			for i := 0; i < num; i++ {
				paramInstructions[i](m, stk)
				params = append(params, stk.Top())
				stk.Pop()
			}
			stk.Push(constructor(params...))
		})
	}
}

func (a *AssignNode) Compile(c *Compiler) {
	num := len(a.Lhs)

	lhsInstructions := make([]common.Instruction, 0, num)
	rhsInstructions := make([]common.Instruction, 0, num)

	for _, lhs := range a.Lhs {
		lhs.SetLhs()
		lhs.Compile(c)
	}
	for _, rhs := range a.Rhs {
		rhs.Compile(c)
	}
	for i := 0; i < num; i++ {
		rhsInstructions = append(rhsInstructions, c.InstructionPop())
	}
	for i := 0; i < num; i++ {
		lhsInstructions = append(lhsInstructions, c.InstructionPop())
	}
	var assignInsts []common.Instruction
	for i := 0; i < num; i++ {
		lhsInstruction := lhsInstructions[i]
		rhsInstruction := rhsInstructions[i]
		convertFunc := lambda_chains.GetConvertFunc(a.Rhs[num-1-i].GetDataType(), a.Lhs[num-1-i].GetDataType())
		assignInsts = append(assignInsts, func(m *common.Memory, stk *common.Stack) {
			rhsInstruction(m, stk)
			if stk.Top() != nil {
				lhsInstruction(m, stk)
				*stk.Top().(*interface{}) = convertFunc(stk.TopIndex(1))
				stk.PopN(2)
			}
		})
	}
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		for i := 0; i < num; i++ {
			assignInsts[i](m, stk)
		}
		stk.Push(nil)
	})
}

func (v *ValueNode) Compile(c *Compiler) {
	c.InstructionPush(func(_ *common.Memory, stk *common.Stack) {
		stk.Push(v.Val)
	})
}

func (v *ValueNode) GetConstantKind() int {
	return v.DataType.Kind.Kind
}

func (v *ValueNode) GetConstantValue() interface{} {
	return v.Val
}
