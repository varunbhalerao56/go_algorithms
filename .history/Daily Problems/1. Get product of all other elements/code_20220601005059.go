package main

import "fmt"

/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1] ,the expected output would be [2, 3, 6].

Follow-up: What if you can't use division?
*/

func main() {
	//	getProductOfAllOtherElements([]int{3, 2, 1})
	getProductOfAllOtherElements([]int{1, 2, 3, 4})
}

func getProductOfAllOtherElements(inputArray []int) []int {

	inputLen := len(inputArray)

	left_products := make([]int, inputLen)
	right_product := make([]int, inputLen)

	left_products[0] = 1
	right_product[inputLen-1] = 1

	outputArray := []int{}

	for i := 1; i < inputLen; i++ {
		left_products[i] = left_products[i-1] * inputArray[i]
	}

	for j := inputLen - 2; j > 0; j-- {

		right_product[j] = right_product[j] * inputArray[j]
	}

	fmt.Println(left_products)
	fmt.Println(right_product)

	return outputArray

}
