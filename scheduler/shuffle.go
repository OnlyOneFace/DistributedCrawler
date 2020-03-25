package scheduler

import (
	"math/rand"
	"time"
)

/*
* @
* @Author:洗牌算法 rand.Perm
* @Date: 2020/3/24 22:39
 */
func Shuffle(arr []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var (
		n         = len(arr)
		randIndex int
	)
	for i := n; i > 0; i-- {
		randIndex = r.Intn(i)
		arr[i-1], arr[randIndex] = arr[randIndex], arr[i-1]
	}
}
