# 001 — Answer (open only after committing your attempt)

```
read 1: vm="cow-1" cpus=2 ok=true
read 2: vm="cow-2" cpus=2 ok=true
read 3: vm="cow-3" cpus=4 ok=true
read 4: vm="" cpus=0 ok=false
read 5: vm="" cpus=0 ok=false
```

## Why

- Closing a channel does **not** throw away values already in its buffer. Reads 1–3 drain them.
- Once the channel is closed **and empty**, every receive returns **immediately** with the
  **zero value** of the element type (`Job{}` → `""` and `0`) and `ok == false`. It never blocks
  and never panics.
- `ok` is the only way to tell "a real zero value was sent" from "the channel is closed".

## Bonus

`for j := range jobs` stops by itself after the third job: `range` ends when the channel is
closed and drained. That's why idiomatic workers range over the channel.

## The rule to remember

| Operation | Open | Closed | nil |
|---|---|---|---|
| receive | waits for a value | buffered values, then zero value + ok=false at once | blocks forever |
| send | waits for room / receiver | **panic** | blocks forever |
| close | ok | **panic** | **panic** |
