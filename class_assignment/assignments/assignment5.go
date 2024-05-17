package assignments

import (
	"fmt"
	"sync"
)

var wg5 sync.WaitGroup

func producer5(num chan int, done chan int) {
	for i := 1; i <= 25; i++ {
		<-done
		fmt.Print("produce ", i, " ")
		num <- i
	}
	<-done
	wg5.Done()
}

func square(num chan int, num4cube chan int) {
	for i := 1; i <= 25; i++ {
		v := <-num
		fmt.Print(" square ", v*v, " ")
		num4cube <- v

	}
	defer close(num)
	// defer close(done)
}

func cube(num4cube chan int, done chan int) {
	for i := 1; i <= 25; i++ {
		v := <-num4cube
		fmt.Println(" cube ", v*v*v)
		done <- 1
	}

	defer close(num4cube)

}

func Run5() {

	num := make(chan int)
	num4cube := make(chan int)
	done := make(chan int)

	// done <- 1
	go func() {
		done <- 1
	}()
	wg5.Add(1)
	go producer5(num, done)
	go square(num, num4cube)
	go cube(num4cube, done)
	// time.Sleep(1 * time.Second)
	wg5.Wait()

}