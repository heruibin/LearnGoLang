package quicksort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	t.Log(values)
	QuickSort2(values)
	t.Log(values)
}
