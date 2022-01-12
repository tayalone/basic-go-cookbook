package main

import "fmt"

func sumOfFirst(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	checkPoint := n / 2
	for i := 2; i <= checkPoint; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func tryInfinityLoop() {
	i := 0
	for {
		i += 1
		if i == 10000 {
			break
		}
	}
	fmt.Println("I'll break when i == 10000, and i =", i)
}

func main() {
	fmt.Println(sumOfFirst(3))
	fmt.Println(isPrime(1))
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(4))
	tryInfinityLoop()
}
