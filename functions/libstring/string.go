package libstring

import (
	"github.com/lw396285v/goscript/common"
	"strings"
)

func StartWith(m *common.Memory, stk *common.Stack) {
	str := stk.TopIndex(1).(string)
	prefix := stk.TopIndex(0).(string)
	if len(str) < len(prefix) {
		stk.Pop()
		stk.Set(0, false)
		return
	}
	stk.ReturnVal(str[:len(prefix)] == prefix)
}

func EndWith(m *common.Memory, stk *common.Stack) {
	str := stk.TopIndex(1).(string)
	suffix := stk.Top().(string)
	if len(str) < len(suffix) {
		stk.ReturnVal(false)
		return
	}
	stk.ReturnVal(str[len(str)-len(suffix):] == suffix)
}

func Contain(m *common.Memory, stk *common.Stack) {
	str := stk.TopIndex(1).(string)
	substr := stk.Top().(string)
	stk.ReturnVal(strings.Contains(str, substr))
}

func StrCat(m *common.Memory, stk *common.Stack) {
	delimiter := stk.TopIndex(1).(string)
	val := (stk.Top()).([]interface{})
	str := make([]string, 0, len(val))
	for _, param := range val {
		str = append(str, param.(string))
	}
	stk.ReturnVal(strings.Join(str, delimiter))
}

func StrSplit(m *common.Memory, stk *common.Stack) {
	str := stk.TopIndex(1).(string)
	delimiter := (stk.Top()).(string)

	items := strings.Split(str, delimiter)
	ret := make([]interface{}, 0, len(items))
	for _, item := range items {
		ret = append(ret, item)
	}
	stk.ReturnVal(ret)
}
