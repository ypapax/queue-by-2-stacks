package main

import (
	"sync"
)

type StackArr struct {
	slice []int
	mtx   *sync.RWMutex
}

func (sa *StackArr) Push(i int) {
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	sa.slice = append(sa.slice, i)
}

func (sa *StackArr) Pop() *int {
	r := sa.Peek()
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	if len(sa.slice) == 0 {
		return nil
	}
	sa.slice = sa.slice[:len(sa.slice)-1]
	return r
}

func (sa *StackArr) Peek() *int {
	sa.mtx.RLock()
	defer sa.mtx.RUnlock()
	if len(sa.slice) == 0 {
		return nil
	}
	r := sa.slice[len(sa.slice)-1]
	return &r
}
