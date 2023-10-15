package lru

import (
	"strconv"
	"testing"
)

/*
   @Time: 2023/10/15 00:15
   @Author: david
   @File: lru_test
*/

func TestNewLRU(t *testing.T) {

}

func NewPool(size int) []string {
	start := 1679510
	pool := make([]string, size)
	for i := 0; i < size; i++ {
		uuid := strconv.FormatInt(int64(start+i), 10)
		pool[i] = uuid
	}
	return pool
}
