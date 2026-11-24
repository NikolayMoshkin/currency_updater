package service

import (
	"context"
	"currency_updater/internal/config"
	"currency_updater/internal/source"
	"net/http"
	"time"
)

type SourceChannelFiller struct {
	ctx    context.Context
	client *http.Client
}

func NewSourceChannelFiller(
	ctx context.Context,
	client *http.Client,
) *SourceChannelFiller {
	return &SourceChannelFiller{
		ctx:    ctx,
		client: client,
	}
}

func (sc *SourceChannelFiller) InitSources(cnf config.SourceConfig, tick time.Duration) <-chan source.Source {
	srcs := sc.init(cnf)
	ch := make(chan source.Source, len(srcs))
	go func() {
		ticker := time.NewTicker(tick)
	loop:
		for {
			for _, s := range srcs {
				ch <- s
			}
			select {
			case <-ticker.C: // запускаем опрос источников раз в период
				continue
			case <-sc.ctx.Done():
				break loop
			}
		}

		close(ch)
	}()

	return ch
}

func (sc *SourceChannelFiller) init(cnf config.SourceConfig) []source.Source {
	return []source.Source{
		source.NewDummySource(),
		source.NewErrorSource(),
		source.NewFrankfurterSource(sc.ctx, sc.client),
		source.NewWithCredSource(cnf.SourceWithCredsApiKey),
	}
}
