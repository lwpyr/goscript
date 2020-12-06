package test

import (
	"fmt"
	"github.com/lw396285v/goscript"
	"github.com/lw396285v/goscript/common"
	"github.com/lw396285v/goscript/program"
	"testing"
	"time"
)

func compilePro(expr string) *program.ScriptProgram {
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
	p := compilePro(expr)
	start := time.Now()
	p.RunOnMemory(mem)
	fmt.Println(time.Since(start))
}

func TestB(t *testing.T) {
	setupParse()
	var expr string
	expr = `
num2 = 0;
for (num1=1; num1<=100; num1 = num1 + 1) {
	num2 = num2 + num1;
}
print(num2);
`
	p := compilePro(expr)
	p.RunOnMemory(mem)
}

func TestC(t *testing.T) {
	setupParse()
	var expr string
	expr = `
var sum int64;
sum = 0;
for (local i int64 = 1; i <= 100; i++) {
	sum = sum + i;
}
print(sum);
`
	p := compilePro(expr)
	p.RunOnMemory(mem)
}

func BenchmarkC(b *testing.B) {
	setupParse()
	var expr string
	expr = `
var sum int64;
sum = 0;
for (local i int64 = 1; i <= 100; i++) {
	sum = sum + i;
}
`
	p := compilePro(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemory(mem)
	}
}

func setupParse() {
	// mock data
	c = go_script.NewCompiler()

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
		Data: make([]interface{}, 100),
	}

	c.Include("common")
	c.Include("string")
	c.Include("json")
	c.Include("base64")
	c.Include("datetime")

	scope := go_script.NewScope(nil)

	scope.AddVariable(go_script.NewVariable("Tom", c.FindType("Person")))
	scope.AddVariable(go_script.NewVariable("Jerry", c.FindType("Person")))
	scope.AddVariable(go_script.NewVariable("Friends", c.FindSliceType("Person")))
	scope.AddVariable(go_script.NewVariable("Class", c.FindMapType("int64", "Person")))
	scope.AddVariable(go_script.NewVariable("newPerson", c.FindType("Person")))
	scope.AddVariable(go_script.NewVariable("DavidId", c.FindType("int64")))
	scope.AddVariable(go_script.NewVariable("Teacher", c.FindType("Person")))
	scope.AddVariable(go_script.NewVariable("RandNumber", c.FindType("float64")))
	scope.AddVariable(go_script.NewVariable("num1", c.FindType("int64")))
	scope.AddVariable(go_script.NewVariable("num2", c.FindType("int64")))
	scope.AddVariable(go_script.NewVariable("num3", c.FindType("int64")))
	scope.AddVariable(go_script.NewVariable("num4", c.FindType("int64")))
	scope.AddVariable(go_script.NewVariable("testString", c.FindType("string")))
	scope.AddVariable(go_script.NewVariable("NewClass", c.FindMapType("int64", "Person")))
	scope.AddVariable(go_script.NewVariable("jsonObj", c.FindType("JSONObject")))
	scope.AddVariable(go_script.NewVariable("testString2", c.FindType("string")))
	scope.AddVariable(go_script.NewVariable("header", c.FindMapType("string", "string")))
	scope.AddVariable(go_script.NewVariable("fruitEnum", c.FindType("fruits")))
	scope.AddVariable(go_script.NewVariable("stringMap", c.FindMapType("string", "Person")))
	scope.AddVariable(go_script.NewVariable("float32Map", c.FindMapType("float32", "Person")))

	c.Scope = scope
}
