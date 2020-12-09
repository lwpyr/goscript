package libjson

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/lwpyr/goscript/common"
	"strconv"
	"unicode"
)

func JsonMarshal(m *common.Memory, stk *common.Stack) {
	bytes, err := jsoniter.Marshal(stk.Top())
	if err != nil {
		panic(err)
	}
	stk.ReturnValue(bytes)
}

func JsonUnmarshal(m *common.Memory, stk *common.Stack) {
	bytes := (stk.TopIndex(1)).([]byte)
	T := (stk.Top()).(*common.DataType)
	stk.Pop()

	iter := jsoniter.ConfigFastest.BorrowIterator(bytes)
	defer jsoniter.ConfigFastest.ReturnIterator(iter)
	stk.ReturnValue(T.Unmarshal(iter))
}

func JsonObject(m *common.Memory, stk *common.Stack) {
	bytes := (stk.Top()).([]byte)
	val := map[string]interface{}{}
	err := jsoniter.Unmarshal(bytes, &val)
	if err != nil {
		panic(err)
	}
	stk.ReturnValue(val)
}

func JsonPath(m *common.Memory, stk *common.Stack) {
	jsonObject := stk.TopIndex(1)
	path := (stk.Top()).(string)

	p := NewParser(path)
	for tokeType, token := p.next(); tokeType != eof; tokeType, token = p.next() {
		if tokeType == selector {
			jsonObject = jsonObject.(map[string]interface{})[token.(string)]
		} else {
			jsonObject = jsonObject.([]interface{})[token.(int64)]
		}
	}
	stk.ReturnValue(jsonObject)
}

type simpleParser struct {
	idx  int
	path string
}

const (
	selector = iota
	index
	eof
)

func NewParser(path string) *simpleParser {
	return &simpleParser{
		idx:  0,
		path: path,
	}
}

func (s *simpleParser) next() (int, interface{}) {
	if s.idx >= len(s.path) {
		return eof, nil
	}
	for s.idx < len(s.path) && s.path[s.idx] == ' ' {
		s.idx++
	}
	if s.idx == 0 {
		if s.path[0] != '$' {
			panic("invalid path")
		}
		s.idx++
	}
	if s.path[s.idx] == '.' {
		s.idx++
		if s.idx == len(s.path) {
			panic("invalid path")
		}
		begin := s.idx
		for s.idx < len(s.path) && s.path[s.idx] != '.' && s.path[s.idx] != '[' {
			if s.path[s.idx] == ' ' {
				panic("invalid path")
			}
			s.idx++
		}
		return selector, s.path[begin:s.idx]
	} else if s.path[s.idx] == '[' {
		s.idx++
		for s.idx < len(s.path) && s.path[s.idx] == ' ' {
			s.idx++
		}
		if s.idx == len(s.path) {
			panic("invalid path")
		}
		begin := s.idx
		for s.idx < len(s.path) && unicode.IsDigit(rune(s.path[s.idx])) {
			s.idx++
		}
		end := s.idx
		if begin == end {
			panic("invalid path")
		}
		for s.idx < len(s.path) && s.path[s.idx] == ' ' {
			s.idx++
		}
		if s.idx == len(s.path) || s.path[s.idx] != ']' {
			panic("invalid path")
		}
		s.idx++
		num, _ := strconv.ParseInt(s.path[begin:end], 10, 64)
		return index, num
	} else {
		panic("invalid path")
	}
}
