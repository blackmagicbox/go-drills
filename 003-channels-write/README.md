# 003 — Merge a VM's log streams

- **Topic:** nil channels in `select` (Q4), closed-channel receive (Q3); mixed in: goroutine leaks (Q6)
- **Format:** write
- **Time:** ~20–30 min
- **Graded by:** `go test -race ./003-channels-write/`

## Scenario

labctl streams a VM's console output as two channels: `stdout` and `stderr`.
The UI wants **one** stream. Implement `Merge` in `merge.go`.

## Spec (the tests check all of it)

1. Every line from both inputs appears on the output exactly once.
2. Lines from the same input keep their order.
3. The output channel is **closed** once *both* inputs are closed and drained — not earlier.
4. Inputs may close at different times (stderr often ends first).
5. No goroutine is left running after the output is closed.
6. Passes with `-race`.

## Constraints

- One goroutine and one `select` loop. No `sync.WaitGroup` this time.
- Before writing code, answer in `NOTES.md`: once `stderr` is closed, what does
  `case line, ok := <-stderr` do on every loop iteration — and how do you make `select`
  stop picking it?

Hints are only in your memory and the `state.md` weak spots. No `ANSWER.md` until you've tried.
