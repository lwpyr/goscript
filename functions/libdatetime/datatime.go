package libdatetime

import (
	"github.com/lwpyr/goscript/common"
	"strconv"
	"time"
)

func DateTimeAdd(m *common.Memory, stk *common.Stack) {
	durationSecond := stk.Top().(int64)
	t := time.Now()
	t = t.Add(time.Second * time.Duration(durationSecond))
	stk.ReturnValue(t.Format("Mon, 2 Jan 2006 15:04:05 MST"))
}

func Now(m *common.Memory, stk *common.Stack) {
	t := time.Now()
	stk.ReturnValue(t.UTC().UnixNano() / 1000000)
}

func NowStr(m *common.Memory, stk *common.Stack) {
	n := time.Now().UTC().UnixNano() / 1000000
	stk.ReturnValue(strconv.FormatInt(n, 10))
}
