package main

func main() {
	// Q2 (1)For-loop
	println("================= Q2 start =============== \n")

	// 1
	println("====== 1.for =======")
	println("  for1 ++++++++++")
	for i := 0; i < 10; i++ {
		print(i, " ")
	}

	println("\n  for2 ++++++++++")
	number := 0
	for number < 10 {
		number++
		print(number, " ")
	}

	println("\n  for3 ++++++++++")
	for {
		print(number, " ")
		number--
		if number < 0 {
			break
		}
	}

	// 2
	println("\n====== 2.goto =======")
	number = 0
Here:
	number++
	print(number, " ")
	if number < 10 {
		goto Here
	}

	// 3
	println("\n====== 3.array =======")
	numbers := [...]int{0, 1, 2, 3, 4, 5, 6}
	for n := range numbers {
		print(n, " ")
	}

	println("\n================= Q3 start =============== \n")
	for i := 0; i < 100; i++ {

		multiple3 := (i%3 == 0)
		multiple5 := (i%5 == 0)

		if multiple3 && multiple5 {
			println("FizzBuzz")
			continue
		}

		if multiple3 {
			println("Fizz")
			continue
		}

		if multiple5 {
			println("Buzz")
			continue
		}
		println(i)

	}

	println("================= Q4 start =============== \n")
	println("\n====== 1.array =======")

}
