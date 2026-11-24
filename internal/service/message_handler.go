package service

import (
	"context"
	"currency_updater/internal/config"
	"currency_updater/internal/db"
	"currency_updater/internal/source"
	"fmt"
	"log/slog"
)

type MessageHandler struct {
	db  *db.DataBaseSaver
	ctx context.Context
}

func NewMessageHandler(ctx context.Context) *MessageHandler {
	return &MessageHandler{
		ctx: ctx,
	}
}

func (h *MessageHandler) Init(cnf config.DBConfig) error {

	conStr := fmt.Sprintf("postgres://postgres:%s@%s:%d/%s?sslmode=disable", cnf.Password, cnf.Host, cnf.Port, cnf.Name)
	dbs, err := db.NewDBHandler(conStr)

	if err != nil {
		return err
	}
	if err := dbs.Init(); err != nil {
		return err
	}

	h.db = dbs

	return nil
}

func (h *MessageHandler) Handle(ch <-chan *source.Message) {
loop:
	for {
		select {
		case <-h.ctx.Done():
			break loop
		case mes, ok := <-ch:
			if !ok {
				break loop
			}
			if mes.GerError() != nil {
				slog.Error(mes.GerError().Error(), "source", mes.Source)
				continue
			}

			if err := h.db.SaveCurrency(mes); err != nil {
				slog.Error("failed to save currency data", "error", err, "source", mes.Source)
			}
		}
	}
}
