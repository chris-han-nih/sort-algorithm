package bubble

func Sort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
