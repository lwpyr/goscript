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
type lambda func(any)any // func alias

func Identity(x any) any {
	return x;
}

func SelfApply(x any) any {
	return x.(lambda)(x);
}

func Apply(x any) any {
	return func(y any) any {
		return x.(lambda)(y);
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
			return c.(lambda)(x).(lambda)(y);
		};
	};
}

func NOT(x any) any {
	return x.(lambda)(FALSE).(lambda)(TRUE);
}

func AND(x any) any {
	return func(y any)any {
		return x.(lambda)(y).(lambda)(FALSE);
	};
}

func OR(x any) any {
	return func(y any)any {
		return x.(lambda)(TRUE).(lambda)(y);
	};
}

print('=== test BASIC ===');
print(Identity == Identity(Identity)); // true
print(Identity == SelfApply(Identity)); // true

print('=== test APPLY ===');
print(Apply(Identity).(lambda)(Identity) == Identity); // true

print('=== test BOOL ===');
print(TRUE(Identity).(lambda)(Apply) == Identity); // true
print(FALSE(Identity).(lambda)(Apply) == Apply); // true

print('=== test COND ===');
print(CONDITION(Identity).(lambda)(Apply).(lambda)(TRUE) == Identity); // true
print(CONDITION(Identity).(lambda)(Apply).(lambda)(FALSE) == Apply); // true
print(CONDITION(Apply).(lambda)(Identity).(lambda)(TRUE) == Apply); // true
print(CONDITION(Apply).(lambda)(Identity).(lambda)(FALSE) == Identity); // true

print('=== test NOT ===');
print(NOT(TRUE) == FALSE); // true
print(NOT(NOT(TRUE)) == TRUE); // true

print('=== test AND ===');
print(AND(TRUE). (lambda)(TRUE)  == TRUE); // true
print(AND(TRUE). (lambda)(FALSE) == FALSE); // true
print(AND(FALSE).(lambda)(TRUE)  == FALSE); // true
print(AND(FALSE).(lambda)(FALSE) == FALSE); // true

print('=== test OR ===');
print(OR(TRUE). (lambda)(TRUE)  == TRUE); // true
print(OR(TRUE). (lambda)(FALSE) == TRUE); // true
print(OR(FALSE).(lambda)(TRUE)  == TRUE); // true
print(OR(FALSE).(lambda)(FALSE) == FALSE); // true

// Natural Number
func Next(x any) lambda {
	return func(y any) any {
		return y.(lambda)(FALSE).(lambda)(x);
	};
}

func IsZero(x any) any {
	return x.(lambda)(TRUE);
}

var zero  lambda = Identity;
var one   lambda = Next(zero);
var two   lambda = Next(one);
var three lambda = Next(two);
var four  lambda = Next(three);
var five  lambda = Next(four);
var six   lambda = Next(five);
var seven lambda = Next(six);
var eight lambda = Next(seven);
var nine  lambda = Next(eight);

print('=== test IsZero ===');
print(IsZero(zero) == TRUE); // true
print(IsZero(one) == FALSE); // true

func Prev(x any) any {
	return IsZero(x).(lambda)(zero).(lambda)(x.(lambda)(FALSE));
}

print('=== test Prev ===');
print(Prev(one) == zero); // true
print(Prev(nine) == eight); // true
print(Prev(nine) != seven); // true
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}
