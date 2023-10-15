package main

/*
   @Time: 2023/10/15 19:41
   @Author: david
   @File: main
*/

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
