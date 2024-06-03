package main

import (
	"fmt"
	"time"
)

//
//func RunProcessor(wg sync.WaitGroup, prices []map[string]float64) {
// go func() {
//     defer wg.Done()
//     for _, price := range prices {
//         for key, value := range price {
//             price[key] = value + 1
//         }
//         fmt.Println(price)
//     }
// }()
//}

func RunWriter() <-chan map[string]float64 {
	var prices = make(chan map[string]float64)
	go func() {
		var currentPrice = map[string]float64{
			"inst1": 1.1,
			"inst2": 2.1,
			"inst3": 3.1,
			"inst4": 4.1,
		}
		for i := 1; i < 5; i++ {
			for key, value := range currentPrice {
				currentPrice[key] = value + 1
			}
			prices <- currentPrice
			time.Sleep(time.Second)
		}
		close(prices)
	}()
	return prices
}
func main() {
	p := RunWriter()
	var prices []map[string]float64

	for price := range p {
		prices = append(prices, price)
	}

	for _, price := range prices {
		fmt.Println(price)
	}

	//wg := sync.WaitGroup{}
	//wg.Add(3)
	//RunProcessor(wg, prices)
	//RunProcessor(wg, prices)
	//RunProcessor(wg, prices)
	//wg.Wait()
}
