package practice

import "testing"

func TestPushPop(t *testing.T) {
	c := new(Stack)
	c.push(5)
	println(c.pop())
}
