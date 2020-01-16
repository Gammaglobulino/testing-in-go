package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestBubbleSortASC(t *testing.T) {
	//init
	elements := []int{3, 4, 1, 0, 8, 9, 6, 4, 5, 3, 2}
	fmt.Println(elements)

	//execution
	BubbleSort(elements)

	//validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
	fmt.Println(elements)
}
func getElements(n int) []int {
	elements := make([]int, n)
	for i := 0; i < n; i++ {
		elements[i] = rand.Intn(n)
	}
	return elements
}
func TestSortASC(t *testing.T) {
	//init
	elements := []int{3, 4, 1, 0, 8, 9, 6, 4, 5, 3, 2}
	fmt.Println(elements)

	//execution
	sort.Ints(elements)

	//validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
	fmt.Println(elements)
}

var elements []int = getElements(100000)

func BenchmarkBubbleSort(b *testing.B) {
	fmt.Println(elements)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
func BenchmarkSort(b *testing.B) {
	fmt.Println(elements)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
