package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"strconv"
	"time"
)

// Merge two sorted slices into a single sorted slice
func merge(left, right []int) []int {
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

// Parallel merge sort function
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
    return merge(leftSorted, rightSorted)
}

func main() {
	N := 10000

	// Open the input file
	file, err := os.Open("./01_one_million_random_numbers.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to store the numbers
	var list[] int

	// Read each line and convert it to an integer, and then add it to the slice
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		list = append(list, num)
	}
	// Setting the time
	start := time.Now()

	// Sprt the list using parallel merge sort
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		list = parallelMergeSort(list[:N], &wg)
	}()
	wg.Wait()

	elapsed := time.Since(start)

	fmt.Println("First 100 elements: ", list[:100])
	fmt.Println("Last 100 elements: ", list[N-100:])
	fmt.Println("List ordered")
	fmt.Println("Time taken: ", elapsed)
	fmt.Println("Time taken in miliseconds: ", elapsed.Milliseconds())
}
