package common

type Stack struct {
	Bp   int
	Sp   int
	Pc   int
	Data []interface{}
}

//go:norace
//go:nosplit
func (s *Stack) Top() interface{} { // equals to TopIndex(0)
	return s.Data[s.Sp]
}

//go:norace
//go:nosplit
func (s *Stack) TopIndex(i int) interface{} {
	return s.Data[s.Sp-i]
}

//go:norace
//go:nosplit
func (s *Stack) Pop() { // equals to PopN(1)
	s.Data = s.Data[:s.Sp]
	s.Sp--
}

//go:norace
//go:nosplit
func (s *Stack) Set(spOffset int, i interface{}) {
	s.Data[s.Sp-spOffset] = i
}

//go:norace
//go:nosplit
func (s *Stack) PopN(i int) {
	s.Data = s.Data[:s.Sp-i+1]
	s.Sp -= i
}

//go:norace
//go:nosplit
func (s *Stack) ReturnValue(val interface{}) {
	s.Data[s.Bp-1] = val
	s.Pc = -1
}

//go:norace
func (s *Stack) ReturnValues(values ...interface{}) {
	for i, val := range values {
		s.Data[s.Bp-i-1] = val
	}
	s.Pc = -1
}

//go:norace
//go:nosplit
func (s *Stack) Push(i interface{}) {
	s.Data = append(s.Data, i)
	s.Sp++
}

//go:norace
//go:nosplit
func (s *Stack) Get(v *Symbol) interface{} {
	return s.Data[s.Bp+v.Offset]
}

//go:norace
//go:nosplit
func (s *Stack) MustGet(v *Symbol) *interface{} {
	return &s.Data[s.Bp+v.Offset]
}

//go:norace
//go:nosplit
func (s *Stack) Call(inst Instruction, mem *Memory, stk *Stack, params int) {
	bp := s.Bp
	pc := s.Pc
	s.Bp = s.Sp - params + 1
	s.Pc = 0
	inst(mem, stk)
	s.Sp = s.Bp - 1
	s.Bp = bp
	s.Data = s.Data[:s.Sp+1]
	s.Pc = pc
}

func (s *Stack) Reset() *Stack {
	s.Pc = 0
	s.Sp = -1
	s.Bp = 0
	s.Data = s.Data[:0]
	return s
}
