# Goscript: a runtime script for Golang
Goscript is completely written in Golang. Goscript's purpose is to build a completely data-compatible-with-golang script language and could dynamicly extend your runtime code in a Golang server.

## NOTICE: Goscript is in a very early stage, there are still many thing to be done and many bugs to be fixed.

### TODO
check lhs flag cannot be set to intermediate result (throw panic).

### TODO LIST

- [x] Type System: Basic value type and basic reference type(Message, Map, Slice)
- [x] Json marshal and unmarshal of typed data
- [x] Basic Ops
- [x] Function Call
- [x] Control statement: if, switch, for loop
- [x] Function Definition
- [x] Lambda function and function type
- [x] type define in script
- [x] Golang channel type
- [ ] Unit test

#### consider to support
- [ ] Protocol Buffer marshal and unmarshal of typed data
- [ ] Golang Struct type (actually reflect struct)
- [ ] Package support (could use 'import')
- [ ] Support using context.Context to control timeout
- [ ] More function libs

### Design principle
1. Goscript is a simple and type-strict with implicit type-conversion script language with Golang-similar grammar for Golang.

2. Goscript should be easily extended for specific usage.

3. Goscript should be fast, data conversion when context switching should be easy.

### What can Goscript do
1. **As a rule parse engine.**

    Goscript has two modes, one is Single-line program mode and another is Script mode. Single-line mode is just one executable statement.
    
    With Single-line mode, Goscript could do evaluation jobs. For example 
    
    ```shell
    "Request.Method=='PUT' && Request.Body != nil"
    ```
    
    

2. **As a extension for your running server.**

   Some times we want to update some server-side logic without stop the server. Goscript's script mode can meet this cand of need.

   ```shell
   if (Player.VIP == true) {
   	Player.offer = 100;
   } else if (Player.TimeSpentLastDay > 3600) {
   	Player.offer = 80;
   }
   ```

   

### What data structure does Goscript use

If you don't extend Gosript's type registry with customed object type, Goscript only use slice, map and basic type in Golang. Very simple and no-coversion in context switching between Goscript and Golang runtime.

1. Basic type: int32, int64, uint32, uint64, string, bool and etc.
2. Message type: this is like a struct in Golang, the underlying data structure is map[string]interface{}
3. Map type: depends on the key, Map type could be map[string]interface{}, map[int32]interface{} and etc.

4. Slice type: underlying data structure is []interface{}
5. Object type: this is for extend usage. For example there is a JsonObject type in json lib of Goscript.

### What is the performance?

Not slow, not fast, but good enough in pure-golang implemented script language

all the experiment is done under MacBook Pro 16 2019, Intel Core i9, 16GB

1. Single-line mode compared with other evaluation lib

   [antonmede/expr](https://github.com/antonmedv/expr)

   ```shell
   #Friends = [{name('friend1'), age(13)}, {name('friend2'), age(15)}, {name('friend3'), age(15)}]
   #Tom = {name('Tom'), age(12)}
   
   #expr:= "Friends[1].age * Tom.age"
   
   Goscript    	16122072	        71.9 ns/op
   AntonmedeExpr    	 2356504	       449 ns/op
   
   #expr := "Friends[0].age == 14 || Tom.age == 12 || Friends[2].age == 15"
   Goscript    	11847440	        96.7 ns/op
   AntonmedeExpr    	 2481778	       486 ns/op
   ```

   Goscript has overwhelming advantage over antonmede/expr in performance side. But use Goscript is complexer than antonmede/expr, all the variables and types should be pre-defined. antonmede/exp directly process the data with no prior knowledge over the data, and this may be the key reason of why Goscript is faster.

1. Script mode compared with other Golang script
   
   [yuin/gopher-lua](https://github.com/yuin/gopher-lua)

   a. we use fib(35) to compare with GopherLua and Python3.8 (loop 10 times and take the average)

   ```shell
   # Goscript
   # ===
   # func fib(v int64) int64 {
   # 	if(v<2) {
   # 		return v;
   # 	}
   # 	return fib(v-1)+fib(v-2);
   # }
   # 
   # print(fib(35));
   # ===
   Goscript 2.67 sec
   
   # GopherLua
   # ===
   # local function fib(n)
   #     if n < 2 then return n end
   #     return fib(n - 2) + fib(n - 1)
   # end
   # 
   # print(fib(35))
   # ===
   GopherLua 3.69 sec
   
   # native python 3.8
   # ===
   # def fib(n):
   # if n < 2:
   #       return n
   # return fib(n - 2) + fib(n - 1)
   #
   # print(fib(35))
   # ===
   Python3.8 2.36 sec
   ```

   Faster than GopherLua, a little slower than native Python3.8

   

   b. calculate sum from 0 to 100

   ```shell
   # Goscript
   # ===
   # var sum int64;
   # sum = 0;
   # for (local i int64 = 1; i <= 100; i++) {
   # 	sum = sum + i;
   # }
   # ===
   Goscript 9637 ns/op
   
   # GopherLua
   # ===
   # sum = 0
   # for i=1,100 do
   #     sum = sum + i
   # end
   # ===
   GopherLua 123201 ns/op
   ```

   I don't have a close look into GopherLua, even though I already use **compile** mode of GopherLua, and I don't know why GopherLua perform badly in this simple case.

### Example code

see more in the **test** folder

1. Single-line mode

```
// todo
```



1. Script-mode

lambda calculus
```go
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

print('=== test BASIC ===');
print(Identity == Identity(Identity)); // true
print(Identity == SelfApply(Identity)); // true

print('=== test APPLY ===');
print(Apply(Identity).(func(object)object)(Identity) == Identity); // true

print('=== test BOOL ===');
print(TRUE(Identity).(func(object)object)(Apply) == Identity); // true
print(FALSE(Identity).(func(object)object)(Apply) == Apply); // true

print('=== test COND ===');
print(CONDITION(Identity).(func(object)object)(Apply).(func(object)object)(TRUE) == Identity); // true
print(CONDITION(Identity).(func(object)object)(Apply).(func(object)object)(FALSE) == Apply); // true
print(CONDITION(Apply).(func(object)object)(Identity).(func(object)object)(TRUE) == Apply); // true
print(CONDITION(Apply).(func(object)object)(Identity).(func(object)object)(FALSE) == Identity); // true

print('=== test NOT ===');
print(NOT(TRUE) == FALSE); // true
print(NOT(NOT(TRUE)) == TRUE); // true

print('=== test AND ===');
print(AND(TRUE).(func(object)object)(TRUE) == TRUE); // true
print(AND(TRUE).(func(object)object)(FALSE) == FALSE); // true
print(AND(FALSE).(func(object)object)(TRUE) == FALSE); // true
print(AND(FALSE).(func(object)object)(FALSE) == FALSE); // true

print('=== test OR ===');
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

print('=== test IsZero ===');
print(IsZero(zero) == TRUE); // true
print(IsZero(one) == FALSE); // true

func Prev(x object) object {
	return IsZero(x).(func(object)object)(zero).(func(object)object)(x.(func(object)object)(FALSE));
}

print('=== test Prev ===');
print(Prev(one) == zero); // true
print(Prev(nine) == eight); // true
print(Prev(nine) == seven); // false
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}
```


## Someone wants to contribute?

Awesome! Goscript is an open source project and contribution is welcomed!