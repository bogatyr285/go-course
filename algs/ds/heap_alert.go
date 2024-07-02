package main

import (
	"container/heap"
	"fmt"
)

type Alert struct {
	Priority    int // Priority level, lower number means higher priority. 1- Critical
	Description string
	Index       int // The index of the item in the heap (used internally by the heap).
}

// max heap of Alerts.
type AlertHeap []*Alert

func (h AlertHeap) Len() int { return len(h) }
func (h AlertHeap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
} // Lower priority number is more urgent
func (h AlertHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *AlertHeap) Push(x interface{}) {
	n := len(*h)
	alert := x.(*Alert)
	alert.Index = n
	*h = append(*h, alert)
}

func (h *AlertHeap) Pop() interface{} {
	old := *h         // Get the slice of current elements.
	n := len(old)     // Get the current length of the heap.
	alert := old[n-1] // Get the element at the end of the slice (highest priority).
	old[n-1] = nil    // Prevent memory leak by setting the last element to nil.
	alert.Index = -1  // Mark the Index field to -1 for safety and debug
	*h = old[0 : n-1] // Reduce the size of the slice by one (removing the last element).
	return alert      // Return the removed element.
}

type MonitoringSystem struct {
	alertHeap *AlertHeap
}

func NewMonitoringSystem() *MonitoringSystem {
	h := &AlertHeap{}
	heap.Init(h)
	return &MonitoringSystem{alertHeap: h}
}

func (ms *MonitoringSystem) AddAlert(priority int, description string) {
	alert := &Alert{
		Priority:    priority,
		Description: description,
	}
	heap.Push(ms.alertHeap, alert)
}

func (ms *MonitoringSystem) ProcessAlerts() {
	for ms.alertHeap.Len() > 0 {
		alert := heap.Pop(ms.alertHeap).(*Alert)
		fmt.Printf("Processing Alert: %s (Priority: %d)\n", alert.Description, alert.Priority)
	}
}

func main() {
	ms := NewMonitoringSystem()

	ms.AddAlert(1, "Critical: Database is down!")
	ms.AddAlert(2, "High: High memory usage detected.")
	ms.AddAlert(3, "Medium: Disk space is running low.")
	ms.AddAlert(1, "Critical: Service A is not responding.")
	ms.AddAlert(2, "High: High CPU usage detected.")

	// Process alerts by priority
	ms.ProcessAlerts()
}
