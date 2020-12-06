package common

type FunctionMeta struct {
	Name      string
	In        []*DataType
	Out       []*DataType
	F         Instruction
	TailArray bool
	ConstExpr bool // for future optimization
}

type Func func(v ...*interface{}) []*interface{}

type FunctionLib interface {
	Init(tr *TypeRegistry) map[string]*FunctionMeta
}

var RegisteredLibs = map[string]FunctionLib{}

func RegisterLib(name string, lib FunctionLib) {
	RegisteredLibs[name] = lib
}
