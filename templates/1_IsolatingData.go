package templates

import "fmt"

func IsolatingData() {
	// makeFibGen()s local variables will sustain its values till gen is in scope.
	gen := makeFibGen()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}
func IsolatingData2() {
	next := incrementor()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func makeFibGen() func() int {
	// f1 & f2 will be unique of every variable of type makeFibGen()
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

func makeFibGen2() int {
	// f1 & f2 will be unique of every variable of type makeFibGen()
	f1 := 0
	f2 := 1
	f2, f1 = (f1 + f2), f2
	return f1
}
