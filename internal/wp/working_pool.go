package wp

import (
	"context"
	"currency_updater/internal/source"
	"sync"
)

type Worker struct {
	num      int
	jobsDone int
}

type WorkingPool struct {
	ctx     context.Context
	workers chan *Worker
	inCh    <-chan source.Source
}

func NewWorkingPool(ctx context.Context, workers int, in <-chan source.Source) *WorkingPool {
	ch := make(chan *Worker, workers)
	for i := range workers {
		ch <- &Worker{
			num:      i,
			jobsDone: 0,
		}
	}
	return &WorkingPool{
		ctx:     ctx,
		workers: ch,
		inCh:    in,
	}
}

func (wp *WorkingPool) Run(f func(ctx context.Context, s source.Source) *source.Message) <-chan *source.Message {
	outCh := make(chan *source.Message)

	wg := &sync.WaitGroup{}

	for range len(wp.workers) {
		worker := <-wp.workers
		wg.Add(1)
		go func() {
			defer func() {
				wp.workers <- worker
				wg.Done()
			}()
			for {
				select {
				case <-wp.ctx.Done():
					return
				case s, ok := <-wp.inCh:
					if !ok {
						return
					}
					outCh <- f(wp.ctx, s)
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}
