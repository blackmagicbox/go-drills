# 002 — Health check that leaks

- **Topic:** goroutine leaks (Q6) + `context`; mixed in: errors (`errors.Is`)
- **Format:** fix
- **Time:** ~20 min
- **Graded by:** `go test -race ./002-goroutines-fix/`

## Scenario

berlin-herd's health checker probes each node. `CheckHealth` runs the probe in a goroutine
and gives up when the context deadline passes. In production the checker's memory grows
slowly forever, and `runtime.NumGoroutine()` keeps climbing.

## Task

1. Read `health.go` and the tests. **Before running:** write in `NOTES.md` which test fails,
   what number you expect the leak test to report, and why.
2. Run the tests. Copy the exact failure message into `NOTES.md`.
3. Fix `health.go` so all tests pass. Change as little as possible.
4. Then fix it a **second, different way** (keep both versions in `NOTES.md`; leave one in the file).
5. In one sentence each: why is this a *leak* and not a *deadlock*? Why can't the caller
   "drain" the channel after `CheckHealth` has returned?

Rules: don't change the test file or the function signature.
