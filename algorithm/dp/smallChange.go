package main

import (
	"fmt"
	"math"
)

/**
零钱相关问题
题目：给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，再给一个总金额 n，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，则回答 -1
[1,2,5]  total=11

推导公式 总额为11的
1. 凑成面值为10的硬币,加上一枚1块的硬币
2. 凑成面值为9的硬币,加上一枚2块的硬币
3. 凑成面值为6的硬币,加上一枚5块的硬币

fn = min(dp[total-1]+1,dp[total-2]+1,dp[total-5]+1))
*/

func main() {
	m := []int{1, 2, 5}
	total := 11
	count := smalltest3(m, total)
	fmt.Println(count)
	//change := smallChange3(m, 11)
	//fmt.Println(change)
}

//func smallTest(money []int, total int) int {
//	if total <= 0 {
//		return 0
//	}
//	//定义一个最大值
//	ans := math.MaxInt32
//	//循环这个数组 递归进行调用
//	for _, val := range money {
//		if total-val < 0 {
//			continue
//		}
//		subPron := smallTest(money, total-val)
//		if subPron == -1 {
//			continue
//		}
//		// 获取最小的
//		ans = int(math.Min(float64(subPron+1), float64(ans)))
//	}
//	if ans == math.MaxInt32 {
//		return -1
//	} else {
//		return ans
//	}
//}
//
//func smalltest2(money []int, total int) int {
//	// 生成一个用来保存值得数组
//	memo := make([]int, total+1)
//	for k := range memo {
//		memo[k] = -2
//	}
//	return smalltest2Help(money, total, memo)
//
//}
//
//func smalltest2Help(money []int, total int, memo []int) int {
//	if total <= 0 {
//		return 0
//	}
//	if memo[total] != -2 {
//		return memo[total]
//	}
//	ans := math.MaxInt32
//	for _, val := range money {
//		if total-val < 0 {
//			continue
//		}
//		subPron := smalltest2Help(money, total-val, memo)
//		if subPron == -1 {
//			continue
//		}
//		ans = int(math.Min(float64(subPron+1), float64(ans)))
//	}
//	if ans == math.MaxInt32 {
//		memo[total] = -1
//	} else {
//		memo[total] = ans
//	}
//	return memo[total]
//
//}
//
//func smalltest3(money []int, total int) int {
//	// 生成一个数组用来保存值
//	memo := make([]int, total+1)
//	for k := range memo {
//		memo[k] = math.MaxInt32
//	}
//	memo[0] = 0
//	for i := 1; i <= total; i++ {
//		for _, val := range money {
//			if i-val < 0 {
//				continue
//			}
//			memo[i] = int(math.Min(float64(memo[i]), float64(memo[i-val]+1)))
//		}
//	}
//	if memo[total] == math.MaxInt32 {
//		return -1
//	} else {
//		return memo[total]
//	}
//}

func smallChange(money []int, total int) int {
	if total == 0 {
		return -1
	}
	ans := math.MaxInt32
	for _, val := range money {
		//  金额不可达
		if total-val < 0 {
			continue
		}
		subPron := smallChange(money, total-val)
		if subPron == -1 {
			continue
		}
		ans = int(math.Min(float64(ans), float64(subPron+1)))
	}
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func smallChange2(money []int, total int) int {
	memo := make([]int, total+1)
	for k := range memo {
		memo[k] = -2
	}
	return helperSmall(money, total, memo)
}

func helperSmall(money []int, total int, memo []int) int {
	if total == 0 {
		return 0
	}
	if memo[total] != -2 {
		return memo[total]
	}
	ans := math.MaxInt32
	for _, val := range money {
		//  金额不可达
		if total-val < 0 {
			continue
		}
		subPron := helperSmall(money, total-val, memo)
		if subPron == -1 {
			continue
		}
		ans = int(math.Min(float64(ans), float64(subPron+1)))
	}

	if ans == math.MaxInt32 {
		memo[total] = -1
	} else {
		memo[total] = ans
	}
	return memo[total]
}
func smallChange3(money []int, total int) int {
	memo := make([]int, total+1)
	for k := range memo {
		memo[k] = total + 1
	}
	memo[0] = 0
	for i := 0; i < len(memo); i++ {
		for _, val := range money {
			if i-val < 0 {
				continue
			}
			memo[i] = int(math.Min(float64(memo[i]), float64(1+memo[i-val])))
		}
	}
	if memo[total] == total+1 {
		return -1
	} else {
		return memo[total]
	}

}
