package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	//Only this is called. The rest is closed over.
	//Quite ingenious. a will be switched to the next iteration of b
	//and b will be the a + b and b is returned.
	return func() int {
		a, b = b, a+b
		return b
	}
}

// I wrote this sample recursive so I see the differences.
//func fib (n int) int {
//	if (n == 1 || n == 2) { return 1 }
//	return fib(n - 1) + fib(n - 2)
//}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
