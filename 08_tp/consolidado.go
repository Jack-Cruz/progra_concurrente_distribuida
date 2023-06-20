package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Generador de N números enteros aleatorios
func generate_numbers(N int) []int {
	// Create a slice of 1,000,000 numbers
    list := make([]int, N)
	
	for i := 0; i < N; i++ {
		list[i] = rand.Intn(N)
	}

	return list
}

// Merge Sort tradicional
func merge(arr []int, left []int, right []int) {
    i := 0 // index of left sub-array
    j := 0 // index of right sub-array
    k := 0 // index of merged array

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    // Copy the remaining elements of left sub-array
    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }

    // Copy the remaining elements of right sub-array
    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
}

func MergeSort(arr []int) {
    if len(arr) < 2 {
        return
    }

    mid := len(arr) / 2
    left := make([]int, mid)
    right := make([]int, len(arr)-mid)

    for i := 0; i < mid; i++ {
        left[i] = arr[i]
    }

    for i := mid; i < len(arr); i++ {
        right[i-mid] = arr[i]
    }

    MergeSort(left)
    MergeSort(right)
    merge(arr, left, right)
}

// Merge Sort Parallelo
func parallelmerge(left, right []int) []int {
    result := make([]int, len(left)+len(right))
    i := 0
    j := 0
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result[i+j] = left[i]
            i++
        } else {
            result[i+j] = right[j]
            j++
        }
    }
    for i < len(left) {
        result[i+j] = left[i]
        i++
    }
    for j < len(right) {
        result[i+j] = right[j]
        j++
    }
    return result
}
func parallelMergeSort(nums []int, wg *sync.WaitGroup) []int {
    defer wg.Done()

    // If the slice has only one element, return it
    if len(nums) == 1 {
        return nums
    }

    // Split the slice into two halves
    mid := len(nums) / 2
    var left, right []int
    left = nums[:mid]
    right = nums[mid:]

    // Use a wait group to synchronize the two goroutines that sort the halves
    var leftSorted, rightSorted []int
    var wg2 sync.WaitGroup
    wg2.Add(2)
    go func() {
        leftSorted = parallelMergeSort(left, &wg2)
    }()
    go func() {
        rightSorted = parallelMergeSort(right, &wg2)
    }()
    wg2.Wait()

    // Merge the two sorted halves
    return parallelmerge(leftSorted, rightSorted)
}

// Quick Sort tradicional
func quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[0]

	var left, right []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	return append(append(left, pivot), right...)
}

// Quick sort paralelo
func quicksortparallel(nums []int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(nums) <= 1 {
		return
	}

	pivot := nums[0]

	var left, right []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	var wgLeft, wgRight sync.WaitGroup
	wgLeft.Add(1)
	go quicksortparallel(left, &wgLeft)

	wgRight.Add(1)
	go quicksortparallel(right, &wgRight)

	wgLeft.Wait()
	wgRight.Wait()

	copy(nums, append(append(left, pivot), right...))
}

func parallelQuicksort(nums []int) {
	var wg sync.WaitGroup
	wg.Add(1)
	quicksortparallel(nums, &wg)
	wg.Wait()
}

func main() {
	N := 1000
	list := make([]int, N)

	var tiempo_merge_tradicional[] int64
	var tiempo_merge_parallel[] int64
	var tiempo_quick_tradicional[] int64
	var tiempo_quick_parallel[] int64

	for i := 0; i < 50; i++ {
		fmt.Println("Iteración: ", i)
		list = generate_numbers(N)
		
		for i := 0; i < 10; i++ {
			fmt.Println(list[i])
		}
		fmt.Println()

		// Merge sort parallel
		start := time.Now()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			parallelMergeSort(list[:N], &wg)
		}()
		wg.Wait()
		elapsed := time.Since(start)
		tiempo_merge_parallel = append(tiempo_merge_parallel, elapsed.Milliseconds())
		
		for i := 0; i < 10; i++ {
			fmt.Println(list[i])
		}
		fmt.Println()

		// Merge sort tradicional
		start = time.Now()
		MergeSort(list)
		elapsed = time.Since(start)
		tiempo_merge_tradicional = append(tiempo_merge_tradicional, elapsed.Milliseconds())


		list = generate_numbers(N)
		
		for i := 0; i < 10; i++ {
			fmt.Println(list[i])
		}
		fmt.Println()

		// Quick sort tradicional
		start = time.Now()
		quicksort(list)
		elapsed = time.Since(start)
		tiempo_quick_tradicional = append(tiempo_quick_tradicional, elapsed.Milliseconds())

		for i := 0; i < 10; i++ {
			fmt.Println(list[i])
		}

		// Quick parallel
		start = time.Now()
		parallelQuicksort(list)
		elapsed = time.Since(start)
		tiempo_quick_parallel = append(tiempo_quick_parallel, elapsed.Milliseconds())
	}
}