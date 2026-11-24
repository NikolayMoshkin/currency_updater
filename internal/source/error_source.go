package source

import (
	"errors"
	"time"
)

type ErrorSource struct {
	name string
}

func NewErrorSource() *ErrorSource {
	return &ErrorSource{
		name: "Error",
	}
}

func (s *ErrorSource) ParseCurrency() *Message {
	return &Message{
		Date:   time.Now(),
		Source: s.name,
		err:    errors.New("ErrorSource error"),
	}
}
