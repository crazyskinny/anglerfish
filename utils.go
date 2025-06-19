package anglerfish

import (
	"errors"
	"math/rand"
)

type WeightedItem[T any] struct {
	Item   T
	Weight int
}

// WeightedRandomPick selects an item from a list of weighted items based on their weights.
func WeightedRandomPick[T any](items []WeightedItem[T]) (T, error) {
	var zero T
	if len(items) == 0 {
		return zero, errors.New("items list is empty")
	}
	prefixSum := make([]int, len(items))
	total := 0
	for i, wi := range items {
		if wi.Weight < 0 {
			return zero, errors.New("weight must be non-negative")
		}
		total += wi.Weight
		prefixSum[i] = total
	}
	if total == 0 {
		return zero, errors.New("total weight is zero")
	}
	r := rand.Intn(total)
	for i, bound := range prefixSum {
		if r < bound {
			return items[i].Item, nil
		}
	}
	return zero, errors.New("unexpected selection failure")
}
