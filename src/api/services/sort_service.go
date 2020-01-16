package services

import "api/utis/sort"

func Sort(elements []int) {
	if len(elements) >= 10000 {
		sort.Sort(elements)
		return
	}
	sort.BubbleSort(elements)

}
