package main

import (
	"fmt"
	"math"
)

// Function to find the duplicate number using a map
func findDuplicate(nums []int) int {
	// Create a map to track seen numbers
	m := make(map[int]bool)
	fmt.Println("Map: ", m)
	// Iterate over the numbers
	for _, num := range nums {
		// Check if the number is already in the map
		if _, exists := m[num]; exists {
			// If it exists, it's a duplicate, return it
			fmt.Println("Duplicate found:", num)
			return num
		}
		// If it doesn't exist, add it to the map
		m[num] = true
		// Log the current state of the map
		fmt.Println("Map state:", m)
	}
	// Return -1 if no duplicate is found
	return -1
}

// Function to find the duplicate number by marking the array
func findDiuplicateMark(nums []int) int {
	// Iterate over the numbers
	for _, num := range nums {
		// Get the absolute value of the number as the index
		idx := int(math.Abs(float64(num)))
		// Log the current index and value
		fmt.Println("Current index:", idx, "Current value:", nums[idx])

		// Check if the number at the index is negative
		if nums[idx] < 0 {
			// If it's negative, the index is a duplicate, return it
			fmt.Println("Duplicate found at index:", idx)
			return idx
		}
		// Mark the number at the index as negative
		nums[idx] = -nums[idx]
		// Log the current state of the array
		fmt.Println("Array state:", nums)
	}
	// Return -1 if no duplicate is found
	return -1
}

func main() {
	// Define a slice with duplicate numbers
	nums := []int{3, 1, 3, 4, 2}

	// Call the findDuplicate function and print the result
	fmt.Println("findDuplicate result:", findDuplicate(nums))

	// Reset the slice for the next function call
	nums = []int{3, 1, 3, 4, 2}

	// Call the findDiuplicateMark function and print the result
	fmt.Println("findDiuplicateMark result:", findDiuplicateMark(nums))
}
