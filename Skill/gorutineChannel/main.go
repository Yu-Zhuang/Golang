package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	num := 10
	go runOne(num, c1)
	go runTwo(num, c2)
	for i := 1; i < num; i++ {
		ret1, ret2 := <-c1, <-c2
		fmt.Printf("%d-1: runOne(%d) result : %d\n", i, num, ret1)
		fmt.Printf("%d-2: runTwo(%d) result : %d\n", i, num, ret2)
	}
	fmt.Println("end")
}

func runOne(num int, c chan int) {
	fmt.Println("--- run One Start--- ")
	ret := 1
	for i := 1; i < num; i++ {
		ret += i
		c <- ret
	}
}

func runTwo(num int, c chan int) {
	fmt.Println("--- run Two Start--- ")
	ret := 1
	for i := 1; i < num; i++ {
		ret *= i
		c <- ret
	}
}
