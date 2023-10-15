package app

import (
	"errors"
	"gocache/app/lru_list"
)

/*
   @Time: 2023/10/14 22:17
   @Author: david
   @File: lru
*/

// LRU fixed size LRU cache
// non-thread safe
type LRU[K comparable, V any] struct {
	size     int
	list     *lru_list.LruList[K, V]
	itemsMap map[K]*lru_list.Entry[K, V]
}

func NewLRU[K comparable, V any](size int) (*LRU[K, V], error) {
	if size <= 0 {
		return nil, errors.New("size invalid")
	}

	ins := &LRU[K, V]{
		size:     size,
		list:     lru_list.New[K, V](),
		itemsMap: make(map[K]*lru_list.Entry[K, V]),
	}
	return ins, nil
}

func (ins *LRU[K, V]) Get(key K) (value V, ok bool) {
	var ent *lru_list.Entry[K, V]
	if ent, ok = ins.itemsMap[key]; ok {
		ins.list.LMove(ent)
		return ent.Value, true
	}
	return
}

func (ins *LRU[K, V]) Set(key K, value V) {
	if ent, ok := ins.itemsMap[key]; ok {
		ins.list.LMove(ent)
		ent.Value = value
		return
	}

	ent := ins.list.LPush(key, value)
	ins.itemsMap[key] = ent
	if ins.full() {
		ins.remLeast()
	}
}

func (ins *LRU[K, V]) Remove(key K) (present bool) {
	if ent, ok := ins.itemsMap[key]; ok {
		ins.list.Remove(ent)
		delete(ins.itemsMap, ent.Key)
		return true
	}
	return false
}

func (ins *LRU[K, V]) remLeast() {
	if ent := ins.list.RPop(); ent != nil {
		delete(ins.itemsMap, ent.Key)
	}
}

func (ins *LRU[K, V]) full() bool {
	return ins.list.Len() > ins.size
}
