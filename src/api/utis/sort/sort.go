package sort

import (
	"math/rand"
	"sort"
)

//using Bubble Sort homemade
func BubbleSort(elements []int) {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}

		}
	}
}

// using Int Sort from Sort package
func Sort(elements []int) {
	sort.Ints(elements)

}
func GetElements(n int) []int {
	elements := make([]int, n)
	for i := 0; i < n-1; i++ {
		elements[i] = rand.Intn(n)
	}
	elements[0] = 0
	elements[n-1] = n - 1

	return elements
}
