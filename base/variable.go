package main

func control () int {
	x := 5
	if x > 0 {
		return 5
	} else {
		return 10
	}
}

func controlIf() {
	if true && true {
		println("true")
	}

	if ! false {
		println("true")
	}
}

func main() {
	control()
	controlIf()
}