package other

import (
	"sort"
)

func KClosest(points [][]int, K int) [][]int {
	kclosest := make([][]int, K)
	l := make([]pair, len(points))
	for i, point := range points {
		key := point[0]*point[0] + point[1]*point[1]
		entry := pair{
			sumSquare: key,
			value:     point,
		}
		l[i] = entry
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].sumSquare < l[j].sumSquare
	})
	//fmt.Println(l)
	for i := 0; i < K; i++ {
		kclosest[i] = l[i].value
	}
	return kclosest
}

type pair struct {
	sumSquare int
	value     []int
}
