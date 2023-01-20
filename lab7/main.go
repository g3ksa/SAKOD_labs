package main

import (
	"fmt"
	binaryinsertionsort "lab7/BinaryInsertionSort"
	bubblesort "lab7/BubbleSort"
	cocktailsort "lab7/CocktailSort"
	insertionsort "lab7/InsertionSort"
	selectsort "lab7/SelectSort"
	shellSort "lab7/ShellSort"
	task3 "lab7/Task3"
	"math/rand"
	"time"
)

func task2() [5][]time.Duration {
	var start time.Time
	var end time.Duration
	var N int = 5000
	var arr = make([]int, N)
	var tempArr = make([]int, N)
	var res [5][]time.Duration

	// N = 5000
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	// Bubble sort
	start = time.Now()
	bubblesort.BubbleSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	cocktailsort.CocktailSort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	selectsort.SelectSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	binaryinsertionsort.BinaryInsertionSort(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N = 10000
	N = 10000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	// Bubble sort
	start = time.Now()
	bubblesort.BubbleSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	cocktailsort.CocktailSort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	selectsort.SelectSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	binaryinsertionsort.BinaryInsertionSort(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N = 20000
	N = 20000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	// Bubble sort
	start = time.Now()
	bubblesort.BubbleSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	cocktailsort.CocktailSort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	selectsort.SelectSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	binaryinsertionsort.BinaryInsertionSort(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N =20000(возрастание)
	N = 20000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
		tempArr[i] = i
	}
	// Bubble sort
	start = time.Now()
	bubblesort.BubbleSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	cocktailsort.CocktailSort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	selectsort.SelectSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	binaryinsertionsort.BinaryInsertionSort(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	// N =20000(убывание)
	N = 20000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i, j := N-1, 0; i >= 0; i-- {
		arr[j] = i
		tempArr[j] = i
		j++
	}
	// Bubble sort
	start = time.Now()
	bubblesort.BubbleSort(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	cocktailsort.CocktailSort(arr)
	end = time.Since(start)
	res[1] = append(res[1], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Cocktail sort
	start = time.Now()
	selectsort.SelectSort(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Insertion sort
	start = time.Now()
	insertionsort.InsertionSort(arr)
	end = time.Since(start)
	res[3] = append(res[3], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	// Binary insertion sort
	start = time.Now()
	binaryinsertionsort.BinaryInsertionSort(arr)
	end = time.Since(start)
	res[4] = append(res[4], end)

	return res
}

func task4() [3][]time.Duration {
	var start time.Time
	var end time.Duration
	var N int = 25000
	var arr = make([]int, N)
	var tempArr = make([]int, N)
	var res [3][]time.Duration

	// N = 25000
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	//First
	start = time.Now()
	shellSort.ShellSortFirst(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	//Second
	start = time.Now()
	shellSort.ShellSortSecond(arr)
	end = time.Since(start)
	res[1] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	//Third
	start = time.Now()
	shellSort.ShellSortThird(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}

	// N = 50000
	N = 50000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	//First
	start = time.Now()
	shellSort.ShellSortFirst(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	//Second
	start = time.Now()
	shellSort.ShellSortSecond(arr)
	end = time.Since(start)
	res[1] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	//Third
	start = time.Now()
	shellSort.ShellSortThird(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}

	// N = 100000
	N = 100000
	arr = make([]int, N)
	tempArr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		tempArr[i] = arr[i]
	}
	//First
	start = time.Now()
	shellSort.ShellSortFirst(arr)
	end = time.Since(start)
	res[0] = append(res[0], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	//Second
	start = time.Now()
	shellSort.ShellSortSecond(arr)
	end = time.Since(start)
	res[1] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	//Third
	start = time.Now()
	shellSort.ShellSortThird(arr)
	end = time.Since(start)
	res[2] = append(res[2], end)
	for i := 0; i < N; i++ {
		arr[i] = tempArr[i]
	}
	return res
}

func main() {
	fmt.Println("Task 2")
	arr := task2()
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("Task 3")
	arr3 := task3.Task3()
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("Task 4")
	arr2 := task4()
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
}
