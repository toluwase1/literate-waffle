package main

import (
	"fmt"
	"sort"
)

/*
1. Find the missing number in the array
You are given an array of positive numbers from 1 to n, such that
all numbers from 1 to n are present except one number x. You
have to find x. The input array is not sorted. Look at the below
array and give it a try before checking the solution.
*/

func missingNumber(arr []int) int {
	//sort
	//check the diff b/2 i and i+1, if not 1 return arr[i]+arr[i+1]/2
	sort.Ints(arr)
	var missingNo int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]-arr[i+1] == -1 {
			continue
		} else {
			missingNo = (arr[i] + arr[i+1]) / 2
		}
	}
	return missingNo
}

/*
Given an array of integers and a data, determine if there are any two integers in the array
whose sum is equal to the given data. Return true if the sum exists and return false if it
does not. Consider this array and the target sums:
*/
func findSumOfTwo(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return true
			}
		}
	}
	return false
}

func missingNumber2(arr []int) int {
	val := []int{}

	fmt.Println(arr)
	sort.Ints(arr)
	sum1 := 0

	for _, j := range arr {
		sum1 += j
	}
	fmt.Println(sum1)
	i := 1
	//v := 0
	val = append(val, arr...)
	fmt.Println(val)
	sum2 := val[0]
	for i <= len(val)+1 {
		fmt.Println("hii", i)
		sum2 += sum2 + i
		i++
	}
	fmt.Println(sum2, sum1)
	return sum2 - sum1
	//sum + len(arr)+1
}

//func main() {
//	arr := []int{3, 1, 2, 8, 4, 5, 6}
//	fmt.Println(missingNumber(arr))
//	arrr := []int{5, 7, 1, 2, 8, 4, 3}
//	fmt.Println(findSumOfTwo(arrr, 20))
//}
