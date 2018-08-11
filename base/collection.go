package main

import "fmt"

func main() {

	var arr [10]int
	arr[0] = 42
	arr[1] = 43
	fmt.Println("The first element is %d", arr[0])

	a := [2][2]int{ [2]int{1,2}, [2]int{3,4}}
	fmt.Println(a)

	a1 := [2][2]int{ [...]int{1,2}, [...]int{3,4}}
	fmt.Println(a1)

	slice := arr[0:2]
	arr[0] = 46
	fmt.Println(slice)
	fmt.Println("slice length: %d", len(slice))
	fmt.Println("slice cap: ", cap(slice))

	a2 := [...]int{1, 2, 3, 4, 5}
	// a2[n:m] 从 array 创建了一个 slice ，具有元素 n 到 m-1
	s1 := a2[2:4]
	s2 := a2[1:5]
	s3 := a2[:]
	s4 := a2[:4]
	s5 := a2[:]
	fmt.Println(" s +++++++++++++++ start ++++++++")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(" s +++++++++++++++ end ++++++++")


}
