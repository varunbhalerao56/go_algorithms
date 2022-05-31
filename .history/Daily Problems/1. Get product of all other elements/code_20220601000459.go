package main

import "fmt"

/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1] ,the expected output would be [2, 3, 6].

Follow-up: What if you can't use division?
*/

func main() {
	getProductOfAllOtherElements([]int{1, 2, 3, 4, 5})
}

func getProductOfAllOtherElements(inputArray []int) []int {
	outputArray := []int{}
	total := inputArray[0]

	for i := 1; i < len(inputArray); i++ {

		total += (total * inputArray[i])
		fmt.Println(total)
	}

	for j := len(inputArray) - 1; j >= 0; j-- {
		outputArray = append(outputArray, total/inputArray[j])
	}

	fmt.Println(outputArray)

	return outputArray

}
