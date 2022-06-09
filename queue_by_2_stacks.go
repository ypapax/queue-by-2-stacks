package main

import (
	"log"
	"sync"
)

type QueueByTwoStacks struct {
	Stack1 pushPopPeek
	Stack2 pushPopPeek
	mtx *sync.RWMutex
}

func (sa *QueueByTwoStacks) Push(i int) {
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	notEmpty, _ := sa.notEmptyAndEmptyOnes()
	notEmpty.Push(i)

}

func (sa *QueueByTwoStacks) notEmptyAndEmptyOnes() (pushPopPeek, pushPopPeek) {
	var full, empty pushPopPeek
	if sa.Stack1.Peek() != nil {
		log.Println("first not empty")
		full = sa.Stack1
		empty = sa.Stack2
	} else {
		log.Println("second not empty")
		full = sa.Stack2
		empty = sa.Stack1
	}
	return full, empty
}

func (sa *QueueByTwoStacks) Peek() *int {
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	notEmpty, empty := sa.notEmptyAndEmptyOnes()

	move(notEmpty, empty)
	defer move(empty, notEmpty)
	return empty.Peek()
}

func (sa *QueueByTwoStacks) Pop() *int {
	sa.mtx.Lock()
	defer sa.mtx.Unlock()
	notEmpty, empty := sa.notEmptyAndEmptyOnes()

	move(notEmpty, empty)
	defer move(empty, notEmpty)
	return empty.Pop()
}

func move(notEmpty, empty pushPopPeek) {
	for {
		v := notEmpty.Pop()
		if v == nil {
			break
		}
		empty.Push(*v)
		log.Println("pushed", *v)
	}
}