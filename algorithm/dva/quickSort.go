package main

import "fmt"

func main() {
	var a = []int{5,2,3,1}
	quickSort2(a,0,len(a)-1)
	fmt.Println(a)
}

//func quickSort(nums []int, begin, last int) {
//	if len(nums) == 0 || len(nums) == 1 {
//		return
//	}
//	// 分支交换法
//	// 假设开始的就为基准点
//	base := nums[begin]
//	i := begin
//	j := last
//	mid := 0
//	// l为最左边 r为最右边 base为最左边的一个
//	for {
//		//从后面开始往前面找,和前面进行颠倒,
//		for j!=i  {
//			if nums[j] < base {
//				nums[i] = nums[j]
//				nums[j] = base
//				break
//			}else{
//				j --
//			}
//		}
//		if i == j  {
//			mid = i
//			break
//		}
//		// 从前面开始往后面找,如果找到的数字大于base的话 则进行交换
//		for i != j {
//			if nums[i] > base {
//				nums[j] = nums[i]
//				nums[i] = base
//				break
//			} else {
//				i++
//			}
//		}
//		if i == j {
//			mid = i
//			break
//		}
//		// 如果这个过程 l和r相同的话 则找到了
//	}
//	//如果mid在左边的话
//	if mid-begin > 1 {
//		quickSort(nums,begin,mid-1)
//	}
//	if last-mid > 1 {
//		quickSort(nums,mid+1 ,last)
//	}
//
//}
func quickSort2(nums []int,startIndex ,lastIndex int)  {
	// 递归结束的条件
	if startIndex >= lastIndex {
		return
	}
	// 获取基准点的位置
	mid := parrent2(nums,startIndex,lastIndex)
	quickSort2(nums,startIndex,mid-1)
	quickSort2(nums,mid+1,lastIndex)
}

func parrent(nums []int,startIndex,lastIndex int) int {
	// 选取第一个位置为起始节点
	base :=nums[startIndex]
	// 创建左右两个指针
	left :=startIndex
	right:= lastIndex
	// 只要left不等于right 就进行循环
	for left !=right {
		/**
		//1. 先从右指针开始移动 如果当前的数比left小于的话 就把指针左移
		// 直到遇到比他大的数 将右指针停止 移动左指针 当当前的数比基准大于的话
		// 就移动左指针 当比他小的时候 就停止  然后交换两边指针的数据
		// 最后交换起始位置 和当前left指针的数据  即左边的小于中间的 右边的大于中间的
		 */



		// 先从右指针开始循环 如果右指针大于基准元素 则向左移动
		for left<right && nums[right] > base {
			right--
		}
		// 从左指针开始 当左指针小于基准元素的时候 向右移动
		for left<right && nums[left] <= base {
			left++
		}
		// 当移动到位置的时候,将左右的元素进行交换
		if left < right {
			p:=nums[left]
			nums[left] = nums[right]
			nums[right] = p
		}
	}
	// 循环结束了之后 即left=right 需要将当前left的数据和基准元素进行交换
	nums[startIndex] = nums[left]
	nums[left]  = base
	return left
}
func parrent2(nums []int,startIndex,lastIndex int) int  {
	base:= nums[startIndex]
	mark := startIndex
	//从基准元素的下一个开始移动
	for i:=startIndex+1;i<= lastIndex;i++ {
		// 如果当前的值 小于基准元素 则进行交换
		// 把mark指针也向前移动
		if nums[i] < base {
			mark++
			p:= nums[mark]
			nums[mark] = nums[i]
			nums[i] = p
		}
	}
	//最后 将基准元素和当前mark的值进行交换
	nums[startIndex] = nums[mark]
	nums[mark] = base
	return mark
}