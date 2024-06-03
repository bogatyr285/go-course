package main

import "testing"

func TestPrice(t *testing.T) {
	var nights uint = 3
	var rate uint = 10000
	var cityTax uint = 132

	var actual uint = 30396
	got := totalPrice(nights,
		rate,
		cityTax,
	)

	if got != actual {
		t.Errorf("totalPrice(%d, %d, %d) = %d; want %d", nights, rate, cityTax, got, actual)
	}
}
