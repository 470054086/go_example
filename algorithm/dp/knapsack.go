package main

import (
	"fmt"
	"math"
)

/**
有N件物品和一个体积为V的背包。（每种物品均只有一件）第i件物品的体积是volume[i]，价值是value[i]。
求解将哪些物品装入背包可使这些物品的体积总和不超过背包体积，且价值总和最大

// 背包问题
我们说的是如果第i个物品的体积小于背包体积V，可以放可以不放.
那么就比较前i-1个物品在V背包容量的最优解价值和前i-1的物品在容积为V减去第i个物品体积的背包容量的最优解加上第i个物品的价值，哪个大？
如果前者大，那第i个物品就不放，如果后者大，那就放。因为如果要放第i个必须腾出位置来，来比较放入和不放入的价值。这样我们状态转换方程就出来了：
	p[i][j]=MAX{p[i-1][j-volume[i]]+value[i],p[i-1][j]};

p[i][j]代表前i件物品组合在容量为j的背包的最优解。将前i件物品放入容量为v的背包中这个子问题，若只考虑第i件物品的策略（放或不放），
那么就可以转化为一个只牵扯前i-1件物品的问题。如果不放第i件物品，那么问题就转化为“前i-1件物品放入容量为v的背包中，价值为p[i-1][v]；
如果放第i件物品，那么问题就转化为“前i-1件物品放入剩下的容量为v-volume[i]的背包中”，
此时能获得的最大价值就是p[i-1][j-volume[i]]再加上通过放入第i件物品获得的价值value[i]


*/
func main() {
	num := 5
	val := 10
	total := []int{0, 6, 3, 5, 4, 6}
	weight := []int{0, 2, 2, 6, 5, 4}
	res := knapsack(num, val, weight, total)
	fmt.Println(res)
}

/**

 */
func knapsack(n, w int, weight []int, total []int) int {
	//  生成一个二维数组
	var dp [6][11]int

	for i := 1; i <= n; i++ { //物品
		for j := 1; j <= w; j++ { //背包
			//当前物品比背包中,装不下，肯定就是不装
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else { //装的下 max(装物品,不装物品)
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-weight[i]]+total[i])))
			}
		}
	}
	return dp[n][w]

}
