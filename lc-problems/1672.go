package problems

import (
	"sync"
)

func MaximumWealth(accounts [][]int) int {
	maxAmount := 0

	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
			if sum > maxAmount {
				maxAmount = sum
			}
		}
	}
	return maxAmount
}

/*
Example 1:

Input: accounts = [[1,2,3],[3,2,1]]
Output: 6
Explanation:
1st customer has wealth = 1 + 2 + 3 = 6
2nd customer has wealth = 3 + 2 + 1 = 6
Both customers are considered the richest with a wealth of 6 each, so return 6.
Example 2:

Input: accounts = [[1,5],[7,3],[3,5]]
Output: 10
Explanation:
1st customer has wealth = 6
2nd customer has wealth = 10
3rd customer has wealth = 8
The 2nd customer is the richest with a wealth of 10.
Example 3:

Input: accounts = [[2,8,7],[7,1,3],[1,9,5]]
Output: 17

*/

func MaximumWealthUsingChan(accounts [][]int) int {
	wg := sync.WaitGroup{}
	maxAmount := 0
	ch := make(chan int, len(accounts))

	for _, account := range accounts {
		wg.Add(1)
		go func(account []int) {
			defer wg.Done()
			accountSum := 0
			for _, val := range account {
				accountSum += val
			}
			ch <- accountSum
		}(account)
	}

	wg.Wait()
	close(ch)

	for accountSum := range ch {
		if accountSum > maxAmount {
			maxAmount = accountSum
		}
	}

	return maxAmount
}
