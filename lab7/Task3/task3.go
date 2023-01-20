package task3

import (
	"math/rand"
	"time"
)

type Item struct {
	Number      int
	Title       string
	Description string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Task3() [5][]time.Duration {
	var start time.Time
	var end time.Duration
	var N int = 5000
	var arr = make([]Item, N)
	var tempArr = make([]Item, N)
	var res [5][]time.Duration

	for i := 0; i < N; i++ {
		arr[i].Number = rand.Intn(N)
		arr[i].Title = RandStringRunes(5)
		arr[i].Description = RandStringRunes(10)
		tempArr[i] = arr[i]
	}
	// Bubble sort
	start = time.Now()
	BubbleSortCustomStruct(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	CocktailSortCustomObject(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Selection sort
	start = time.Now()
	SelectSortCustomObject(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	InsertionSortCustomObject(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	BinaryInsertionSortCustomObject(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N = 10000
	N = 10000
	arr = make([]Item, N)
	tempArr = make([]Item, N)
	for i := 0; i < N; i++ {
		arr[i].Number = rand.Intn(N)
		arr[i].Title = RandStringRunes(5)
		arr[i].Description = RandStringRunes(10)
		tempArr[i] = arr[i]
	}
	// Bubble sort
	start = time.Now()
	BubbleSortCustomStruct(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	CocktailSortCustomObject(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	SelectSortCustomObject(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	InsertionSortCustomObject(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	BinaryInsertionSortCustomObject(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N = 20000
	N = 20000
	arr = make([]Item, N)
	tempArr = make([]Item, N)
	for i := 0; i < N; i++ {
		arr[i].Number = rand.Intn(N)
		arr[i].Title = RandStringRunes(5)
		arr[i].Description = RandStringRunes(10)
		tempArr[i] = arr[i]
	}
	// Bubble sort
	start = time.Now()
	BubbleSortCustomStruct(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	CocktailSortCustomObject(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	SelectSortCustomObject(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	InsertionSortCustomObject(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	BinaryInsertionSortCustomObject(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N = 20000(по возрастанию)
	N = 20000
	arr = make([]Item, N)
	tempArr = make([]Item, N)
	for i := 0; i < N; i++ {
		arr[i].Number = i
		arr[i].Title = RandStringRunes(5)
		arr[i].Description = RandStringRunes(10)
		tempArr[i] = arr[i]
	}
	// Bubble sort
	start = time.Now()
	BubbleSortCustomStruct(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	CocktailSortCustomObject(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	SelectSortCustomObject(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	InsertionSortCustomObject(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	BinaryInsertionSortCustomObject(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N = 20000(по убыванию)
	N = 20000
	arr = make([]Item, N)
	tempArr = make([]Item, N)
	for i, j := N-1, 0; i >= 0; i-- {
		arr[i].Number = i
		arr[i].Title = RandStringRunes(5)
		arr[i].Description = RandStringRunes(10)
		tempArr[i] = arr[i]
		j++
	}
	// Bubble sort
	start = time.Now()
	BubbleSortCustomStruct(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	CocktailSortCustomObject(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	SelectSortCustomObject(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	InsertionSortCustomObject(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	BinaryInsertionSortCustomObject(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	return res
}

func BubbleSortCustomStruct(arr []Item) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i].Number > arr[j].Number {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func CocktailSortCustomObject(arr []Item) {
	left := 0
	right := len(arr) - 1
	count := 0

	for left < right {
		for i := left; i < right; i++ {
			count++
			if arr[i].Number > arr[i+1].Number {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--

		for i := right; i > left; i-- {
			count++
			if arr[i-1].Number > arr[i].Number {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		left++
	}
}

func SelectSortCustomObject(arr []Item) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min].Number > arr[j].Number {
				min = j
			}
		}

		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

func InsertionSortCustomObject(ar []Item) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1].Number > x.Number; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}

func binarySearchCustomObject(arr []Item, item Item, low, high int) int {
	if high <= low {
		if item.Number > arr[low].Number {
			return low + 1
		}
		return low
	}

	var mid = (low + high) / 2

	if item == arr[mid] {
		return mid + 1
	}

	if item.Number > arr[mid].Number {
		return binarySearchCustomObject(arr, item, mid+1, high)
	}
	return binarySearchCustomObject(arr, item, low, mid-1)
}

func BinaryInsertionSortCustomObject(arr []Item) {
	var loc, j int
	var selected Item

	for i := 1; i < len(arr); i++ {
		j = i - 1
		selected = arr[i]

		loc = binarySearchCustomObject(arr, selected, 0, j)

		for j >= loc {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = selected
	}
}
