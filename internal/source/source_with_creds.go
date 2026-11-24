package source

import "time"

type WithCredSource struct {
	cred string
	name string
}

func (s *WithCredSource) GetUrl() string {
	return ""
}

func NewWithCredSource(cred string) *WithCredSource {
	return &WithCredSource{
		cred: cred,
		name: "credit_source",
	}
}

func (s *WithCredSource) ParseCurrency() *Message {
	return &Message{
		Data:     666,
		Currency: "RUB",
		Date:     time.Now(),
		Source:   s.name,
	}
}
