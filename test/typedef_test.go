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
type Person message {
	name    string
	age     int32
	hobbies slice<string>
}

var Tom Person = {name('Tommy'), age(12), hobbies(['play_games', 'sports'])};
print(Tom);
var testMap StringMap = {1:'abc', 2:'def'};
print(testMap);

type StringMap map<int32, string>

type PersonArray slice<Person>

var testArr PersonArray = [{name('aaa'),age(1)},{name('bbb'),age(2)}];
print(testArr[0]);
print(testArr[1]);

var testArr2 slice<Person> = [{name('ccc'),age(3)},{name('ddd'),age(4)}];

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
type SelfRef message {
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
type Fruit enum {
	apple: 0
	banana: 1
	orange: 2
	watermelon: 3
	grape: 5
}

type Person message {
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