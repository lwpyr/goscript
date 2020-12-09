package test

import "testing"

func TestLambdaCalculus(t *testing.T) {
	setupParse()
	var expr string
	expr = `
func Identity(x object) object {
	return x;
}

func SelfApply(x object) object {
	local f func(object)object = x;
	return f(x);
}

func Apply(x object) func(object)object {
	local f func(object)object = x;
	return func(y object) object {
		return f(y);
	};
}

func TRUE(x object) func(object)object {
	local f func(object)object = x;
	return func(y object) object {
		return f;
	};
}

func FALSE(x object) func(object)object {
	return func(y object) object {
		return y;
	};
}

func CONDITION(x object) func(object)(func(object)object) {
	local f func(object)object = x;
	return func(y object) object {
		return func(c object) object {
			local BOOL func(object)(func(object)object) = c;
			return BOOL(x)(y);
		};
	};
}

func NOT(x object) object {
	local f func(object)(func(object)object) = x;
	return f(FALSE)(TRUE);
}

func AND(x object) func(object)object {
	local f func(object)(func(object)object) = x;
	return func(y object)object {
		return f(y)(FALSE);
	};
}

func OR(x object) func(object)object {
	local f func(object)(func(object)object) = x;
	return func(y object)object {
		return f(TRUE)(y);
	};
}

print(Identity == Identity(Identity)); // true
print(Identity == SelfApply(Identity)); // true

print(Apply(Identity)(Identity) == Identity); // true

print(TRUE(Identity)(Apply) == Identity); // true
print(FALSE(Identity)(Apply) == Apply); // true

print(CONDITION(Identity)(Apply)(TRUE) == Identity); // true
print(CONDITION(Identity)(Apply)(FALSE) == Identity); // false
print(CONDITION(Apply)(Identity)(TRUE) == Identity); // false
print(CONDITION(Apply)(Identity)(FALSE) == Identity); // true

print(NOT(TRUE) == FALSE); // true
print(NOT(NOT(TRUE)) == TRUE); // true

print(OR(TRUE)(TRUE) == TRUE); // true
print(OR(TRUE)(FALSE) == TRUE); // true
print(OR(FALSE)(TRUE) == TRUE); // true
print(OR(FALSE)(FALSE) == FALSE); // true

// Natural Number
func Next(x object) func(object)object {
	return func(y object) object {
		local f func(object)(func(object)object) = y;
		return f(FALSE)(x);
	};
}

func IsZero(x object) object {
	local f func(object)object = x;
	return f(TRUE);
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
	local n func(object)object = x;
	local z func(object)(func(object)object) = IsZero(n);
	return z(zero)(n(FALSE));
}

print(Prev(one) == zero); // true
print(Prev(nine) == eight); // true
print(Prev(nine) == seven); // false
`
	p := compilePro(expr)
	p.RunOnMemory(mem)
}
