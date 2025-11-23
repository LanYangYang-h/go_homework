package main

import "fmt"

func once_number(arr []int) int {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}
	for num, count := range counts {
		if count == 1 {
			return num
		}
	}
	return -1
}

func main() {

	arr := []int{4, 5, 6, 7, 5, 6, 4}
	result := once_number(arr)
	fmt.Println("The number that appears only once is:", result)
}
