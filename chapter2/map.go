package main

import "fmt"

func main() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}

	monthdays["Mi"] = 69
	fmt.Println(monthdays)

	year := 0
	for _, days := range monthdays {
		year += days
	}
	fmt.Println("Number of days in a year: ", year)

	monthdays["Undecim"] = 30
	monthdays["Feb"] = 29

	var value int
	var present bool

	value, present = monthdays["Jan"]
	fmt.Println("value = ", value, " present : ", present)

	v, ok := monthdays["Jan"]
	fmt.Println("v = ", v, " ok = ", ok)

	fmt.Println("monthdays = ", monthdays)
	delete(monthdays, "Mar")
	fmt.Println("monthdays = ", monthdays)

}
