package main

import "fmt"

func main() {
	fmt.Printf("%f\n", myPow(5, 25))
	fmt.Println(myPow(2, -10))
	fmt.Println(myPow(3, 0))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n = 0 - n
	}

	result := myPow(x*x, n>>1)
	if n%2 == 1 {
		result = result * x
	}

	return result
}
