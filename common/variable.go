package common

type VarType int

const (
	Global VarType = iota
	Local
	Captured
)

type Variable struct {
	Offset       int // for slice mem
	Symbol       string
	VariableType VarType
	DataType     *DataType
	Scope        *Scope
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

	EnableCapture bool
	CaptureOffset int
	Capture       []*Variable
	CapturedSet   map[string]*Variable
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
func (s *Scope) SetCaptureMode() {
	s.EnableCapture = true
	s.CapturedSet = map[string]*Variable{}
}

func (s *Scope) AddConstant(name string, constVal *Constant) {
	if s.Constants == nil {
		s.Constants = map[string]*Constant{}
	}
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
	v.VariableType = Global
	v.Scope = s
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
		v := s.Outer.GetVariable(name)
		if v != nil && s.EnableCapture && v.VariableType >= Local {
			if capture, ok := s.CapturedSet[v.Symbol]; ok {
				return capture
			} else {
				s.Capture = append(s.Capture, v)
				capture := &Variable{
					Offset:       s.CaptureOffset,
					Symbol:       v.Symbol,
					VariableType: Captured,
					Scope:        s,
					DataType:     v.DataType,
				}
				s.CaptureOffset++
				s.CapturedSet[capture.Symbol] = capture
				return capture
			}
		} else {
			return v
		}
	}
	return nil
}

func (s *Scope) AddParameterVariable(v *Variable) {
	s.Parameters[v.Symbol] = v
	v.Offset = s.LocalIndex
	v.VariableType = Local
	v.Scope = s
	s.LocalIndex++
}

func (s *Scope) AddReturnVariable(v *Variable) {
	s.Parameters[v.Symbol] = v
	v.Offset = s.ReturnIndex
	v.VariableType = Local
	v.Scope = s
	s.ReturnIndex--
}
