package health

import (
	"context"
	"errors"
	"runtime"
	"testing"
	"time"
)

var errNodeDown = errors.New("node down")

func TestReturnsProbeError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := CheckHealth(ctx, func() error { return errNodeDown })
	if !errors.Is(err, errNodeDown) {
		t.Fatalf("got %v, want errNodeDown", err)
	}
}

func TestTimesOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	err := CheckHealth(ctx, func() error {
		time.Sleep(50 * time.Millisecond)
		return nil
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v, want context.DeadlineExceeded", err)
	}
}

func TestNoGoroutineLeak(t *testing.T) {
	before := runtime.NumGoroutine()

	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		_ = CheckHealth(ctx, func() error {
			time.Sleep(50 * time.Millisecond)
			return nil
		})
		cancel()
	}

	// Give every slow probe time to finish.
	time.Sleep(200 * time.Millisecond)

	after := runtime.NumGoroutine()
	if leaked := after - before; leaked > 0 {
		t.Fatalf("%d goroutines leaked after 10 timed-out checks", leaked)
	}
}
