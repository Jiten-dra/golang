package assignments

import (
	"fmt"
	"sync"
)

/*
Assignment 1: Goroutine with Channel Problem:
Write a Go program that calculates the sum of numbers from 1 to N concurrently using goroutines and channels.
The program should take the value of N as input from the user.
*/

var wg sync.WaitGroup

func producer(n int, nums chan<- int) {
	defer wg.Done()
	defer close(nums)

	for i := 0; i < n; i++ {
		fmt.Println("insert to nums ", i)
		nums <- i
	}
}

func calculateSum(nums <-chan int) <-chan int {
	defer wg.Done()
	sum := 0
	ans := make(chan int)
	go func() {
		for num := range nums {
			fmt.Println("added to sum ", sum)
			sum += num
		}
		ans <- sum
	}()
	return ans
}

func Run1() {
	wg.Add(2)
	nums := make(chan int)
	go producer(25, nums)
	ans := calculateSum(nums)
	fmt.Println("Sum is: ", <-ans)
	wg.Wait()
}
