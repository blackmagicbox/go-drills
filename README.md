# go-drills

Short practice exercises (15–20 min) between labs. Part of the
"Learning Go without working with Go" course (state lives in the `brain` vault:
`courses/learning-go-without-working-with-go/state.md`).

## Rules

1. **Predict before you run.** Write your prediction down (in the exercise's `NOTES.md`) first.
2. **No looking things up on the first attempt.** Write down where you got stuck, *then* check docs.
3. **Copy exact error messages.** Never "something like that".
4. **Don't open `ANSWER.md` until you've committed your attempt.**
5. Done = tests pass with `go test -race ./NNN-*/...` (or the self-check in the README).

## Formats

| Format | What you do | Graded by |
|---|---|---|
| predict | Read code, predict output in free text, then run | Running it + `ANSWER.md` |
| fix | Realistic code with a bug; make the tests pass | `go test -race` |
| explain | Explain real code (stdlib / k8s / Prometheus) in plain English | Claude review |
| complete | Fill the gaps in nearly-finished code | `go test -race` |
| parsons | Put scrambled lines back in order | `go test -race` |
| write | Implement from a README + tests only | `go test -race` |
| quiz | Multiple choice / true–false warm-ups | In chat with Claude |

Folder names: `NNN-topic-format`. Each README lists topic, format, time and how it's graded.

## Current set — channels & goroutine leaks (+ errors mixed in)

| # | Exercise | Format |
|---|---|---|
| 001 | Job queue after close | predict |
| 002 | Health check that leaks | fix |
| 003 | Merge a VM's log streams | write |
