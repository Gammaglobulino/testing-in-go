package services

import (
	"api/utis/sort"
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
