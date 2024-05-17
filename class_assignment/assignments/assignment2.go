package assignments

import (
	"fmt"
)

/*
Assignment 1: Goroutine with Channel Problem:
Write a Go program that calculates the sum of numbers from 1 to N concurrently using goroutines and channels.
The program should take the value of N as input from the user.
*/

func producer2(nums chan<- int) {
	defer close(nums)

	for i := 1; i < 101; i++ {
		nums <- i
	}
}

func consumer2(nums <-chan int) {
	defer wg.Done()
	for num:= range nums {
		fmt.Println(num)
	}
}

func Run2() {
	wg.Add(1)
	nums := make(chan int)
	go producer2(nums)
	go consumer2(nums)
	wg.Wait()
}
