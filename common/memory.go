package common

type Memory struct {
	Data []interface{}
}

func (m *Memory) Copy() *Memory {
	ret := &Memory{Data: make([]interface{}, len(m.Data))}
	copy(ret.Data, m.Data)
	return ret
}

func (m *Memory) Get(v *Variable) interface{} {
	return m.Data[v.Offset]
}

func (m *Memory) MustGet(v *Variable) *interface{} {
	return &m.Data[v.Offset]
}
