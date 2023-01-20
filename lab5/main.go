package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func RecBinarySearch(array []int, M int, watch *[]interface{}) {
	start := time.Now()
	for i := 0; i < M; i++ {
		Svalue := rand.Intn(1000)
		BinarySearch(array, Svalue, 0, len(array)-1)
	}
	end := time.Since(start)
	*watch = append(*watch, end)
}

func RecInterpolationSearch(array []int, M int, watch *[]interface{}) {
	start := time.Now()
	for i := 0; i < M; i++ {
		Svalue := rand.Intn(1000)
		InterpolationSearch(array, Svalue)
	}
	end := time.Since(start)
	*watch = append(*watch, end)
}

func RecConsistentSearch(array []int, M int, watch *[]interface{}) {
	start := time.Now()
	for i := 0; i < M; i++ {
		Svalue := rand.Intn(1000)
		ConsistentSearch(array, Svalue)
	}
	end := time.Since(start)
	*watch = append(*watch, end)
}

func RecConsistentSearchWithBarrier(array []int, M int, watch *[]interface{}) {
	start := time.Now()
	for i := 0; i < M; i++ {
		Svalue := rand.Intn(1000)
		ConsistentSearchWithBarrier(array, Svalue)
	}
	end := time.Since(start)
	*watch = append(*watch, end)
}

func ConsistentSearchWithBarrier(array []int, SValue int) int {
	last := array[len(array)-1]
	array[len(array)-1] = SValue
	i := 0
	for array[i] != SValue {
		i++
	}

	if i == len(array)-1 && last != SValue {
		return -1
	} else {
		return i
	}
}

func ConsistentSearch(array []int, Svalue int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == Svalue {
			return i
		}
	}
	return -1
}

func InterpolationSearch(array []int, Svalue int) int {
	var mid int
	var low int = 0
	high := len(array) - 1

	for array[low] < Svalue && array[high] > Svalue {
		if array[high] == array[low] {
			break
		}
		mid = low + ((Svalue-array[low])*(high-low))/(array[high]-array[low])

		if array[mid] < Svalue {
			low = mid + 1
		} else if array[mid] > Svalue {
			high = mid - 1
		} else {
			return mid
		}
	}

	if array[low] == Svalue {
		return low
	}
	if array[high] == Svalue {
		return high
	}

	return -1
}

func BinarySearch(array []int, Svalue, first, last int) int {
	if first > last {
		return -1
	}

	middle := (first + last) / 2
	middleValue := array[middle]

	if middleValue == Svalue {
		return middle
	} else {
		if middleValue > Svalue {
			return BinarySearch(array, Svalue, first, middle-1)
		} else {
			return BinarySearch(array, Svalue, middle+1, last)
		}
	}
}

func main() {
	var watch []interface{}

	var array1 = make([]int, 1000)
	var array2 = make([]int, 2000)
	var array3 = make([]int, 4000)
	var array4 = make([]int, 8000)
	var array5 = make([]int, 16000)

	for i := 0; i < len(array5); i++ {
		if i < 1000 {
			array1[i] = rand.Intn(1000)
		}
		if i < 2000 {
			array2[i] = rand.Intn(1000)
		}
		if i < 4000 {
			array3[i] = rand.Intn(1000)
		}
		if i < 8000 {
			array4[i] = rand.Intn(1000)
		}
		array5[i] = rand.Intn(1000)
	}
	fmt.Println("Пункт А:")
	fmt.Println("  M   Алгоритм                N=1000          N=2000            N=4000            N=8000            N=16000")
	RecConsistentSearch(array1, 5000, &watch)
	RecConsistentSearch(array2, 5000, &watch)
	RecConsistentSearch(array3, 5000, &watch)
	RecConsistentSearch(array4, 5000, &watch)
	RecConsistentSearch(array5, 5000, &watch)
	fmt.Printf("5000  Последовательный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecConsistentSearchWithBarrier(array1, 5000, &watch)
	RecConsistentSearchWithBarrier(array2, 5000, &watch)
	RecConsistentSearchWithBarrier(array3, 5000, &watch)
	RecConsistentSearchWithBarrier(array4, 5000, &watch)
	RecConsistentSearchWithBarrier(array5, 5000, &watch)
	fmt.Printf("5000  Последовательный с барьером     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecConsistentSearch(array1, 10000, &watch)
	RecConsistentSearch(array2, 10000, &watch)
	RecConsistentSearch(array3, 10000, &watch)
	RecConsistentSearch(array4, 10000, &watch)
	RecConsistentSearch(array5, 10000, &watch)
	fmt.Printf("10000  Последовательный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecConsistentSearchWithBarrier(array1, 10000, &watch)
	RecConsistentSearchWithBarrier(array2, 10000, &watch)
	RecConsistentSearchWithBarrier(array3, 10000, &watch)
	RecConsistentSearchWithBarrier(array4, 10000, &watch)
	RecConsistentSearchWithBarrier(array5, 10000, &watch)
	fmt.Printf("10000  Последовательный с барьером     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecConsistentSearch(array1, 20000, &watch)
	RecConsistentSearch(array2, 20000, &watch)
	RecConsistentSearch(array3, 20000, &watch)
	RecConsistentSearch(array4, 20000, &watch)
	RecConsistentSearch(array5, 20000, &watch)
	fmt.Printf("20000  Последовательный    %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecConsistentSearchWithBarrier(array1, 20000, &watch)
	RecConsistentSearchWithBarrier(array2, 20000, &watch)
	RecConsistentSearchWithBarrier(array3, 20000, &watch)
	RecConsistentSearchWithBarrier(array4, 20000, &watch)
	RecConsistentSearchWithBarrier(array5, 20000, &watch)
	fmt.Printf("20000  Последовательный с барьером     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()
	fmt.Println()
	fmt.Println()

	sort.Ints(array1)
	sort.Ints(array2)
	sort.Ints(array3)
	sort.Ints(array4)
	sort.Ints(array5)

	fmt.Println("\nПункт Б:")
	fmt.Println("  M   Алгоритм                N=1000          N=2000            N=4000            N=8000            N=16000")
	RecBinarySearch(array1, 5000, &watch)
	RecBinarySearch(array2, 5000, &watch)
	RecBinarySearch(array3, 5000, &watch)
	RecBinarySearch(array4, 5000, &watch)
	RecBinarySearch(array5, 5000, &watch)
	fmt.Printf("5000  Бинарный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecInterpolationSearch(array1, 5000, &watch)
	RecInterpolationSearch(array2, 5000, &watch)
	RecInterpolationSearch(array3, 5000, &watch)
	RecInterpolationSearch(array4, 5000, &watch)
	RecInterpolationSearch(array5, 5000, &watch)
	fmt.Printf("5000  Интерполяционный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecBinarySearch(array1, 10000, &watch)
	RecBinarySearch(array2, 10000, &watch)
	RecBinarySearch(array3, 10000, &watch)
	RecBinarySearch(array4, 10000, &watch)
	RecBinarySearch(array5, 10000, &watch)
	fmt.Printf("10000  Бинарный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecInterpolationSearch(array1, 10000, &watch)
	RecInterpolationSearch(array2, 10000, &watch)
	RecInterpolationSearch(array3, 10000, &watch)
	RecInterpolationSearch(array4, 10000, &watch)
	RecInterpolationSearch(array5, 10000, &watch)
	fmt.Printf("10000  Интерполяционный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecBinarySearch(array1, 20000, &watch)
	RecBinarySearch(array2, 20000, &watch)
	RecBinarySearch(array3, 20000, &watch)
	RecBinarySearch(array4, 20000, &watch)
	RecBinarySearch(array5, 20000, &watch)
	fmt.Printf("20000  Бинарный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])

	watch = watch[:0]
	fmt.Println()

	RecInterpolationSearch(array1, 20000, &watch)
	RecInterpolationSearch(array2, 20000, &watch)
	RecInterpolationSearch(array3, 20000, &watch)
	RecInterpolationSearch(array4, 20000, &watch)
	RecInterpolationSearch(array5, 20000, &watch)
	fmt.Printf("20000  Интерполяционный     %v       %v       %v       %v         %v", watch[0], watch[1], watch[2], watch[3], watch[4])
}
