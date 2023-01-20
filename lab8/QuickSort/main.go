package QuickSort

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(arr []int) int {
	var pivot = arr[0]
	var i = 0
	var j = len(arr) - 1

	for i < j {
		for arr[i] <= pivot && i < len(arr)-1 {
			i++
		}
		for arr[j] > pivot && j > 0 {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[0] = arr[j]
	arr[j] = pivot

	return j
}
