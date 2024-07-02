package main

import (
	"container/ring"
	"fmt"
)

type WeightedItem struct {
	Value  interface{}
	Weight int
}

type WeightedRoundRobin struct {
	ring *ring.Ring
}

// NewWeightedRoundRobin initializes a new weighted round-robin structure with the given items.
func NewWeightedRoundRobin(items []WeightedItem) *WeightedRoundRobin {
	totalWeight := 0
	for _, item := range items {
		totalWeight += item.Weight
	}

	r := ring.New(totalWeight)

	// Populate the ring according to the weight of each item
	for _, item := range items {
		for i := 0; i < item.Weight; i++ {
			r.Value = item.Value
			r = r.Next()
		}
	}

	return &WeightedRoundRobin{ring: r}
}

func (wr *WeightedRoundRobin) Next() interface{} {
	// Get the value of the current element in the ring.
	value := wr.ring.Value

	wr.ring = wr.ring.Next()

	return value
}

func main() {
	items := []WeightedItem{
		{Value: "A", Weight: 1},
		{Value: "B", Weight: 3},
		{Value: "C", Weight: 2},
	}

	wrr := NewWeightedRoundRobin(items)

	for i := 0; i < 100; i++ {
		fmt.Println(wrr.Next())
	}
}
