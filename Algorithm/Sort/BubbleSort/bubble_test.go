package bubble

import (
	"testing"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	arrays := []int{2, 3, 1, 5, 6, 4, 9}
	arrays = BubbleSort(arrays)
	for i := 0; i < len(arrays); i++ {
		t.Log(arrays[i])
	}
	fmt.Println(arrays)
}
