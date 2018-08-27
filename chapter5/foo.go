package main

import "fmt"

type foo struct{ int }
type bar foo

func main() {
	var b bar = bar{1}
	//var f foo = b
	var f foo = foo(b)
	fmt.Printf("%v", f)
}
