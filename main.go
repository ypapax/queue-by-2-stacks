package main

import (
	"log"
	"sync"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	stack := &StackArr{mtx: &sync.RWMutex{}}
	playAround(stack, 5)
	query := &QueueByTwoStacks{mtx: &sync.RWMutex{}, Stack1: &StackArr{mtx: &sync.RWMutex{}}, Stack2: &StackArr{mtx: &sync.RWMutex{}}}
	playAround(query, 10)
}

func playAround(s pushPopPeek, n int) {
	for i := 1; i <= n; i++ {
		s.Push(i)
	}
	var counter int
	for {
		v := s.Pop()
		if v == nil {
			break
		}
		log.Printf("%T, counter: %+v, result:----%+v\n", s, counter, *v)
		counter++
	}
	/*for v := s.Pop();v != nil; {
		log.Println(*v)
	}*/
	//for i := 1; i <= n; i++ {
	//	v := s.Pop()
	//	if v != nil {
	//		log.Println(*v)
	//	} else {
	//		log.Println("nil")
	//	}
	//}
}

type pushPopPeek interface {
	Push(int)
	Pop() *int
	Peek() *int
}
