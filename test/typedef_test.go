package test

import (
	"github.com/lwpyr/goscript"
	"github.com/lwpyr/goscript/common"
	"testing"
)

func TestTypeDef_A(t *testing.T) {
	setupClean()
	var expr string
	expr = `
typedef Person {
	name    string
	age     int32
	type    string
	hobbies []string
}

var Tom Person = {name('Tommy'), age(12), hobbies(['play_games', 'sports'])};
print(Tom);
var testMap StringMap = {1:'abc', 2:'def'};
print(testMap);

typedef StringMap [int32]string

typedef PersonArray []Person

var testArr PersonArray = [{name('aaa'),age(1)},{name('bbb'),age(2)}];
print(testArr[0]);
print(testArr[1]);

var testArr2 []Person = [{name('ccc'),age(3)},{name('ddd'),age(4)}];

testArr[1] = testArr2[0];
print(testArr);
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestTypeDef_B(t *testing.T) {
	setupClean()
	var expr string
	expr = `
typedef SelfRef {
	name string
	ref  SelfRef
}

var objA SelfRef = {name('objA')};
var objB SelfRef = {name('objB'), ref(objA)};
objA.ref = objB;

// you cannot directly call print(objA), this will cause stack overflow
print(objA.name);
print(objA.ref.name);
print(objB.name);
print(objB.ref.name);

`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestTypeDef_C(t *testing.T) {
	setupClean()
	var expr string
	expr = `
typedef Fruit {
	apple: 0
	banana: 1
	orange: 2
	watermelon: 3
	grape: 5
}

typedef Person {
	name    string
	like    Fruit
}

var Tom Person = {name('Tommy'), like(1)};
var Dave Person = {name('David'), like(Fruit.grape)};

print(Tom.name, ' likes ', enumString(Tom.like));
print(Dave.name, ' likes ', enumString(Dave.like));
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestTypeDef_D(t *testing.T) {
	setupClean()
	var expr string
	expr = `
typedef MessageType1 {
	str string
} 

typedef MessageType2 {
	str string
} 

typedef MessageType3 {
	str string
} 

typedef OneOfTest {
	oneof Type1AndType2 {
		msg1 MessageType1
		msg2 MessageType2
	}
	msg3 MessageType3
}

var obj OneOfTest = new OneOfTest();
obj.msg1 = {str('I\'m msg1')};
print(obj);
obj.msg3 = {str('I\'m msg3')};
print(obj);
obj.msg2 = {str('I\'m msg2')};
print(obj);
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func setupClean() {
	c = goscript.NewCompiler()
	mem = &common.Memory{
		Data: make([]interface{}, 1000),
	}

	scope := goscript.NewScope(nil)
	c.Scope = scope

	c.Include("common")
	c.Include("string")
	c.Include("json")
	c.Include("base64")
	c.Include("datetime")
}
