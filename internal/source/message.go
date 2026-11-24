package source

import "time"

type Message struct {
	Data     float64   `db:"data"`
	Currency string    `db:"currency"`
	Source   string    `db:"source"`
	Date     time.Time `db:"created_at"`
	err      error
}

func (m *Message) GetData() float64 {
	return m.Data
}

func (m *Message) GetCurrency() string {
	return m.Currency
}

func (m *Message) GetDate() time.Time {
	return m.Date
}

func (m *Message) GetSource() string {
	return m.Source
}

func (m *Message) GerError() error {
	return m.err
}
