package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// merge function merges two sorted sub-arrays
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

// MergeSort function sorts the array using merge sort algorithm
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

func main() {
	N := 1000000
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
	
    MergeSort(list)
	
	elapsed := time.Since(start)
	
	fmt.Println("First 100 elements: ", list[:100])
	fmt.Println("Last 100 elements: ", list[N-100:])
	fmt.Println("List ordered")
	fmt.Println("Time taken: ", elapsed)
	fmt.Println("Time taken in miliseconds: ", elapsed.Milliseconds())
}
