// https://leetcode.com/problems/two-sum/description/
package main

func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int, len(nums))

	for i, j := range nums {
		value, ok := mapNums[target-j]

		if ok {
			return []int{i, value}
		}

		mapNums[j] = i
	}
	return nil
}

func main() {
}
