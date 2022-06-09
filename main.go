package main

import (
	"log"
	"sync"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := &StackArr{mtx: &sync.Mutex{}}
	playAroundWithStack(s, 5)
}

func playAroundWithStack(s stack, n int) {
	for i := 1; i <= n; i++ {
		s.Push(i)
	}
	for i := 1; i <= n; i++ {
		v := s.Pop()
		if v != nil {
			log.Println(*v)
		} else {
			log.Println("nil")
		}
	}
}

type stack interface {
	Push(int)
	Pop() *int
}

type queue interface {
	Push(int)
	Pop() *int
}
