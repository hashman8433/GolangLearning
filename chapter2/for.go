package main

func main() {
	//i := 0
	//Here:
	//	println(i)
	//	i += 11000000
	//	goto Here

	//1. for init; condition; post {}
	//2. for condition
	//3. for {}

	for i := 0; i < 10; i++ {
		println(i)
	}

	a := [10]int{1, 3, 6, 9, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		println("i = ", i, "  j = ", j)
	}

	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		println("k = ", k, " v = ", v)
	}

}
