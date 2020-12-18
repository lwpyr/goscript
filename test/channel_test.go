package test

import "testing"

func TestChannel_A(t *testing.T) {
	setupClean()
	var expr string
	expr = `
type Person message {
	name    string
	age     int32
}

var personChan chan[Person] = new chan[Person](1);

var Tom Person = {name:'Tommy', age:12};

personChan <<- Tom;
var TomCopy Person = <<- personChan;
print(TomCopy);
TomCopy.age = 15;
print(Tom);
print(TomCopy);
`
	p := compileScript(expr)
	p.RunOnMemory(mem)
}
