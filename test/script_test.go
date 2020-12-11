package test

import (
	"fmt"
	"github.com/lwpyr/goscript"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/program"
	"testing"
	"time"
)

func compileScript(expr string) *program.ScriptProgram {
	ret, err := c.CompileScript(expr)
	if err != nil {
		panic(err)
	}
	return &program.ScriptProgram{Root: ret}
}

func TestA(t *testing.T) {
	setupParse()
	var expr string
	expr = `
func fib(v int64) int64 {
	if(v<2) {
		return v;
	}
	return fib(v-1)+fib(v-2);
}

print(fib(35));`
	p := compileScript(expr)
	start := time.Now()
	p.RunOnMemory(mem)
	fmt.Println(time.Since(start))
}

func TestB(t *testing.T) {
	setupParse()
	var expr string
	expr = `
var num2 int64 = 0;
var num1 int64 = 0;
for (num1=1; num1<=100; num1 = num1 + 1) {
	num2 = num2 + num1;
}
print(num2);
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestC(t *testing.T) {
	setupParse()
	var expr string
	expr = `
var sum int64 = 0;
for (local i int64 = 1; i <= 100; i++) {
	sum = sum + i;
}
print(sum);
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestD(t *testing.T) {
	setupParse()
	var expr string
	expr = `
var Tom Person = new Person();
func who(p Person) {
	if(p.name == 'Tom') {
		print('Tommy');
	} else if(p.name == 'Dave') {
		print('David');
	} else {
		print('I don\'t know');
	}
}
Tom = {name('Tom')};
who(Tom);
Tom = {name('Dave')};
who(Tom);
Tom = {name('Lisa')};
who(Tom);
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestE(t *testing.T) {
	setupParse()
	var expr string
	expr = `
var f func(int64,int64)(int64) = func(a int64, b int64) int64 {
	return a+b;
}; 
print(f(1,2));
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestF(t *testing.T) {
	setupParse()
	var expr string
	expr = `
var fs []func()int64 = new []func()int64(2);
func main(a int64, b int64) func()(int64) {
	return func()int64{
		return a+b;
	};
}

fs[0] = main(1,2);
fs[1] = main(3,4);
print(fs[0]());
print(fs[1]());
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func TestG(t *testing.T) {
	setupParse()
	var expr string
	expr = `
const NAME string = 'Wonderful Goscript!';

var str1 string = NAME + ' yes!';
var str2 string = NAME + ' ok!';

print(NAME);
print(str1);
print(str2);

var str3 string = NAME;
var str4 string = NAME;

str3 += ' yes!';
str4 += ' ok!';

print(NAME);
print(str3);
print(str4);
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}

func BenchmarkA(b *testing.B) {
	setupParse()
	var expr string
	expr = `
var sum int64;
sum = 0;
for (local i int64 = 1; i <= 100; i++) {
	sum = sum + i;
}
`
	p := compileScript(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemory(mem)
	}
}

func setupParse() {
	// mock data
	c = goscript.NewCompiler()

	dtb := common.NewDataTypeBuilder("Person")
	dtb.SetKind(common.Message)
	dtb.SetField("name", c.FindType("string"))
	dtb.SetField("age", c.FindType("int32"))
	dtb.SetField("hobbies", c.FindSliceType("string"))

	c.AddBuiltType(dtb)

	c.AddEnum("fruits", map[string]int32{
		"apple":      int32(0),
		"banana":     int32(1),
		"orange":     int32(2),
		"strawberry": int32(3),
	})

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
