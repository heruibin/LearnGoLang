package bubble

func BubbleSort(arrays []int) []int {
	for i := 0; i < len(arrays); i++ {
		flag := true
		for j := 0; j < len(arrays)-i-1; j++ {
			if arrays[j] > arrays[j+1] {
				arrays[j], arrays[j+1] = arrays[j+1], arrays[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return arrays
}
