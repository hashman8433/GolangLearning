package main

import "fmt"

func main() {

	var p *[]int = new([]int)
	//var v []int = make([]int, 100)

	*p = make([]int, 100, 100)
	v := make([]int, 100)
	v[0] = 6

	fmt.Println(*p)
	fmt.Println(v)

}
