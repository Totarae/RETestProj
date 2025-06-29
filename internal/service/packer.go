package service

import (
	"sort"
)

type PackResult struct {
	Packs map[int]int // pack size and count
	Total int         // total number of itms shipped
}

type dpState struct {
	packCount int
	packMap   map[int]int
}

// Coin Change problem
// OptimizePacks optimal set of packs.
// rule 1: minimize total items >= order
// rul2 2: minimize number of packs
func OptimizePacks(order int, packSizes []int) PackResult {
	// empty
	if order <= 0 || len(packSizes) == 0 {
		return PackResult{Packs: map[int]int{}, Total: 0}
	}

	//Presort all available packs asc order
	sort.Ints(packSizes)

	// maxSize is the largest pack size available
	// maxLimit is the largest number of items to consider
	maxSize := packSizes[len(packSizes)-1]
	maxLimit := order + maxSize

	dp := make([]*dpState, maxLimit+1)
	dp[0] = &dpState{packCount: 0, packMap: map[int]int{}}

	for i := 1; i <= maxLimit; i++ {
		bestPackCount := -1
		var bestPackMap map[int]int

		for _, size := range packSizes {
			prev := i - size
			if prev < 0 || dp[prev] == nil {
				continue
			}

			candidatePackCount := dp[prev].packCount + 1
			if bestPackCount == -1 || candidatePackCount < bestPackCount {
				newMap := copyMap(dp[prev].packMap)
				newMap[size]++
				bestPackCount = candidatePackCount
				bestPackMap = newMap
			}
		}

		if bestPackMap != nil {
			dp[i] = &dpState{
				packCount: bestPackCount,
				packMap:   bestPackMap,
			}
		}
	}

	return checkSolution(dp, order, maxLimit)
}

// checkSolution chaecks a valid combination.
func checkSolution(dp []*dpState, order int, maxLimit int) PackResult {
	var best *dpState

	// for equal combinations
	minIndex := -1

	for i := order; i <= maxLimit; i++ {
		if dp[i] == nil {
			continue
		}
		if best == nil ||
			i < minIndex ||
			(i == minIndex && dp[i].packCount < best.packCount) {

			best = dp[i]
			minIndex = i
		}
	}

	if best == nil {
		return PackResult{Packs: map[int]int{}, Total: 0}
	}

	return PackResult{
		Packs: copyMap(best.packMap),
		Total: minIndex}
}

func copyMap(original map[int]int) map[int]int {
	copy := make(map[int]int)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}
