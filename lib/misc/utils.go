package misc

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
   @Time: 2023/10/15 00:17
   @Author: david
   @File: utils
*/

func UsedNano(start, count int64) string {
	used := time.Now().UnixNano() - start
	return strings.Join([]string{"used: ", strconv.FormatInt(used, 10), "ns , ", strconv.FormatInt(used/count, 10), " ns/op "}, " ")
}

func PrintMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	println(fmt.Sprintf("TotalAlloc = %v MiB, HeapAlloc = %v MiB, NumGC = %v, HeapObjs = %v, Goroutine = %v",
		bToMb(m.TotalAlloc), bToMb(m.HeapAlloc), m.NumGC, m.HeapObjects, runtime.NumGoroutine()))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
