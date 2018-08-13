package main

import "fmt"

func nextInt(b []byte, i int) (int, int) {
	x := 0
	x = x*10 + int(b[i]) - '0'
	i++
	return x, i
}

func main() {
	a := []byte{'1', '2', '3', '4', '5'}
	var x int
	for i := 0; i < len(a); {
		x, i = nextInt(a, i)
		fmt.Println("x = ", x, " index = ", i)
	}
}
