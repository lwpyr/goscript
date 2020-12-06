package program

import (
	"fmt"
	"github.com/lw396285v/goscript/common"
	"runtime/debug"
)

type SingleLineProgram struct {
	Root common.Instruction
}

func (p *SingleLineProgram) RunOnMemory(m *common.Memory) (ret interface{}) {
	stk := common.Get()
	defer common.Put(stk)
	p.Root(m, stk)
	return stk.Top()
}

func (p *SingleLineProgram) RunOnMemoryNoPanic(m *common.Memory) (ret interface{}, err error) {
	stk := common.Get()
	defer func() {
		common.Put(stk)
		r := recover()
		if r != nil {
			err = fmt.Errorf("%v %s", r, debug.Stack())
		}
	}()
	p.Root(m, stk)
	return stk.Top(), nil
}

func (p *SingleLineProgram) RunOnMemoryNoPanicNoStackTrace(m *common.Memory) (ret interface{}, err error) {
	stk := common.Get()
	defer func() {
		common.Put(stk)
		r := recover()
		if r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	p.Root(m, stk)
	return stk.Top(), nil
}

func (p *SingleLineProgram) Run(m *common.Memory, stk *common.Stack) interface{} {
	p.Root(m, stk)
	return stk.Top()
}
