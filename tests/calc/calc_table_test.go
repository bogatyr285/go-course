package main

import "testing"

func TestTotalPrice(t *testing.T) {
	tests := []struct {
		name     string
		nights   uint
		rate     uint
		cityTax  uint
		expected uint
	}{
		{"Normal case", 3, 100, 10, 30},  // (3 * (100 + 10)) = 330
		{"Zero nights", 0, 100, 10, 0},   // (0 * (100 + 10)) = 0
		{"Zero rate", 3, 0, 10, 30},      // (3 * (0 + 10)) = 30
		{"Zero cityTax", 3, 100, 0, 300}, // (3 * (100 + 0)) = 300
		{"All zeros", 1, 0, 0, 0},        // (0 * (0 + 0)) = 0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalPrice(tt.nights, tt.rate, tt.cityTax)
			if got != tt.expected {
				t.Errorf("totalPrice(%d, %d, %d) = %d; want %d", tt.nights, tt.rate, tt.cityTax, got, tt.expected)
			}
		})
	}
}

func Test_totalPrice(t *testing.T) {
	type parameters struct {
		nights  uint
		rate    uint
		cityTax uint
	}
	type testCase struct {
		name string
		args parameters
		want uint
	}
	tests := []testCase{
		{
			name: "test 0 nights",
			args: parameters{nights: 0, rate: 150, cityTax: 12},
			want: 0,
		},
		{
			name: "test 1 nights",
			args: parameters{nights: 1, rate: 100, cityTax: 12},
			want: 112,
		},
		{
			name: "test 2 nights",
			args: parameters{nights: 2, rate: 100, cityTax: 12},
			want: 224,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalPrice(tt.args.nights, tt.args.rate, tt.args.cityTax); got != tt.want {
				t.Errorf("totalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
