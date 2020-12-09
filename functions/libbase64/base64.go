package libbase64

import (
	"encoding/base64"
	"github.com/lwpyr/goscript/common"
)

func EncodeBase64(m *common.Memory, stk *common.Stack) {
	stk.ReturnValue(base64.StdEncoding.EncodeToString(stk.Top().([]byte)))
}

func DecodeBase64(m *common.Memory, stk *common.Stack) {
	val := (stk.Top()).(string)
	decoded, err := base64.StdEncoding.DecodeString(val)

	if err != nil {
		panic(err)
	}
	stk.ReturnValue(decoded)
}
