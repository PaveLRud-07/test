package main

import (
	"fmt"
	"math"
)

func Prime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
func findPrimesInRange(min, max int) []int {
	primes := []int{}
	for num := min; num <= max; num++ {
		if Prime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func main() {
	min := 11
	max := 20
	primesInRange := findPrimesInRange(min, max)
	fmt.Printf("Простые числа в диапазоне от %d до %d: %v\n", min, max, primesInRange)
}
