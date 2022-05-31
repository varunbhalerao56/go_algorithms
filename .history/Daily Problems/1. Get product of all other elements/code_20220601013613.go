package main

import "fmt"

/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1] ,the expected output would be [2, 3, 6].

Follow-up: What if you can't use division?
*/

func main() {
	getProductOfAllOtherElements([]int{3, 2, 1})
	getProductOfAllOtherElements([]int{1, 2, 3, 4})
	getProductOfAllOtherElements([]int{1, 2, 3, 4, 5})

	getProductOfAllOtherElementsBad([]int{3, 2, 1})
	getProductOfAllOtherElementsBad([]int{1, 2, 3, 4})
	getProductOfAllOtherElementsBad([]int{1, 2, 3, 4, 5})
}

func getProductOfAllOtherElementsBad(inputArray []int) []int {

	inputLen := len(inputArray)
	outputArray := make([]int, inputLen)

	for i := 0; i < inputLen; i++ {

		temp := 1
		for j := 0; j < inputLen; j++ {
			if j == i {
				continue
			} else {
				temp = temp * inputArray[j]
				outputArray[i] = temp
			}
		}
	}

	fmt.Println("Slow Solution O(n)^2:", outputArray)

	return outputArray
}

func getProductOfAllOtherElements(inputArray []int) []int {

	inputLen := len(inputArray)

	// Make array that calculates products of all items excluding the first index (left to right)
	left_products := make([]int, inputLen)

	// Make array that calculates products of all items excluding the last index (right to left)
	right_product := make([]int, inputLen)
	outputArray := make([]int, inputLen)

	left_products[0] = 1
	right_product[inputLen-1] = 1

	//* Loop to calculate all excluding first index
	for i := 1; i < inputLen; i++ {

		//* [1, (Index Start), 0]
		//* [1, 0, 0] -> [1, 3, 0] -> [1, 3, 6]
		leftOfIndex := i - 1
		left_products[i] = left_products[leftOfIndex] * inputArray[leftOfIndex]
	}

	//* Loop to calculate all excluding last index
	for j := inputLen - 2; j >= 0; j-- {

		//* [0, (Index Start), 1]
		//* [0, 0, 1] -> [0, 1, 1] -> [2, 1, 1]
		rightOfIndex := j + 1
		right_product[j] = right_product[rightOfIndex] * inputArray[rightOfIndex]
	}

	//* Multiply left & right to get answer [1, 3, 6] * [2, 1, 1] = [2, 3, 6]
	for m := 0; m < inputLen; m++ {
		outputArray[m] = left_products[m] * right_product[m]
	}

	fmt.Println("Fast Solution O(n):", outputArray)

	return outputArray
}
