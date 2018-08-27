package main

import "fmt"

type NameAge struct {
	name string
	age  int
}

func main() {
	a := new(NameAge)
	a.name = "Pate"
	a.age = 42

	fmt.Printf("%v\n", a)
	fmt.Printf("%s", a.name)

}
