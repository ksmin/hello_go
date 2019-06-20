package main

import "fmt"

func Sum(arr *[3]float64) (sum float64) {
	for idx, val := range *arr {
		fmt.Println("idx:", idx, "val:", val)
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("Hello, world.")

	// make example
	sliceA := make([]int, 0, 3)

	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println("sliceA len:", len(sliceA), "cap:", cap(sliceA))
	}
	fmt.Println("sliceA:", sliceA)

	sliceB := []int{4, 5, 6}
	sliceA = append(sliceA, sliceB...)
	fmt.Println("sliceA:", sliceA)

	target := make([]int, len(sliceA), cap(sliceA))
	copy(target, sliceA)
	fmt.Println("target:", target)

	// Array example
	array := [...]float64{7.0, 8.5, 9.1}
	sum := Sum(&array)
	fmt.Println("sum:", sum)
}
