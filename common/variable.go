package common

type Variable struct {
	Offset      int    // for slice mem
	Symbol      string // for map mem
	IsParameter bool
	Type        *DataType
}

type Constant struct {
	Symbol string
	Type   *DataType
	Data   interface{}
}

type Scope struct {
	Outer       *Scope
	VarIndex    *int
	LocalIndex  int
	ReturnIndex int
	Variables   map[string]*Variable
	Parameters  map[string]*Variable
	Constants   map[string]*Constant
}

func NewScope(outer *Scope) *Scope {
	var varIndex *int
	var localIndex int
	var returnIndex int
	if outer != nil {
		varIndex = outer.VarIndex
		localIndex = outer.LocalIndex
		returnIndex = outer.ReturnIndex
	} else {
		index := 0
		varIndex = &index
		localIndex = 0
		returnIndex = -1
	}
	return &Scope{
		Outer:       outer,
		VarIndex:    varIndex,
		LocalIndex:  localIndex,
		ReturnIndex: returnIndex,
		Variables:   map[string]*Variable{},
		Constants:   map[string]*Constant{},
		Parameters:  map[string]*Variable{},
	}
}

func (s *Scope) AddConstant(name string, constVal *Constant) {
	s.Constants[name] = constVal
}

func (s *Scope) GetConstant(name string) *Constant {
	if ret, ok := s.Constants[name]; ok {
		return ret
	}
	if s.Outer != nil {
		return s.Outer.GetConstant(name)
	}
	return nil
}

func (s *Scope) AddVariable(v *Variable) {
	s.Variables[v.Symbol] = v
	v.Offset = *s.VarIndex
	v.IsParameter = false
	*s.VarIndex = *s.VarIndex + 1
}

func (s *Scope) GetVariable(name string) *Variable {
	if ret, ok := s.Parameters[name]; ok {
		return ret
	}
	if ret, ok := s.Variables[name]; ok {
		return ret
	}
	if s.Outer != nil {
		return s.Outer.GetVariable(name)
	}
	return nil
}

func (s *Scope) AddParameterVariable(v *Variable) {
	s.Parameters[v.Symbol] = v
	v.Offset = s.LocalIndex
	v.IsParameter = true
	s.LocalIndex++
}

func (s *Scope) AddReturnVariable(v *Variable) {
	s.Parameters[v.Symbol] = v
	v.Offset = s.ReturnIndex
	v.IsParameter = true
	s.ReturnIndex--
}
