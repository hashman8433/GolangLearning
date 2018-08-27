package main

import "fmt"

func main() {

	b := []byte{'h', 'e', 'l', 'l', 'o'}
	s := string(b)
	i := []rune{257, 1024, 65}
	r := string(i)

	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", r)

}
