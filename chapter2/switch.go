package main

func main() {

	switch 0 {
	case 0:
		println("0")
	case 1:
		println("1")
	default:
		println("default")
	}

	println("+++++++++++ fall through ++++++++++++++")

	switch 0 {
	case 0:
		println("0")
		fallthrough
	case 1:
		println("1")
		fallthrough
	default:
		println("default")
	}
}
