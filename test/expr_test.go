package test

import (
	"fmt"
	"github.com/lwpyr/goscript"
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/program"
	"github.com/stretchr/testify/assert"
	"testing"
)

func compile(expr string) *program.SingleLineProgram {
	ret, err := c.CompileExpression(expr)
	if err != nil {
		panic(err)
	}
	return &program.SingleLineProgram{Root: ret}
}

func TestExpression(t *testing.T) {
	setup()
	var expr string
	var p *program.SingleLineProgram

	expr = "Tom.name"
	p = compile(expr)
	assert.Equal(t, "Tom", p.RunOnMemory(mem))

	expr = "Tom.name + ' and ' + Jerry.name"
	p = compile(expr)
	assert.Equal(t, "Tom and Jerry", p.RunOnMemory(mem))

	expr = "Tom.age + Jerry.age"
	p = compile(expr)
	assert.Equal(t, int32(20), p.RunOnMemory(mem))

	expr = "Tom.age + Jerry.age"
	p = compile(expr)
	assert.Equal(t, int32(20), p.RunOnMemory(mem))

	expr = "Friends[0].name"
	p = compile(expr)
	assert.Equal(t, "friend1", p.RunOnMemory(mem))

	expr = "Friends[1].age"
	p = compile(expr)
	assert.Equal(t, int32(15), p.RunOnMemory(mem))

	expr = "Friends[1:][0].age"
	p = compile(expr)
	assert.Equal(t, int32(15), p.RunOnMemory(mem))

	expr = "Friends[0:3:2][1].age"
	p = compile(expr)
	assert.Equal(t, int32(15), p.RunOnMemory(mem))

	expr = "Friends[0:3:2][1].age * Tom.age"
	p = compile(expr)
	assert.Equal(t, int32(180), p.RunOnMemory(mem))

	expr = "Friends[0:3:2][1].age * Tom.age" // slice index of slice
	p = compile(expr)
	assert.Equal(t, int32(180), p.RunOnMemory(mem))

	expr = "Friends[0,0,0][2].age * Tom.age" // slice index of slice
	p = compile(expr)
	assert.Equal(t, int32(156), p.RunOnMemory(mem))

	expr = "Class[1001].name"
	p = compile(expr)
	assert.Equal(t, "Jackson", p.RunOnMemory(mem))

	expr = "Class[ [1001, 1002]][0].name" // slice index of map
	p = compile(expr)
	assert.Equal(t, "Jackson", p.RunOnMemory(mem))

	expr = "Class[DavidId].name"
	p = compile(expr)
	assert.Equal(t, "David", p.RunOnMemory(mem))
}

func TestAssign(t *testing.T) {
	setup()
	var expr string
	var p *program.SingleLineProgram

	expr = "newPerson == nil"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name == nil"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name == ''"
	p = compile(expr)
	assert.Equal(t, false, p.RunOnMemory(mem))

	expr = "newPerson.name = '\"a\r\n\"'"
	p = compile(expr)
	assert.Equal(t, "\"a\r\n\"", p.RunOnMemory(mem))

	expr = "newPerson.name == '\"a\r\n\"'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name = 'James'"
	p = compile(expr)
	assert.Equal(t, "James", p.RunOnMemory(mem))

	expr = "newPerson == nil"
	p = compile(expr)
	assert.Equal(t, false, p.RunOnMemory(mem))

	expr = "newPerson.name == nil"
	p = compile(expr)
	assert.Equal(t, false, p.RunOnMemory(mem))

	expr = "newPerson.name == 'James'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name, Teacher.name = 'Steve', 'Dannel'"
	p = compile(expr)
	assert.Equal(t, nil, p.RunOnMemory(mem))

	expr = "newPerson.name == 'Steve'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Teacher.name == 'Dannel'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Friends[1].age = 27.0"
	compile(expr).RunOnMemory(mem)
	expr = "Friends[1].age"
	assert.Equal(t, int32(27), compile(expr).RunOnMemory(mem))

	expr = "Friends[1].age == 27"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name = nil"
	assert.Panics(t, func() { compile(expr) })

	expr = "newPerson.name == ''"
	p = compile(expr)
	assert.Equal(t, false, p.RunOnMemory(mem))

	expr = "newPerson.name == nil"
	p = compile(expr)
	assert.Equal(t, false, p.RunOnMemory(mem))

	expr = "forMap[1] = 'Good'"
	compile(expr).RunOnMemory(mem)
	expr = "forMap[1] == 'Good'"
	assert.Equal(t, true, compile(expr).RunOnMemory(mem))

	expr = "Tom.age = Jerry.age++"
	compile(expr).RunOnMemory(mem)
	expr = "Tom.age == 8"
	assert.Equal(t, true, compile(expr).RunOnMemory(mem))
	expr = "Jerry.age == 9"
	assert.Equal(t, true, compile(expr).RunOnMemory(mem))

	expr = "Tom.age = --Jerry.age"
	compile(expr).RunOnMemory(mem)
	expr = "Tom.age == 8"
	assert.Equal(t, true, compile(expr).RunOnMemory(mem))
	expr = "Jerry.age == 8"
	assert.Equal(t, true, compile(expr).RunOnMemory(mem))

	expr = "Tom.name += 'my'"
	compile(expr).RunOnMemory(mem)
	expr = "Tom.name == 'Tommy'"
	assert.Equal(t, true, compile(expr).RunOnMemory(mem))
}

func TestInitializationList(t *testing.T) {
	setup()
	var expr string
	var p *program.SingleLineProgram

	expr = "newPerson = {name(nil), age(nil)}"
	p1 := compile(expr)
	p1.RunOnMemory(mem)

	expr = "newPerson"
	p1 = compile(expr)
	assert.Equal(t, map[string]interface{}{"age": int32(0), "name": ""}, p1.RunOnMemory(mem))

	expr = "newPerson = {name(Tom.name), age(12)}"
	p1 = compile(expr)
	p1.RunOnMemory(mem)

	expr = "newPerson.name == 'Tom'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name[:2] == 'To'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name[2] == 'm'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.age == 12"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson.name = 'Daniel'"
	p = compile(expr)
	assert.Equal(t, "Daniel", p.RunOnMemory(mem))

	assert.Equal(t, map[string]interface{}{
		"name": "Tom",
		"age":  int32(12),
	}, p1.RunOnMemory(mem))

	expr = "Friends = [{name('alpha'), age(100)}, {name('beta'), age(200)}, {name('gamma'), age(300)}]"
	p = compile(expr)
	p.RunOnMemory(mem)

	expr = "Friends[0].name == 'alpha'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Friends[1].name == 'beta'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Friends[2].name == 'gamma'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Class = {111:{name('alpha'), age(100)}, 222:{name('beta'), age(200), hobbies(['dance', 'sing']) }, 333:{name('gamma'), age(300)}}"
	p = compile(expr)
	p.RunOnMemory(mem)

	expr = "Class[111].name == 'alpha'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Class[222].name == 'beta'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Class[222].hobbies[1] == 'sing'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Class[333].name == 'gamma'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "stringMap = {'1001':{name('Jackson'), age(17)}, '1002':{name('David'), age(16)}}"
	compile(expr).RunOnMemory(mem)
	assert.Equal(t, true, compile("stringMap['1001'].age == 17").RunOnMemory(mem))

	expr = "float32Map = {1001:{name('Jackson'), age(17)}, 1002:{name('David'), age(16)}}"
	compile(expr).RunOnMemory(mem)
	assert.Equal(t, true, compile("float32Map[1001.0].age == 17").RunOnMemory(mem))
}

func TestFilter(t *testing.T) {
	setup()
	var expr string
	var p *program.SingleLineProgram

	expr = "Friends = [{name('alpha'), age(100)}, {name('beta'), age(200)}, {name('gamma'), age(300)}]"
	p = compile(expr)
	p.RunOnMemory(mem)

	expr = "Friends[?(@.name == 'beta')][0].age == 200"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Friends[?(@.name == 'gamma')][0].age == 300"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))
}

func TestFunction(t *testing.T) {
	setup()
	var expr string
	var p *program.SingleLineProgram

	expr = "RandNumber = random()"
	p = compile(expr)
	r1 := p.RunOnMemory(mem)

	expr = "RandNumber = random()"
	p = compile(expr)
	r2 := p.RunOnMemory(mem)

	assert.Equal(t, true, r1 != r2)

	expr = "num1 = 27"
	p = compile(expr)
	assert.Equal(t, int64(27), p.RunOnMemory(mem))

	expr = "num2 = 6"
	p = compile(expr)
	assert.Equal(t, int64(6), p.RunOnMemory(mem))

	expr = "num3, num4 = test1(num1, num2)"
	p = compile(expr)
	assert.Equal(t, nil, p.RunOnMemory(mem))

	expr = "num3"
	p = compile(expr)
	assert.Equal(t, int64(4), p.RunOnMemory(mem))

	expr = "num4"
	p = compile(expr)
	assert.Equal(t, int64(3), p.RunOnMemory(mem))

	expr = "startWith('China', 'Chi')"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "startWith('China', 'chi') || startWith('China', 'cHI')"
	p = compile(expr)
	assert.Equal(t, false, p.RunOnMemory(mem))

	expr = "!startWith('China', 'chi')"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "startWith('China', 'chi')"
	p = compile(expr)
	assert.Equal(t, false, p.RunOnMemory(mem))

	expr = "startWith('China', string(nil))"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "strCat(' ', 'Tom', 'and', 'Jerry')"
	p = compile(expr)
	assert.Equal(t, "Tom and Jerry", p.RunOnMemory(mem))

	expr = "strSplit('Tom and Jerry', ' ')"
	p = compile(expr)
	assert.Equal(t, []interface{}{"Tom", "and", "Jerry"}, p.RunOnMemory(mem))

	expr = "Class = {1001: {name('Jackson'), age(6), hobbies(['eat','drink','play','happy'])}, 1002:{name('David'), age(7)}}"
	compile(expr).RunOnMemory(mem)

	// marshal
	expr = "testString = string(JsonMarshal(Class))"
	p = compile(expr)
	fmt.Println("=== Json marshal result is not deterministic ===\nMarshal result" + p.RunOnMemory(mem).(string))

	// unmarshal
	expr = "NewClass = JsonUnmarshal(bytes(testString), typeof(Class))"
	compile(expr).RunOnMemory(mem)

	expr = "NewClass[1001].name == 'Jackson'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))
	expr = "NewClass[1001].age == 6"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))
	expr = "NewClass[1002].name == 'David'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))
	expr = "NewClass[1002].age == 7"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))
	expr = "NewClass[1001].hobbies[0] == 'eat'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))
	expr = "NewClass[1001].hobbies[2] == 'play'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	// jsonObject
	expr = "jsonObj = JsonObject(bytes(testString))"
	compile(expr).RunOnMemory(mem)

	// jsonpath
	expr = "testString = JsonPath(jsonObj, '$.1001.name')"
	compile(expr).RunOnMemory(mem)
	expr = "testString == 'Jackson'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	// jsonpath
	expr = "testString = JsonPath(jsonObj, '$.1001.hobbies[ 3 ]')"
	compile(expr).RunOnMemory(mem)
	expr = "testString == 'happy'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "testString = EncodeBase64(bytes('abcd1234-~1'))"
	compile(expr).RunOnMemory(mem)
	expr = "testString2 = DecodeBase64(testString)"
	compile(expr).RunOnMemory(mem)
	assert.Equal(t, "abcd1234-~1", compile("testString2").RunOnMemory(mem))
}

func TestSlice(t *testing.T) {
	setup()
	var expr string
	var p *program.SingleLineProgram

	expr = "newPerson = {name('GOOD'), age(100)}"
	p = compile(expr)
	p.RunOnMemory(mem)

	expr = "pushBack(Friends, newPerson)"
	p = compile(expr)
	p.RunOnMemory(mem)

	expr = "Friends[3].name == 'GOOD'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "newPerson = {name('BAD'), age(200)}"
	p = compile(expr)
	p.RunOnMemory(mem)

	expr = "pushFront(Friends, newPerson)"
	p = compile(expr)
	p.RunOnMemory(mem)

	expr = "Friends[4].name == 'GOOD'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "Friends[0].name == 'BAD'"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "len(Friends)"
	p = compile(expr)
	assert.Equal(t, int64(5), p.RunOnMemory(mem))

	expr = "typeof(Friends[0]) == typeof(Tom)"
	p = compile(expr)
	assert.Equal(t, true, p.RunOnMemory(mem))

	expr = "itoa(int64(3.1))"
	p = compile(expr)
	assert.Equal(t, "3", p.RunOnMemory(mem))

	expr = "atoi('123')"
	p = compile(expr)
	assert.Equal(t, int64(123), p.RunOnMemory(mem))

	expr = "ftoa(3.1)"
	p = compile(expr)
	assert.Equal(t, "3.100000", p.RunOnMemory(mem))

	expr = "atof('123.7')"
	p = compile(expr)
	assert.Equal(t, float64(123.7), p.RunOnMemory(mem))

	expr = "Tom = new Person()"
	p = compile(expr)
	assert.Equal(t, map[string]interface{}{}, p.RunOnMemory(mem))

	expr = "Tom = {name('alpha'), age(100)}"
	compile(expr).RunOnMemory(mem)
	expr = "delete(Tom, 'name')"
	compile(expr).RunOnMemory(mem)
	expr = "Tom"
	p = compile(expr)
	assert.Equal(t, map[string]interface{}{"age": int32(100)}, p.RunOnMemory(mem))
}

func TestEnum(t *testing.T) {
	setup()
	var expr string
	var p *program.SingleLineProgram

	expr = "fruitEnum = fruits.banana"
	p = compile(expr)
	assert.Equal(t, int32(1), p.RunOnMemory(mem))

	expr = "enumString(fruitEnum)"
	p = compile(expr)
	assert.Equal(t, "banana", p.RunOnMemory(mem))
}

func BenchmarkExpr(b *testing.B) {
	setup()
	expr := "Friends[1].age * Tom.age"
	p := compile(expr)
	stk := &common.Stack{
		Bp:   0,
		Sp:   -1,
		Data: make([]interface{}, 0, 100),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Run(mem, stk)
		stk.Pop()
	}
}

func BenchmarkExprStackPool(b *testing.B) {
	setup()
	expr := "Friends[1].age * Tom.age"
	p := compile(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemory(mem)
	}
}

func BenchmarkExprStackPoolNoPanic(b *testing.B) {
	setup()
	expr := "Friends[1].age * Tom.age"
	p := compile(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemoryNoPanic(mem)
	}
}

func BenchmarkExprFunc(b *testing.B) {
	setup()
	expr := "testString = 'www.google.com'"
	compile(expr).RunOnMemory(mem)
	expr = "startWith(testString, 'www.g')"
	p := compile(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemory(mem)
	}
}

func BenchmarkExpr2(b *testing.B) {
	setup()
	expr := "Friends[0].age == 14 || Tom.age == 12 || Friends[2].age == 15"
	p := compile(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemory(mem)
	}
}

func BenchmarkInitializationList(b *testing.B) {
	setup()
	expr := "Class = {1:{name('alpha'), age(100)}, 2:{name('beta'), age(200), hobbies(['dance', 'sing']) }, 3:{name('gamma'), age(300)}}"
	p := compile(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemory(mem)
	}
}

func BenchmarkInitNaive(b *testing.B) {
	setup()
	p1 := compile("Class = new('map<int64,Person>')")
	p2 := compile("Class[1].name = 'alpha'")
	p3 := compile("Class[1].age = 100")
	p4 := compile("Class[2].name = 'beta'")
	p5 := compile("Class[2].age = 200")
	p5_1 := compile("Class[2].hobbies = ['dance', 'sing']")
	p6 := compile("Class[3].name = 'gamma'")
	p7 := compile("Class[3].age = 300")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p1.RunOnMemory(mem)
		p2.RunOnMemory(mem)
		p3.RunOnMemory(mem)
		p4.RunOnMemory(mem)
		p5.RunOnMemory(mem)
		p5_1.RunOnMemory(mem)
		p6.RunOnMemory(mem)
		p7.RunOnMemory(mem)
	}
}

func BenchmarkFilter(b *testing.B) {
	setup()
	expr := "Friends = [{name('alpha'), age(100)}, {name('beta'), age(200), hobbies(['dance', 'sing']) }, {name('gamma'), age(300)}]"
	p := compile(expr)
	p.RunOnMemory(mem)
	expr = "Friends[?(@.name == 'beta')][0].age"
	p = compile(expr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.RunOnMemory(mem)
	}
}

var c *ast.Compiler
var mem *common.Memory

func setup() {
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
		Data: make([]interface{}, 100),
	}

	scope := goscript.NewScope(nil)
	c.Scope = scope

	c.Include("common")
	c.Include("string")
	c.Include("json")
	c.Include("base64")
	c.Include("datetime")

	scope.AddVariable(goscript.NewVariable("Tom", c.FindType("Person")))
	scope.AddVariable(goscript.NewVariable("Jerry", c.FindType("Person")))
	scope.AddVariable(goscript.NewVariable("Friends", c.FindSliceType("Person")))
	scope.AddVariable(goscript.NewVariable("Class", c.FindMapType("int64", "Person")))
	scope.AddVariable(goscript.NewVariable("newPerson", c.FindType("Person")))
	scope.AddVariable(goscript.NewVariable("DavidId", c.FindType("int64")))
	scope.AddVariable(goscript.NewVariable("Teacher", c.FindType("Person")))
	scope.AddVariable(goscript.NewVariable("RandNumber", c.FindType("float64")))
	scope.AddVariable(goscript.NewVariable("num1", c.FindType("int64")))
	scope.AddVariable(goscript.NewVariable("num2", c.FindType("int64")))
	scope.AddVariable(goscript.NewVariable("num3", c.FindType("int64")))
	scope.AddVariable(goscript.NewVariable("num4", c.FindType("int64")))
	scope.AddVariable(goscript.NewVariable("testString", c.FindType("string")))
	scope.AddVariable(goscript.NewVariable("NewClass", c.FindMapType("int64", "Person")))
	scope.AddVariable(goscript.NewVariable("jsonObj", c.FindType("JSONObject")))
	scope.AddVariable(goscript.NewVariable("testString2", c.FindType("string")))
	scope.AddVariable(goscript.NewVariable("header", c.FindMapType("string", "string")))
	scope.AddVariable(goscript.NewVariable("fruitEnum", c.FindType("fruits")))
	scope.AddVariable(goscript.NewVariable("stringMap", c.FindMapType("string", "Person")))
	scope.AddVariable(goscript.NewVariable("float32Map", c.FindMapType("float32", "Person")))
	scope.AddVariable(goscript.NewVariable("forMap", c.FindMapType("int32", "string")))

	c.Scope = scope

	// init some data
	var expr string
	expr = "Tom = {name('Tom'), age(12)}"
	compile(expr).RunOnMemory(mem)
	expr = "Jerry = {name('Jerry'), age(8)}"
	compile(expr).RunOnMemory(mem)
	expr = "Friends = [{name('friend1'), age(13)}, {name('friend2'), age(15)}, {name('friend3'), age(15)}]"
	compile(expr).RunOnMemory(mem)
	expr = "Class = {1001:{name('Jackson'), age(17)}, 1002:{name('David'), age(16)}}"
	compile(expr).RunOnMemory(mem)
	expr = "DavidId = 1002"
	compile(expr).RunOnMemory(mem)
}
