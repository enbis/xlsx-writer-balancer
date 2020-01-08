package wb

import (
	"container/heap"
)

type Pool []*Work

type Balancer struct {
	pool Pool
	done chan *Work
}

func initBalancer() *Balancer {
	done := make(chan *Work, nWorker)
	b := &Balancer{make(Pool, 0, nWorker), done}
	for i := 0; i < nWorker; i++ {
		w := &Work{wok: make(chan Request)}
		heap.Push(&b.pool, w)
		go w.doWork(b.done)
	}
	return b
}

func (b *Balancer) balance(req chan Request) {

	for {
		select {

		case request := <-req:
			b.dispatch(request)

		case w := <-b.done:
			b.completed(w)
		}
	}
}

func (b *Balancer) dispatch(req Request) {

	w := heap.Pop(&b.pool).(*Work)
	w.wok <- req
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Work) {

	w.pending--
	heap.Remove(&b.pool, w.idx)
	heap.Push(&b.pool, w)

}
