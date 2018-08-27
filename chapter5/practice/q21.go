package main

import (
	"container/list"
)

func main() {

	numbers := list.New()
	numbers.PushFront(1)
	numbers.PushFront(3)
	numbers.PushFront(1)
}
