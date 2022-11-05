package templates

import (
	"fmt"
	"sort"
)

// UnFlexMethod U cant use this template if u need to sort slice where each element is a custom types
// Or if you want to find the index of the first number that is 7 or higher rather than just the first index of 7
func UnFlexMethod() {
	numbers := []int{1, 11, -5, 7, 2, 0, 12}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)
	index := sort.SearchInts(numbers, 7)
	fmt.Println("7 is at index:", index)
}

// BinarySearchingSortSearch Using the sort.Search() function,
// and passing in a closure that can be used to determine
// if the number at a specific index meets your criteria.
func BinarySearchingSortSearch() {
	numbers := []int{1, 11, -5, 8, 2, 0, 12}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)

	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= 7
	})
	fmt.Println("The first number >= 7 is at index:", index)
	fmt.Println("The first number >= 7 is:", numbers[index])
}
