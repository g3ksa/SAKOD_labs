package main

import (
	"fmt"
	//heapsort "lab8/HeapSort"
	"math"

	heapsort "lab8/HeapSort"
	insertionsort "lab8/InsertionSort"
	quicksort "lab8/QuickSort"
	radixSort "lab8/RadixSort"
	"math/rand"
	"time"
)

func quickAndInsertionSort() {
	var arr []int
	var randNum int
	var start time.Time
	var end time.Duration

	fmt.Print("Быстрая сортировка: ")
	N := 250000
	for i := 0; i < N; i++ {
		randNum = rand.Intn(N)
		arr = append(arr, randNum)
	}
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = 500000
	for i := 0; i < N; i++ {
		randNum = rand.Intn(N)
		arr = append(arr, randNum)
	}
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = int(math.Pow(10, 6))
	for i := 0; i < N; i++ {
		randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, randNum)
	}
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = int(math.Pow(10, 6))
	for i := 1; i <= N; i++ {
		//randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, i)
	}
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = int(math.Pow(10, 6))
	for i := N; i >= 1; i-- {
		//randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, i)
	}
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	fmt.Print(end)

	fmt.Println()
	fmt.Print("Сортировка вставками: ")
	N = int(2.5 * math.Pow(10, 5))
	for i := 0; i < N; i++ {
		randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, randNum)
	}
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = int(5 * math.Pow(10, 5))
	for i := 0; i < N; i++ {
		randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, randNum)
	}
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = int(math.Pow(10, 6))
	for i := 0; i < N; i++ {
		randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, randNum)
	}
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = int(math.Pow(10, 6))
	for i := 1; i <= N; i++ {
		//randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, i)
	}
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	fmt.Print(end, '\t')

	arr = []int{}
	N = int(math.Pow(10, 6))
	for i := N; i >= 1; i-- {
		//randNum = rand.Intn(int(math.Pow(float64(N), 2)))
		arr = append(arr, i)
	}
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	fmt.Print(end)
}

func task4() [3][]time.Duration {
	var start time.Time
	var end time.Duration
	var N int = 5000
	var arr = make([]int, N)
	var tempArr = make([]int, N)
	var res [3][]time.Duration

	// N = 5000
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	// Radix sort
	start = time.Now()
	radixSort.RadixSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Quick sort
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Heap sort
	start = time.Now()
	heapsort.HeapSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)

	// N = 10000
	N = 10000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	// Radix sort
	start = time.Now()
	radixSort.RadixSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Quick sort
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Heap sort
	start = time.Now()
	heapsort.HeapSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)

	// N = 20000
	N = 20000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	// Radix sort
	start = time.Now()
	radixSort.RadixSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Quick sort
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Heap sort
	start = time.Now()
	heapsort.HeapSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)

	// N =20000(возрастание)
	N = 20000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
		tempArr[i] = i
	}
	// Radix sort
	start = time.Now()
	radixSort.RadixSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Quick sort
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Heap sort
	start = time.Now()
	heapsort.HeapSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)

	// N =20000(убывание)
	N = 20000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i, j := N-1, 0; i >= 0; i-- {
		arr[j] = i
		tempArr[j] = i
		j++
	}
	// Radix sort
	start = time.Now()
	radixSort.RadixSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Quick sort
	start = time.Now()
	quicksort.Quicksort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Heap sort
	start = time.Now()
	heapsort.HeapSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)

	return res
}

func main() {
	// quickAndInsertionSort()
	// fmt.Println()
	// task4()

	arr := []int{1, 4, 6, 2, 9, 1}
	quicksort.Quicksort(arr)
	fmt.Println(arr)

}
