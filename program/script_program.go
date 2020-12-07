package program

import (
	"fmt"
	"github.com/lwpyr/goscript/common"
)

type ScriptProgram struct {
	Root common.Instruction
}

func (p *ScriptProgram) RunOnMemory(m *common.Memory) {
	stk := common.Get()
	defer common.Put(stk)
	p.Root(m, stk)
}

func (p *ScriptProgram) RunOnMemoryNoPanic(m *common.Memory) (err error) {
	stk := common.Get()
	defer func() {
		common.Put(stk)
		r := recover()
		if r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	p.Root(m, stk)
	return nil
}

func (p *ScriptProgram) Run(m *common.Memory, stk *common.Stack) {
	p.Root(m, stk)
}
