package shellSort

import (
	"math"
)

func ShellSortFirst(arr []int) {
	var j int
	var step int = len(arr) / 2
	for step > 0 {
		for i := 0; i < len(arr)-step; i++ {
			j = i
			for (j >= 0) && (arr[j] > arr[j+step]) {
				var tmp = arr[j]
				arr[j] = arr[j+step]
				arr[j+step] = tmp
				j -= step
			}
		}
		step = step / 2
	}
}

func ShellSortSecond(arr []int) {
	for gap := int(math.Log2(float64(len(arr))) - 1); gap > 0; gap = int(math.Log2(float64(gap)) - 1) {
		for i := gap; i < len(arr); i++ {
			x := arr[i]
			j := i
			for ; j >= gap && arr[j-gap] > x; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = x
		}
	}
}

func ShellSortThird(arr []int) {
	for gap := int(math.Round(log(3, float64(len(arr))) - 1)); gap > 0; gap = int(math.Round(log(3, float64(gap)) - 1)) {
		for i := gap; i < len(arr); i++ {
			x := arr[i]
			j := i
			for ; j >= gap && arr[j-gap] > x; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = x
		}
	}
}

func log(x int, y float64) float64 {
	return math.Log(y) / math.Log(float64(x))
}
