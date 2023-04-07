package main

import "fmt"

func findEven(n []int) []int {
	var result []int
	for i := 0; i < len(n); i++ {
		if n[i]%2 == 0 {
			result = append(result, n[i])
		}
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	var numbers []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}
	fmt.Print(findEven(numbers))
}