package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Event struct {
	Timestamp   time.Time
	Description string
	Index       int // The index of the item in the heap.
}

type PriorityQueue []*Event

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want the earliest timestamp to have the highest priority
	return pq[i].Timestamp.Before(pq[j].Timestamp)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	event := x.(*Event)
	event.Index = n
	*pq = append(*pq, event)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil   // Avoid memory leak
	event.Index = -1 // For safety
	*pq = old[:n-1]
	return event
}

func main() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	events := []*Event{
		{Timestamp: time.Now().Add(10 * time.Minute), Description: "Event 3"},
		{Timestamp: time.Now().Add(5 * time.Minute), Description: "Event 2"},
		{Timestamp: time.Now().Add(1 * time.Minute), Description: "Event 1"},
	}

	for _, event := range events {
		heap.Push(&pq, event)
	}

	// Processing events in the order of their timestamps.
	fmt.Println("Processing events by timestamp:")
	for pq.Len() > 0 {
		event := heap.Pop(&pq).(*Event)
		fmt.Printf("%s: %s\n", event.Timestamp.Format(time.RFC3339), event.Description)
	}
}
