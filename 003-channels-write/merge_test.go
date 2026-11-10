package logs

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func feed(lines []string, delay time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, l := range lines {
			time.Sleep(delay)
			ch <- l
		}
	}()
	return ch
}

// collect reads out until it is closed, failing the test if that takes too long.
func collect(t *testing.T, out <-chan string) []string {
	t.Helper()
	var got []string
	timeout := time.After(2 * time.Second)
	for {
		select {
		case l, ok := <-out:
			if !ok {
				return got
			}
			got = append(got, l)
		case <-timeout:
			t.Fatalf("output not closed after 2s; got %d lines so far: %v", len(got), got)
		}
	}
}

func lines(prefix string, n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = fmt.Sprintf("%s-%d", prefix, i)
	}
	return out
}

func checkAll(t *testing.T, got, a, b []string) {
	t.Helper()
	if len(got) != len(a)+len(b) {
		t.Fatalf("got %d lines, want %d: %v", len(got), len(a)+len(b), got)
	}
	for _, want := range [][]string{a, b} {
		i := 0
		for _, l := range got {
			if i < len(want) && l == want[i] {
				i++
			}
		}
		if i != len(want) {
			t.Fatalf("lines %v missing or out of order in %v", want, got)
		}
	}
}

func TestMergesBoth(t *testing.T) {
	out, errs := lines("out", 5), lines("err", 5)
	got := collect(t, Merge(feed(out, time.Millisecond), feed(errs, time.Millisecond)))
	checkAll(t, got, out, errs)
}

func TestStderrEndsFirst(t *testing.T) {
	out, errs := lines("out", 20), lines("err", 2)
	got := collect(t, Merge(feed(out, 2*time.Millisecond), feed(errs, 0)))
	checkAll(t, got, out, errs)
}

func TestBothEmpty(t *testing.T) {
	got := collect(t, Merge(feed(nil, 0), feed(nil, 0)))
	if len(got) != 0 {
		t.Fatalf("got %v, want no lines", got)
	}
}

func TestNoGoroutineLeak(t *testing.T) {
	before := runtime.NumGoroutine()
	for i := 0; i < 10; i++ {
		collect(t, Merge(feed(lines("out", 3), 0), feed(lines("err", 1), 0)))
	}
	time.Sleep(50 * time.Millisecond)
	if leaked := runtime.NumGoroutine() - before; leaked > 0 {
		t.Fatalf("%d goroutines still running after all outputs closed", leaked)
	}
}
