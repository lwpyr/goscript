package common

type SymbolType int

const (
	Global SymbolType = iota
	Const
	Local
	Captured
)

type Symbol struct {
	Offset     int // for slice mem
	Symbol     string
	SymbolType SymbolType
	DataType   *DataType
	Scope      *Scope
	Data       interface{}
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
	Variables   map[string]*Symbol
	Parameters  map[string]*Symbol
	Constants   map[string]*Symbol

	EnableCapture bool
	CaptureOffset int
	Capture       []*Symbol
	CapturedSet   map[string]*Symbol
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
		Variables:   map[string]*Symbol{},
		Constants:   map[string]*Symbol{},
		Parameters:  map[string]*Symbol{},
	}
}
func (s *Scope) SetCaptureMode() {
	s.EnableCapture = true
	s.CapturedSet = map[string]*Symbol{}
}

func (s *Scope) AddLocalVariable(v *Symbol) {
	s.Parameters[v.Symbol] = v
	v.Offset = s.LocalIndex
	v.SymbolType = Local
	v.Scope = s
	s.LocalIndex++
}

func (s *Scope) AddReturnVariable(v *Symbol) {
	s.Parameters[v.Symbol] = v
	v.Offset = s.ReturnIndex
	v.SymbolType = Local
	v.Scope = s
	s.ReturnIndex--
}

func (s *Scope) AddConstant(name string, constVal *Symbol) {
	s.Constants[name] = constVal
	constVal.Scope = s
	constVal.SymbolType = Const
}

func (s *Scope) AddGlobalVariable(v *Symbol) {
	s.Variables[v.Symbol] = v
	v.Offset = *s.VarIndex
	v.SymbolType = Global
	v.Scope = s
	*s.VarIndex = *s.VarIndex + 1
}

func (s *Scope) GetSymbol(name string) *Symbol {
	if ret, ok := s.Constants[name]; ok {
		return ret
	}
	if ret, ok := s.Parameters[name]; ok {
		return ret
	}
	if ret, ok := s.Variables[name]; ok {
		return ret
	}
	if s.Outer != nil {
		v := s.Outer.GetSymbol(name)
		if v != nil && s.EnableCapture && (v.SymbolType == Local || v.SymbolType == Captured) {
			if capture, ok := s.CapturedSet[v.Symbol]; ok {
				return capture
			} else {
				s.Capture = append(s.Capture, v)
				capture := &Symbol{
					Offset:     s.CaptureOffset,
					Symbol:     v.Symbol,
					SymbolType: Captured,
					Scope:      s,
					DataType:   v.DataType,
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

func IsBuiltIn(name string) bool {
	switch name {
	case "pushBack", "pushFront", "delete", "len", "enumString", "typeof":
		return true
	}
	return false
}
