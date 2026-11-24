package source

import "time"

type DummySource struct {
	name string
}

func (s *DummySource) GetUrl() string {
	return ""
}

func NewDummySource() *DummySource {
	return &DummySource{
		name: "dummy",
	}
}

func (s *DummySource) ParseCurrency() *Message {
	return &Message{
		Data:     1,
		Currency: "RUB",
		Date:     time.Now(),
		Source:   s.name,
	}
}
