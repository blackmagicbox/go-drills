package logs

// Merge returns a channel that emits every line from stdout and stderr,
// and is closed once both inputs are closed. See README.md for the full spec.
func Merge(stdout, stderr <-chan string) <-chan string {
	panic("TODO: implement Merge")
}
