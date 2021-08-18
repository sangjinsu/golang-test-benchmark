package main

import "fmt"

func square(x int) int {
	return x * x
}

func fibo(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return b
}

func fiboRecursive(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	return fiboRecursive(n -1) + fiboRecursive( n - 2)
}

func main() {
	fmt.Printf("9 * 9 = %d\n", square(9))

	fmt.Println(fibo(5))
	fmt.Println(fiboRecursive(5))
}
