package templates

import "fmt"

func DeferringWork() {
	a := 1
	b := 1
	go func() {
		result := doWork1(a, b)
		result = doWork2(result)
		result = doWork3(result)
		// Use the final result
	}()
	fmt.Println("hi!")
}

func doWork1(a, b int) int {
	return a + b
}

func doWork2(c int) int {
	return c * c
}

func doWork3(c int) int {
	return c * 2
}
