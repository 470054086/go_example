package main

import "fmt"

func main()  {
	nums := []int{4,5,1,2}
	quickSort(nums,0,len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int,startIndex,lastIndex int)  {
	// 递归结束的条件
	if startIndex>= lastIndex {
		return
	}
	mid := mid2(nums,startIndex,lastIndex)
	quickSort(nums,startIndex,mid-1)
	quickSort(nums,mid+1,lastIndex)

}

func mid(nums []int,startIndex,lastIndex int) int  {
	// 假设第一个为基准点 左边的指针为startIndex 右边的指针为lastIndex
	base := nums[startIndex]
	left := startIndex
	right :=lastIndex
	//1. 先从右指针开始移动 如果当前的数比left小于的话 就把指针左移
	// 直到遇到比他大的数 将右指针停止 移动左指针 当当前的数比基准大于的话
	// 就移动左指针 当比他小的时候 就停止  然后交换两边指针的数据
	// 最后交换起始位置 和当前left指针的数据  即左边的小于中间的 右边的大于中间的
	for left < right {
		for left < right && nums[right] < base {
			right--
		}
		for left < right && nums[left] >= base {
			left++
		}
		if left < right {
			p := nums[left]
			nums[left] = nums[right]
			nums[right] = p
		}
	}
	nums[startIndex] = nums[left]
	nums[left] = base
	return left;
}
func mid2(nums []int,startIndex,lastIndex int) int {
	//设置基准为起始点 mark为当前节点
	base := nums[startIndex]
	mark := startIndex
	// 从mark的下一个节点开始比较 如果大于mark的话 就mark向前移动一位
	// 并且交换他们的距离
	for i:= startIndex+1 ;i<lastIndex;i ++  {
		if(nums[i] > base) {
			mark++
			p:= nums[mark]
			nums[mark] = nums[i]
			nums[i] =p
		}
	}
	nums[startIndex] = nums[mark]
	nums[mark] = base

	return mark
}