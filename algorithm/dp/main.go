package main

import (
	"sync"
	"time"
)
var wg sync.WaitGroup
func main() {


}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//  使用备忘录的方法
func fib2(n int )int  {
	if n<0 {
		return 0
	}
	//定义一个数组
	num := make([]int,n+1)
	num[1] = 1
	num[2] = 1
	// 将 1 2 都设置为1
	return helper(num,n)
}
func helper(num []int,n int)int  {
	// 将已经计算的备忘录 存储起俩
	if n>0 && num[n]==0 {
		num[n] = helper(num,n-1)+ helper(num,n-2)
	}
	return num[n]
}

func fib3(n int) int  {
	num := make([]int,n+1)
	num[1] = 1
	num[2] = 1
	for i:=3 ;i<n;i++ {
		num[i] = num[i-1] + num[i-2]
	}
	return num[n]
}

func times(i int ) int   {
	time.Sleep(time.Millisecond * time.Duration(i*100))
	return i
}


