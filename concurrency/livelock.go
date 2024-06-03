package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type value struct {
		sync.Mutex
		id     string
		locked bool
	}

	lock := func(v *value) {
		v.Lock()
		v.locked = true
	}
	unlock := func(v *value) {
		v.Unlock()
		v.locked = false
	}
	move := func(wg *sync.WaitGroup, id string, v1, v2 *value) {
		defer wg.Done()
		for i := 0; ; i++ {
			if i >= 3 {
				fmt.Println("canceling goroutine... ", id)
				return
			}

			fmt.Printf("%v: trying to move behind\n", v1.id)
			lock(v1) // <1>

			time.Sleep(3 * time.Second)

			if v2.locked { // <2>
				fmt.Printf("%v: moving ahead, blocked by %v\n", v1.id, v2.id)
				unlock(v1) // <3>
				continue
			}
		}
	}
	a, b, c, d := value{id: "car1"}, value{id: "car2"}, value{id: "car3"}, value{id: "car4"}
	var wg sync.WaitGroup
	wg.Add(4)
	go move(&wg, "first", &a, &b)
	go move(&wg, "second", &b, &c)
	go move(&wg, "third", &c, &d)
	go move(&wg, "fourth", &d, &a)
	wg.Wait()
}
