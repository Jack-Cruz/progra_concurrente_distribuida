package main

import (
	"sync"
	"os"
	"fmt"
	"bufio"
	"strconv"
	"time"
)

func quicksort(nums []int, wg *sync.WaitGroup) {
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
	go quicksort(left, &wgLeft)

	wgRight.Add(1)
	go quicksort(right, &wgRight)

	wgLeft.Wait()
	wgRight.Wait()

	copy(nums, append(append(left, pivot), right...))
}

func parallelQuicksort(nums []int) {
	var wg sync.WaitGroup
	wg.Add(1)
	quicksort(nums, &wg)
	wg.Wait()
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

	// Sort the list using parallel quicksort
	parallelQuicksort(list)

	elapsed := time.Since(start)

	// Print the sorted list
	fmt.Println("First 100 elements: ", list[:100])
	fmt.Println("Last 100 elements: ", list[N-100:])
	fmt.Println("List ordered")
	fmt.Println("Time taken: ", elapsed)
	fmt.Println("Time taken in miliseconds: ", elapsed.Milliseconds())
}
