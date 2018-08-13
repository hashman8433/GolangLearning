package main

func main() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}

	println("f() = ", f())
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}
