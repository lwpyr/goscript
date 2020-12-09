package common

type FunctionMeta struct {
	In        []*DataType
	Out       []*DataType
	TailArray bool
	ConstExpr bool // for future optimization
}

func (f *FunctionMeta) GenerateTypeName() string {
	ret := "func("
	for i, in := range f.In {
		if i != 0 {
			ret += ","
		}
		ret += in.Type
	}
	if f.TailArray {
		ret += "..."
	}
	ret += ")("
	for i, out := range f.Out {
		if i != 0 {
			ret += ","
		}
		ret += out.Type
	}
	ret += ")"
	return ret
}

type Function struct {
	Type *DataType
	F    *Instruction
}

type FunctionLib interface {
	Init(tr *TypeRegistry) map[string]*Function
}

var RegisteredLibs = map[string]FunctionLib{}

func RegisterLib(name string, lib FunctionLib) {
	RegisteredLibs[name] = lib
}
