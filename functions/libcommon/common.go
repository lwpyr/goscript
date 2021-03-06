package libcommon

import (
	"fmt"
	"github.com/lwpyr/goscript/common"
	"math/rand"
	"strconv"
	"time"
)

func PlaceHolder(m *common.Memory, stk *common.Stack) {}

func ItoA(_ *common.Memory, stk *common.Stack) {
	stk.ReturnValue(strconv.FormatInt((stk.Top()).(int64), 10))
}

func AtoI(_ *common.Memory, stk *common.Stack) {
	ret, err := strconv.ParseInt((stk.Top()).(string), 10, 64)
	if err != nil {
		panic(err)
	}
	stk.ReturnValue(ret)
}

func FtoA(_ *common.Memory, stk *common.Stack) {
	stk.ReturnValue(fmt.Sprintf("%f", (stk.Top()).(float64)))
}

func AtoF(_ *common.Memory, stk *common.Stack) {
	ret, err := strconv.ParseFloat((stk.Top()).(string), 64)
	if err != nil {
		panic(err)
	}
	stk.ReturnValue(ret)
}

func Random(_ *common.Memory, stk *common.Stack) {
	stk.ReturnValue(rand.Float64())
}

func Print(_ *common.Memory, stk *common.Stack) {
	data := stk.Top().([]interface{})
	fmt.Println(data...)
}

func Sleep(_ *common.Memory, stk *common.Stack) {
	ms := stk.Top().(int64)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
