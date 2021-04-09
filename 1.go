package main

import "fmt"

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

type LazyMultipleGenerator func() int

func MakeLazyMultipleGenerator() LazyMultipleGenerator {
	v := 0
	f := nextMultiple
	return func() int {
		v = f(v + 1)
		return v
	}
}

func nextMultiple(n int) int {
	if n%3 == 0 || n%5 == 0 {
		return n
	}
	return nextMultiple(n  + 1)
}

func main() {
	sum := 0
	limit := 400000000
	generator := MakeLazyMultipleGenerator()
	for i := generator(); i < limit; i = generator() {
		sum += i
	}
	fmt.Println(sum)
}
