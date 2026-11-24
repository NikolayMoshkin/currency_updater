package app

import (
	"context"
	cli_parser "currency_updater/internal/cli"
	"currency_updater/internal/config"
	"currency_updater/internal/service"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const DEFAULT_TICKER = 5 * time.Second // период опроса источников данных

type App struct {
	cnf *config.Config
}

func NewApp(cnf *config.Config) *App {
	return &App{cnf: cnf}
}

func (a *App) Run() {
	cli_parser.Parse()
	workers := cli_parser.CLI.Workers

	sysContext, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	sf := service.NewSourceChannelFiller(sysContext, &http.Client{})
	sourceCh := sf.InitSources(a.cnf.SourceConfig, DEFAULT_TICKER)

	mh := service.NewMessageHandler(sysContext)
	if err := mh.Init(a.cnf.DBConfig); err != nil {
		slog.Error("failed to connect to database", "error", err)
	}

	sc := service.NewSourceCaller(sysContext, workers, sourceCh)
	msgChan := sc.Call()

	mh.Handle(msgChan)
}
