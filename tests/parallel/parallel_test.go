package parallel

import (
	"testing"
	"time"
)

func TestHighlyLoaded1(t *testing.T) {
	// t.Parallel()
	time.Sleep(300 * time.Millisecond)
}

func TestHighlyLoaded2(t *testing.T) {
	// t.Parallel()
	time.Sleep(300 * time.Millisecond)
}

func TestHighlyLoaded3(t *testing.T) {
	// t.Parallel()
	time.Sleep(300 * time.Millisecond)
}
