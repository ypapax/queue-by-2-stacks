package main

import (
	"log"
	"sync"
)

type StackArr struct {
	slice []int
	mtx   *sync.Mutex
}

func (sa *StackArr) Push(i int) {
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	sa.slice = append(sa.slice, i)
	log.Println("len(sa.slice): ", len(sa.slice))

}

func (sa *StackArr) Pop() *int {
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	log.Println("len(sa.slice): ", len(sa.slice))
	if len(sa.slice) == 0 {
		return nil
	}
	r := sa.slice[len(sa.slice)-1]
	sa.slice = sa.slice[:len(sa.slice)-1]
	return &r
}
