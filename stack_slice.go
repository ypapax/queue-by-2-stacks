package main

import (
	"log"
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
	log.Println("len(sa.slice): ", len(sa.slice))

}

func (sa *StackArr) Pop() *int {
	r := sa.Peek()
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	log.Println("0: in pop len(sa.slice): ", len(sa.slice))
	if len(sa.slice) == 0 {
		return nil
	}
	sa.slice = sa.slice[:len(sa.slice)-1]
	log.Println("in pop len(sa.slice): ", len(sa.slice))
	return r
}

func (sa *StackArr) Peek() *int {
	sa.mtx.RLock()
	defer sa.mtx.RUnlock()
	log.Println("len(sa.slice): ", len(sa.slice))
	if len(sa.slice) == 0 {
		return nil
	}
	r := sa.slice[len(sa.slice)-1]
	return &r
}
