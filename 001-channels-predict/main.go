package main

import "fmt"

type Job struct {
	VM   string
	CPUs int
}

func main() {
	jobs := make(chan Job, 3)

	// Producer: queue three jobs, then close the queue.
	jobs <- Job{VM: "cow-1", CPUs: 2}
	jobs <- Job{VM: "cow-2", CPUs: 2}
	jobs <- Job{VM: "cow-3", CPUs: 4}
	close(jobs)

	// Worker: reads a fixed five times (bad idea, on purpose).
	for i := 1; i <= 5; i++ {
		j, ok := <-jobs
		fmt.Printf("read %d: vm=%q cpus=%d ok=%v\n", i, j.VM, j.CPUs, ok)
	}
}
