package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func makeRandomSlice(numItems, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func checkSorted(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

// Perform linear search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func linearSearch(slice []int, target int) (index, numTests int) {
	var i int
	for i = 0; i < len(slice); i++ {
		if target == slice[i] {
			return i, i + 1
		}
	}
	return -1, i
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	var target string
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	for {
		fmt.Printf("Target: ")
		_, err := fmt.Scanln(&target)
		if err != nil {
			break
		}

		if len(target) == 0 {
			break
		}

		ntarget, err := strconv.Atoi(target)
		if err != nil {
			fmt.Println(err)
			continue
		}

		i, n := linearSearch(slice, ntarget)
		if i == -1 {
			fmt.Printf("Target %d not found, %d tests\n", ntarget, n)
		} else {
			fmt.Printf("values[%d] = %d, %d tests\n", i, ntarget, n)
		}

	}
}
