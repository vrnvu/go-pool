package main

import (
	"flag"
	"log"
	"sync"

	"github.com/pkg/profile"
)

const (
	_       = iota
	kib int = 1 << (10 * iota)
	mib
)

var (
	poolFlag  bool
	callsFlag int
)

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	flag.IntVar(&callsFlag, "c", 0, "number of expensive calls to simulate")
	flag.IntVar(&callsFlag, "calls", 0, "number of expensive calls to simulate")
	flag.BoolVar(&poolFlag, "p", false, "use pool")
	flag.BoolVar(&poolFlag, "pool", false, "use pool")
	flag.Parse()
	if callsFlag <= 0 {
		log.Fatalf("invalid calls flag: %d", callsFlag)
	}

	pool := &sync.Pool{
		New: func() any {
			return make([]byte, 5*mib)
		},
	}

	for i := 0; i < callsFlag; i++ {
		doSomething(pool)
	}
}

func doSomething(pool *sync.Pool) {
	h := pool.Get()
	if poolFlag {
		defer pool.Put(h)
	}
}
