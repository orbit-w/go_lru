# go_lru/v1

This provides the `LRU(Least recently used)` package which implements a fixed-size
thread safe LRU cache.

## Example
```go
package main

import (
	"github.com/orbit-w/go_lru/v1/lru"
)

func main() {
	ins, err := lru.NewLRU[int, int8](10000)
	if err != nil {
		panic("NewLRU failed")
	}
	for i := 0; i < 100; i++ {
		ins.Set(i, 0)
	}
}

```

## Benchmark
```go
package lru

import (
	"github.com/alecthomas/assert/v2"
	"github.com/orbit-w/go_lru/v1/lib/misc"
	"math/rand"
	"testing"
)

func BenchmarkLRU_Set(b *testing.B) {
	lru, err := NewLRU[string, int8](10000)
	assert.NoError(b, err)
	pool := NewPool(10000)
	b.Run("Set", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			uuid := pool[rand.Int31n(10000)]
			lru.Set(uuid, 0)
		}
	})

	misc.PrintMem()
}


````

```
pkg: github.com/orbit-w/go_lru/v1/lru
BenchmarkLRU_Set
BenchmarkLRU_Set/Set
BenchmarkLRU_Set/Set-8         	27125148	        43.45 ns/op	       0 B/op	       0 allocs/op
```