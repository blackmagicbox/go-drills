# 001 — Job queue after close

- **Topic:** channel semantics (Q3) — receiving from a closed channel
- **Format:** predict
- **Time:** ~15 min
- **Graded by:** running it, then `ANSWER.md`

## Scenario

labctl queues VM-provisioning jobs on a channel. The producer queues three jobs and
closes the queue. The worker is badly written: it reads a fixed five times instead of
ranging over the channel.

## Task

1. Read `main.go`. **Don't run it yet.**
2. In `NOTES.md`, write the exact five lines you expect it to print, and one sentence on *why*
   for lines 4 and 5.
3. Bonus prediction: what would happen if the worker used `for j := range jobs` instead?
4. Run `go run ./001-channels-predict` and compare. Note every surprise.
5. Commit, then read `ANSWER.md`.
