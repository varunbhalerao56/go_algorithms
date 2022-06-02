package main

import (
	"fmt"
	"math"
)

//? Given an array of integers that are out of order.
//? Determine the bounds of the smallest window that must be sorted in order for the entire array to be sorted.

//? For example, given [3 , 7 , 5 , 6 , 9] , you should return ( 1 , 3 ) .
//? Also solves https://leetcode.com/problems/shortest-unsorted-continuous-subarray/submissions/

func main() {

	//! Very Bad and unreliable, only works on some inputs (My own solution...)
	locateSmallestWindowFail([]int{3, 7, 5, 6, 9})
	locateSmallestWindowFail([]int{5, 4, 3, 2, 1})
	locateSmallestWindowFail([]int{1, 2, 3, 4, 5})
	locateSmallestWindowFail([]int{2, 6, 4, 8, 10, 9, 15})
	locateSmallestWindowFail([]int{1, 2, 3, 3, 3})
	locateSmallestWindowFail([]int{1, 3, 2, 2, 2})
	locateSmallestWindowFail([]int{10, 12, 20, 30, 25, 40, 32, 31, 35, 50, 60})

	//! Solution that clearly did not work to easily find for both questions
	//! https://leetcode.com/problems/shortest-unsorted-continuous-subarray/discuss/1787702/Go-Multiple-Solutions-Clean-Code-(18ms-100)

	//* Inspiration from online, was hard to find a simple solutions, some solutions did not
	//* even work. Modified the solution to make it faster and fit the code book and leetcode
	locateSmallestWindow([]int{3, 7, 5, 6, 9})
	locateSmallestWindow([]int{5, 4, 3, 2, 1})
	locateSmallestWindow([]int{1, 2, 3, 4, 5})
	locateSmallestWindow([]int{2, 6, 4, 8, 10, 9, 15})
	locateSmallestWindow([]int{1, 2, 3, 3, 3})
	locateSmallestWindow([]int{1, 3, 2, 2, 2})
	locateSmallestWindow([]int{10, 12, 20, 30, 25, 40, 32, 31, 35, 50, 60})
}

func locateSmallestWindowFail(nums []int) []int {

	outputArray := []int{0, 0}
	length := len(nums)

	max := math.MinInt16
	for i := 0; i < length-1; i++ {

		isGreater := nums[i] > nums[i+1]
		isEqual := nums[i] == nums[i+1]
		isMax := nums[i] >= max

		if isEqual {
			if isMax {
				outputArray[1] = i + 1
			} else {
				continue
			}

		} else if !isGreater && !isMax {
			outputArray[1] = i
		} else if isGreater {
			outputArray[1] = i + 1

			if isMax {
				max = nums[i]
			}

		}

	}

	max = math.MinInt16
	for j := len(nums) - 2; j > 0; j-- {

		isGreater := nums[j] > nums[j+1]
		isSmaller := nums[j] < nums[j-1]
		isMax := nums[j] > max

		if !isGreater && !isSmaller && !isMax {
			outputArray[0] = j
		} else if isSmaller {
			outputArray[0] = j - 1
			continue
		} else if isGreater {
			outputArray[0] = j
			continue
		}

	}

	count := 0
	for m := outputArray[0]; m <= outputArray[1]; m++ {

		if outputArray[1] != 0 {
			count++
		}

	}

	fmt.Println("Array:", outputArray, " Len:", count)
	return outputArray
}

func locateSmallestWindow(nums []int) []int {

	min := math.MaxInt
	max := math.MinInt

	outputArray := []int{0, 0}

	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}

		if max > nums[i] {
			outputArray[1] = i
		}

	}

	for i := len(nums) - 1; i >= 0; i-- {
		if min > nums[i] {
			min = nums[i]
		}
		if min < nums[i] {
			outputArray[0] = i
		}
	}

	count := 0

	for m := outputArray[0]; m <= outputArray[1]; m++ {
		if outputArray[1] != 0 {
			count++
		}
	}

	fmt.Println(outputArray, " Count:", count)
	return outputArray
}
