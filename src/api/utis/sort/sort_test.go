package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubbleSortASC(t *testing.T) {
	//init
	elements := GetElements(10)

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

func TestSortASC(t *testing.T) {
	//init
	elements := GetElements(10)

	//execution
	sort.Ints(elements)

	//validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}

}

var elements []int = GetElements(10000)

func BenchmarkBubbleSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
func BenchmarkSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
