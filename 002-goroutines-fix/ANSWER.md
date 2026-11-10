# 002 — Answer (open only after committing your attempt)

`TestNoGoroutineLeak` fails with about `10 goroutines leaked`. The other two pass:
the bug doesn't change what the function *returns*, only what it leaves behind.

## Why it leaks

On timeout, `CheckHealth` returns via `<-ctx.Done()`. Nobody will ever receive from
`result` again (it's a local variable that just went out of scope), so the probe goroutine
blocks forever on `result <- probe()`: unbuffered sends wait for a receiver.

- **Leak, not deadlock:** only that one goroutine is stuck; the rest of the program runs.
  Go's "all goroutines are asleep" detector fires only when *every* goroutine is blocked.
- **The caller can't drain it:** `result` never escapes the function, so nothing outside can
  reach it once `CheckHealth` returns.

## Fix 1: buffered channel of size 1

`make(chan error, 1)` — the send always has room, completes, and the goroutine exits.
The value is simply garbage-collected with the channel. Minimal and idiomatic.

## Fix 2: the goroutine listens to the context too

Inside the goroutine, `select` between sending on `result` and `<-ctx.Done()`.
If the caller has given up, the goroutine takes the `ctx.Done()` branch and exits.
Note: this only works because the goroutine *checks* ctx — passing a context around does
nothing by itself. (The probe still runs to completion; to stop the probe itself, the probe
would need to accept ctx.)
