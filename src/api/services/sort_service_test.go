package services

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	elements := []int{9, 5, 6, 4, 3, 2, 4, 0}
	fmt.Println(elements)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
	fmt.Println(elements)
}
