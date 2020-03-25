package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3,4}
	Permutation(nums, 0)

}

//func Permutation(nums []int, start int) {
//	// 递归退出的条件
//	if start == len(nums) {
//		fmt.Println(nums)
//	}
//	// 从当前位置开始到最后一个数的位置
//	for i:=start;i<=len(nums)-1;i++ {
//		//把第一个元素与后面的元素进行交换,递归的调用子数组进行排序
//		swap(nums,i,start)
//		Permutation(nums,start+1)
//		//子数组排序返回后要将第一个元素交换回来。
//		//如果不交换回来会出错，比如说第一次1、2交换，第一个位置为2，子数组排序返回后如果不将1、2
//		//交换回来第二次交换的时候就会将2、3交换，因此必须将1、2交换使1还是在第一个位置
//		swap(nums,i,start)
//	}
//}
//func swap(nums []int,i,start int)  {
//	p:= nums[i]
//	nums[i] = nums[start]
//	nums[start] = p
//}
func Permutation(nums []int, start int) {

	// 递归结束的条件
	if start == len(nums) {
		fmt.Println(nums)
	}

	//形如 1 2 3 4 第一位其实有四种情形 其实就是 1 2 3 4 四种情况 每次都是把数字进行对换
	for i:=start;i<=len(nums)-1;i++ {
		//将i和start进行交换 每次输出的就是第一个
		swap(nums,i,start)
		//fmt.Println(nums) //
		// // 以上是对第一个数字的判断 可以将数字更加的像后面递归 只需将每次的数组像后面+1
		Permutation(nums,start+1)
		// 此时假如是第二次交换  数组应该是 3 2 1 4  为了下次方便
		// 需要将数组变为原来的样子
		swap(nums,i,start)

	}
}

func swap(nums []int, start,end int)  {
	p:= nums[start]
	nums[start] = nums[end]
	nums[end] = p
}
