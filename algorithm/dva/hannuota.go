package main

import "fmt"

func main()  {
	hannuota(3,"a","b","c")
}
// 汉诺塔问题
func hannuota(num int,a,b,c string)  {
	//1. 分治算法的经典案例
	// 先将盘子看成两个盘子,上面的统一看成一个盘子,因此就为n-1,
	// 下面的一个看成一个盘子为n
	// 一个盘子的步骤为 a-c
	// 两个盘子的思路为 a-b a-c b-c
	// 三个盘子的思路为...
	if num == 1 {
		fmt.Printf("第1个盘子从%s到%s\n",a,c)
	}else {
		// 先看成上面的盘子
		hannuota(num-1,a,c,b)
		// 当前的盘子
		fmt.Printf("第%d盘子从%s到达%s\n",num,a,c)
		// 在B上面的盘子
		hannuota(num-1,b,a,c)
	}
}
