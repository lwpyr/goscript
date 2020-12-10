package test

import (
	"github.com/lwpyr/goscript"
	"github.com/lwpyr/goscript/common"
	"testing"
)

func TestLambdaCalculus(t *testing.T) {
	c = goscript.NewCompiler()
	c.Scope = goscript.NewScope(nil)
	c.Include("common")
	mem = &common.Memory{
		Data: make([]interface{}, 1000),
	}
	var expr string
	expr = `
func Identity(x object) object {
	return x;
}

func SelfApply(x object) object {
	return x.(func(object)object)(x);
}

func Apply(x object) object {
	return func(y object) object {
		return x.(func(object)object)(y);
	};
}

func TRUE(x object) object {
	return func(y object) object {
		return x;
	};
}

func FALSE(x object) object {
	return func(y object) object {
		return y;
	};
}

func CONDITION(x object) object {
	return func(y object) object {
		return func(c object) object {
			return c.(func(object)object)(x).(func(object)object)(y);
		};
	};
}

func NOT(x object) object {
	return x.(func(object)object)(FALSE).(func(object)object)(TRUE);
}

func AND(x object) object {
	return func(y object)object {
		return x.(func(object)object)(y).(func(object)object)(FALSE);
	};
}

func OR(x object) object {
	return func(y object)object {
		return x.(func(object)object)(TRUE).(func(object)object)(y);
	};
}

print(Identity == Identity(Identity)); // true
print(Identity == SelfApply(Identity)); // true

print(Apply(Identity).(func(object)object)(Identity) == Identity); // true

print(TRUE(Identity).(func(object)object)(Apply) == Identity); // true
print(FALSE(Identity).(func(object)object)(Apply) == Apply); // true

print(CONDITION(Identity).(func(object)object)(Apply).(func(object)object)(TRUE) == Identity); // true
print(CONDITION(Identity).(func(object)object)(Apply).(func(object)object)(FALSE) == Identity); // false
print(CONDITION(Apply).(func(object)object)(Identity).(func(object)object)(TRUE) == Identity); // false
print(CONDITION(Apply).(func(object)object)(Identity).(func(object)object)(FALSE) == Identity); // true

print(NOT(TRUE) == FALSE); // true
print(NOT(NOT(TRUE)) == TRUE); // true

print(OR(TRUE).(func(object)object)(TRUE) == TRUE); // true
print(OR(TRUE).(func(object)object)(FALSE) == TRUE); // true
print(OR(FALSE).(func(object)object)(TRUE) == TRUE); // true
print(OR(FALSE).(func(object)object)(FALSE) == FALSE); // true

// Natural Number
func Next(x object) func(object)object {
	return func(y object) object {
		return y.(func(object)object)(FALSE).(func(object)object)(x);
	};
}

func IsZero(x object) object {
	return x.(func(object)object)(TRUE);
}

var zero  func(object)object = Identity;
var one   func(object)object = Next(zero);
var two   func(object)object = Next(one);
var three func(object)object = Next(two);
var four  func(object)object = Next(three);
var five  func(object)object = Next(four);
var six   func(object)object = Next(five);
var seven func(object)object = Next(six);
var eight func(object)object = Next(seven);
var nine  func(object)object = Next(eight);

print(IsZero(zero) == TRUE); // true
print(IsZero(one) == FALSE); // true

func Prev(x object) object {
	return IsZero(x).(func(object)object)(zero).(func(object)object)(x.(func(object)object)(FALSE));
}

print(Prev(one) == zero); // true
print(Prev(nine) == eight); // true
print(Prev(nine) == seven); // false
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}
