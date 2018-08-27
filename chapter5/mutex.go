package main

import "fmt"

type Mutex struct {
}

func (m *Mutex) Lock()   {}
func (m *Mutex) Unlock() {}
func (m *Mutex) get0() (a int) {
	a = 5
	return
}

type NewMutex Mutex

type PrintableMutex struct{ Mutex }

func main() {
	mutex := new(Mutex)
	newMutex := new(NewMutex)
	printableMutex := new(PrintableMutex)

	fmt.Println(mutex.Lock)
	fmt.Println(newMutex)
	fmt.Println(printableMutex.Lock)
	fmt.Println(printableMutex.get0())

}
