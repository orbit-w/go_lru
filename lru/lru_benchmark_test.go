package lru

import (
	"github.com/alecthomas/assert/v2"
	"github.com/orbit-w/go_lru/v1/lib/misc"
	"math/rand"
	"testing"
)

/*
   @Time: 2023/10/15 14:45
   @Author: david
   @File: lru_benchmark_test
*/

func BenchmarkLRU_Set(b *testing.B) {
	lru, err := NewLRU[string, int8](10000)
	assert.NoError(b, err)
	size := int32(10000)
	pool := NewPool(int(size))
	b.Run("Set", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			uuid := pool[rand.Int31n(size)]
			lru.Set(uuid, 0)
		}
	})

	misc.PrintMem()
}
