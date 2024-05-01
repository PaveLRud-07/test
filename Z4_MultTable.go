//go lang

package main

import "fmt"

func MultTable(n int) {

	fmt.Printf("Таблица умножения до %d:\n", n)

	fmt.Print("   ")
	for i := 1; i <= n; i++ {
		fmt.Printf(" %2d  ", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%2d ", i)

		for j := 1; j <= n; j++ {
			result := i * j
			fmt.Printf(" %2d  ", result)
		}
		fmt.Println()
	}
}

func main() {
	MultTable(5)
}
