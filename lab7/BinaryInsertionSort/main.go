package binaryinsertionsort

func binarySearch(arr []int, item, low, high int) int {
	if high <= low {
		if item > arr[low] {
			return low + 1
		}
		return low
	}

	var mid = (low + high) / 2

	if item == arr[mid] {
		return mid + 1
	}

	if item > arr[mid] {
		return binarySearch(arr, item, mid+1, high)
	}
	return binarySearch(arr, item, low, mid-1)
}

func BinaryInsertionSort(arr []int) {
	var loc, j, selected int

	for i := 1; i < len(arr); i++ {
		j = i - 1
		selected = arr[i]

		loc = binarySearch(arr, selected, 0, j)

		for j >= loc {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = selected
	}
}
