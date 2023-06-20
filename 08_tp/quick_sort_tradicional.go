package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"time"
)

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
	
	// Sort the list using quicksort
	list = quicksort(list)

	elapsed := time.Since(start)

	// Print the sorted list
	fmt.Println("First 100 elements: ", list[:100])
	fmt.Println("Last 100 elements: ", list[N-100:])
	fmt.Println("List ordered")
	fmt.Println("List ordered")
	fmt.Println("Time taken: ", elapsed)
	fmt.Println("Time taken in miliseconds: ", elapsed.Milliseconds())
}
