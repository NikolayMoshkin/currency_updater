package db

import (
	"currency_updater/internal/source"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // "_" - вызывает только init() модуля и регистрирует "postgres" драйвер
)

type DataBaseSaver struct {
	db *sqlx.DB
}

func NewDBHandler(connStr string) (*DataBaseSaver, error) {
	if connStr == "" {
		return nil, errors.New("DB connection string not set")
	}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DataBaseSaver{db}, nil
}

func (h *DataBaseSaver) Init() error {
	err := h.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (s *DataBaseSaver) SaveCurrency(msg *source.Message) error {
	_, err := s.db.NamedExec("INSERT INTO currency (data, currency, source, created_at) VALUES (:data, :currency, :source, :created_at)", msg)
	return err
}
