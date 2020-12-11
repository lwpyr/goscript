package test

import (
	"github.com/lwpyr/goscript"
	"testing"
)

func TestLambdaCalculus(t *testing.T) {
	c = goscript.NewCompiler()
	c.Scope = goscript.NewScope(nil)
	c.Include("common")
	mem = goscript.NewMemory(100)
	var expr string
	expr = `
func Identity(x any) any {
	return x;
}

func SelfApply(x any) any {
	return x.(func(any)any)(x);
}

func Apply(x any) any {
	return func(y any) any {
		return x.(func(any)any)(y);
	};
}

func TRUE(x any) any {
	return func(y any) any {
		return x;
	};
}

func FALSE(x any) any {
	return func(y any) any {
		return y;
	};
}

func CONDITION(x any) any {
	return func(y any) any {
		return func(c any) any {
			return c.(func(any)any)(x).(func(any)any)(y);
		};
	};
}

func NOT(x any) any {
	return x.(func(any)any)(FALSE).(func(any)any)(TRUE);
}

func AND(x any) any {
	return func(y any)any {
		return x.(func(any)any)(y).(func(any)any)(FALSE);
	};
}

func OR(x any) any {
	return func(y any)any {
		return x.(func(any)any)(TRUE).(func(any)any)(y);
	};
}

print('=== test BASIC ===');
print(Identity == Identity(Identity)); // true
print(Identity == SelfApply(Identity)); // true

print('=== test APPLY ===');
print(Apply(Identity).(func(any)any)(Identity) == Identity); // true

print('=== test BOOL ===');
print(TRUE(Identity).(func(any)any)(Apply) == Identity); // true
print(FALSE(Identity).(func(any)any)(Apply) == Apply); // true

print('=== test COND ===');
print(CONDITION(Identity).(func(any)any)(Apply).(func(any)any)(TRUE) == Identity); // true
print(CONDITION(Identity).(func(any)any)(Apply).(func(any)any)(FALSE) == Apply); // true
print(CONDITION(Apply).(func(any)any)(Identity).(func(any)any)(TRUE) == Apply); // true
print(CONDITION(Apply).(func(any)any)(Identity).(func(any)any)(FALSE) == Identity); // true

print('=== test NOT ===');
print(NOT(TRUE) == FALSE); // true
print(NOT(NOT(TRUE)) == TRUE); // true

print('=== test AND ===');
print(AND(TRUE).(func(any)any)(TRUE) == TRUE); // true
print(AND(TRUE).(func(any)any)(FALSE) == FALSE); // true
print(AND(FALSE).(func(any)any)(TRUE) == FALSE); // true
print(AND(FALSE).(func(any)any)(FALSE) == FALSE); // true

print('=== test OR ===');
print(OR(TRUE).(func(any)any)(TRUE) == TRUE); // true
print(OR(TRUE).(func(any)any)(FALSE) == TRUE); // true
print(OR(FALSE).(func(any)any)(TRUE) == TRUE); // true
print(OR(FALSE).(func(any)any)(FALSE) == FALSE); // true

// Natural Number
func Next(x any) func(any)any {
	return func(y any) any {
		return y.(func(any)any)(FALSE).(func(any)any)(x);
	};
}

func IsZero(x any) any {
	return x.(func(any)any)(TRUE);
}

var zero  func(any)any = Identity;
var one   func(any)any = Next(zero);
var two   func(any)any = Next(one);
var three func(any)any = Next(two);
var four  func(any)any = Next(three);
var five  func(any)any = Next(four);
var six   func(any)any = Next(five);
var seven func(any)any = Next(six);
var eight func(any)any = Next(seven);
var nine  func(any)any = Next(eight);

print('=== test IsZero ===');
print(IsZero(zero) == TRUE); // true
print(IsZero(one) == FALSE); // true

func Prev(x any) any {
	return IsZero(x).(func(any)any)(zero).(func(any)any)(x.(func(any)any)(FALSE));
}

print('=== test Prev ===');
print(Prev(one) == zero); // true
print(Prev(nine) == eight); // true
print(Prev(nine) == seven); // false
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}
