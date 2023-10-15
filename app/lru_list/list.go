package lru_list

/*
   @Time: 2023/10/14 18:23
   @Author: david
   @File: list
*/

type LruList[K comparable, V any] struct {
	root Entry[K, V]
	len  int
}

func New[K comparable, V any]() *LruList[K, V] {
	list := new(LruList[K, V])
	list.Init()
	return list
}

func (ins *LruList[K, V]) Init() {
	ins.root.root = true
	ins.root.prev = &ins.root
	ins.root.next = &ins.root
}

func (ins *LruList[K, V]) Len() int {
	return ins.len
}

func (ins *LruList[K, V]) LPush(k K, v V) *Entry[K, V] {
	ent := &Entry[K, V]{
		Key:   k,
		Value: v,
	}
	ins.insert(ent, &ins.root)
	return ent
}

func (ins *LruList[K, V]) LPop() *Entry[K, V] {
	if ins.isEmpty() {
		return nil
	}
	ent := ins.root.next
	ins.remove(ent)
	return ent
}

func (ins *LruList[K, V]) Remove(ent *Entry[K, V]) V {
	ins.remove(ent)
	return ent.Value
}

func (ins *LruList[K, V]) LMove(ent *Entry[K, V]) {
	if ins.root.next == ent {
		return
	}

	ins.move(ent, &ins.root)
}

// RPop returns the last element of list l or nil if the list is empty.
func (ins *LruList[K, V]) RPop() *Entry[K, V] {
	if ins.isEmpty() {
		return nil
	}
	ent := ins.root.prev
	ins.remove(ent)
	return ent
}

func (ins *LruList[K, V]) RRange(num int, iter func(k K, v V)) {
	var i int
	for b := ins.root.prev; b != nil && i < num; b = b.Prev() {
		iter(b.Key, b.Value)
	}
}

func (ins *LruList[K, V]) insert(ent, at *Entry[K, V]) {
	ent.prev = at
	ent.next = at.next
	ent.prev.next = ent
	ent.next.prev = ent
	ins.len++
}

func (ins *LruList[K, V]) move(ent, at *Entry[K, V]) {
	if ent == at {
		return
	}

	ent.prev.next = ent.next
	ent.next.prev = ent.prev

	ent.prev = at
	ent.next = at.next
	ent.prev.next = ent
	ent.next.prev = ent
}

func (ins *LruList[K, V]) remove(ent *Entry[K, V]) {
	ent.prev.next = ent.next
	ent.next.prev = ent.prev

	ent.clear()
	ins.len--
}

func (ins *LruList[K, V]) isEmpty() bool {
	return ins.len == 0
}
