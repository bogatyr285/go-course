package coverage

import "testing"

// go test -coverprofile profile
// go tool cover -html=profile

func TestBazBaz(t *testing.T) {
	expected := 3
	actual := BazBaz(3)
	if actual != expected {
		t.Errorf("actual %d, expected %d", actual, expected)
	}
}

func TestBazBaz2(t *testing.T) {
	expected := 25
	actual := BazBaz(25)
	if actual != expected {
		t.Errorf("actual %d, expected %d", actual, expected)
	}
}
