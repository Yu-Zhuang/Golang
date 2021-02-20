package main

import "fmt"

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() (r int) {
	r = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func main() {
	var s stack
	fmt.Printf("stack value: %v ;; stack type: %T\n", s, s)

	for i := 0; i < 10; i += 2 {
		s.push(i)
		fmt.Println(s)
	}

	fmt.Println(`====================`)

	for range s {
		fmt.Println(s.pop(), s)
	}
}
