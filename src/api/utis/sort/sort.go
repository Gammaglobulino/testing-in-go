package sort

import "sort"

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
