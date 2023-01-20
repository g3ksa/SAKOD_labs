package cocktailsort

func CocktailSort(arr []int) {
	left := 0
	right := len(arr) - 1
	count := 0

	for left < right {
		for i := left; i < right; i++ {
			count++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--

		for i := right; i > left; i-- {
			count++
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		left++
	}
}
