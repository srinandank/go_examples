package main

import "fmt"

func twoSum(target int, nums []int) [2]int {
	var my_map = make(map[int]int)
	for idx, num_value := range nums {
		complement := target - num_value
		value, ok := my_map[complement]
		if ok {
			return [2]int{value, idx}
		} else {
			my_map[num_value] = idx
		}
	}
	return [2]int{-1, -1}
}

func main() {
	var target int = 10
	nums := [...]int{1, 3, 4, 5, 7, 9}
	fmt.Printf("Target is: %d, Nums are: %d\n", target, nums)
	elements := twoSum(target, nums[:])
	if elements[0] == -1 {
		fmt.Println("No successful values found to create target! :(")
	} else {
		fmt.Printf("Target sum is satisfied by elements %d and %d of values %d and %d respectively\n", elements[0]+1, elements[1]+1, nums[elements[0]], nums[elements[1]])
	}
}
