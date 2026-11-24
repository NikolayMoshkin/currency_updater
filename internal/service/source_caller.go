package service

import (
	"context"
	"currency_updater/internal/source"
	"currency_updater/internal/wp"
)

type SourceCaller struct {
	ctx      context.Context
	workers  int
	sourceCh <-chan source.Source
}

func NewSourceCaller(ctx context.Context, workers int, sourceCh <-chan source.Source) *SourceCaller {
	return &SourceCaller{
		ctx:      ctx,
		workers:  workers,
		sourceCh: sourceCh,
	}
}

func (sc *SourceCaller) Call() <-chan *source.Message {
	wp := wp.NewWorkingPool(sc.ctx, sc.workers, sc.sourceCh)

	return wp.Run(func(ctx context.Context, s source.Source) *source.Message {
		return s.ParseCurrency()
	})
}
