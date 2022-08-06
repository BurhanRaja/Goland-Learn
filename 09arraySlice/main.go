package main

import (
	"fmt"
	"sort"
)

func main() {
	// Arrays
	var fruitListArr [5]string
	fruitListArr[0] = "Mango"
	fruitListArr[1] = "Peru"
	fruitListArr[2] = "Apple"
	fruitListArr[3] = "Banana"
	fruitListArr[4] = "Watermelon"

	fmt.Println(fruitListArr)

	vegList := [3]string{"Spinach", "Cauliflower", "Potato"}

	fmt.Println(vegList)

	fmt.Println()

	// Slice
	fruitListSlice := []string{"Apple", "Banana", "Mango"}
	fmt.Printf("Type of fruitListSlice %T\n", fruitListSlice)

	fruitListSlice = append(fruitListSlice, "Watermelon", "Peru")
	fmt.Println(fruitListSlice)

	// Memory Allocation
	highScoresSlice := make([]int, 5)
	highScoresSlice[0] = 234
	highScoresSlice[1] = 577
	highScoresSlice[2] = 434
	highScoresSlice[3] = 976
	highScoresSlice[4] = 324

	highScoresSlice = append(highScoresSlice, 134, 289, 300)

	fmt.Println(highScoresSlice)
	sort.Ints(highScoresSlice)
	fmt.Println(highScoresSlice)

	// Remove Mango from fruitListSlice
	index := 2 // given index
	fruitListSlice = append(fruitListSlice[:index], fruitListSlice[index+1:]...)
	fmt.Println(fruitListSlice)
}
