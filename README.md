# go-pool

We simulate an expensive routine by allocating a buffer of 5MiB at each procedure call.

Usage: 
```
go run main.go (-c CALLS) (-p POOL)

Options:
    -c, --calls              Number of expensive calls to simulate
    -p, --pool               If present enables pool

Output:
	Generates a mem.pprof file in the wd.

Example:
	go run main.go --calls 1_000_000 --pool
	go run main.go --calls 1_000_000
```