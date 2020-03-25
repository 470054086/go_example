package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 22, 34, 55, 65}
	search := binarySearchLoop(nums, 32)
	fmt.Println(search)
}

/**
二分查找
*/
func binarySearch(nums []int, find int, startIndex, rightIndex int) int {
	if startIndex > rightIndex {
		return -1
	}
	mid := int(math.Ceil((float64(startIndex) + float64(rightIndex)) / 2))
	// 如果当前的数字 小于查找的数字 则像左边查找
	if nums[mid] > find {
		return binarySearch(nums, find, startIndex, mid-1)
	} else if nums[mid] < find {
		// 如果当前的数字 大于查找的数字 则像右边查找
		return binarySearch(nums, find, mid+1, rightIndex)
	} else {
		return mid
	}
}
func binarySearchLoop(nums []int, find int) int {
	start := 0
	end := len(nums)
	for start <= end {
		mid := int(math.Ceil((float64(start) + float64(end)) / 2))
		// 如果中间的位置大于查找的位置 则需要像左寻找
		if nums[mid] > find {
			end = mid-1
		} else if nums[mid] < find {
		// 如果中间的作为小于查找的位置  则需向右寻找
			start = mid+1
		}else{
			return mid
		}
	}
	return -1
}
