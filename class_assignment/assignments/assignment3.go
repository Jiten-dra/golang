package assignments

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var sg sync.WaitGroup

var num1 int
var num2 int

func withoutMutex() {
	// defer sg.Done()
	sg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer sg.Done()
			temp := num1
			temp++
			num1 = temp
		}()
	}

}

func withMutex() {
	// defer sg.Done()
	sg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer sg.Done()
			mu.Lock()
			temp := num2
			temp++
			num2 = temp
			mu.Unlock()
		}()
	}
}

func Run3() {
	// sg.Add(2)
	withMutex()
	withoutMutex()
	sg.Wait()
	fmt.Println("final value without mutex: ", num1)
	fmt.Println("final value without mutex: ", num2)

}
