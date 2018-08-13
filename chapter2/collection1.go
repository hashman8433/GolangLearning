package main

import "fmt"

func main() {
	var array [100]int
	slice := array[0:99]
	slice[98] = 'a'
	//slice[99] = 'a' // error index out of range

	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 6, 7)
	s3 := append(s2, s0...)
	fmt.Println(s3)

}
