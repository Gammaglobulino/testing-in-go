package services

import (
	"api/utis/sort"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	elements := sort.GetElements(10)
	Sort(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
	fmt.Println(elements)
}

func TestSortMoreThan10000(t *testing.T) {
	elements := sort.GetElements(10001)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 10000 {
		t.Error("last element should be 9")
	}

}
