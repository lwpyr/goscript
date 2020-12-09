package ast

type EnumNode struct {
	Node
	EnumName string
	Enum     map[string]int32
	REnum    map[int32]string
}

func (t *EnumNode) Compile(c *Compiler) {
	c.TypeRegistry.AddEnum(t.EnumName, t.Enum)
}

type TypeAlias struct {
	Node
	TypeName string
	RealType *TypeNameNode
}

func (t *TypeAlias) Compile(c *Compiler) {
	panic("cannot compile")
}

type TypeDef struct {
	Node
	TypeName  string
	FieldType map[string]*TypeNameNode
}

func (t *TypeDef) Compile(c *Compiler) {
	panic("implement me")
}
