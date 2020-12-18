package common

type Memory struct {
	Data []interface{}
}

func (m *Memory) Copy() *Memory {
	ret := &Memory{Data: make([]interface{}, len(m.Data))}
	copy(ret.Data, m.Data)
	return ret
}

//go:norace
//go:nosplit
func (m *Memory) Get(v *Symbol) interface{} {
	return m.Data[v.Offset]
}

//go:norace
//go:nosplit
func (m *Memory) Set(v *Symbol, data interface{}) {
	m.Data[v.Offset] = data
}

//go:norace
//go:nosplit
func (m *Memory) MustGet(v *Symbol) *interface{} {
	return &m.Data[v.Offset]
}
