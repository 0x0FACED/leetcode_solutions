package task1043

import "math"

func MaxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxVal := math.MinInt32
		for j := 1; j <= min(i, k); j++ {
			maxVal = max(maxVal, arr[i-j])
			dp[i] = max(dp[i], dp[i-j]+maxVal*j)
		}
	}

	return dp[n]
}
