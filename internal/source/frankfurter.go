package source

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type FrankfurterSource struct {
	name   string
	client *http.Client
	ctx    context.Context
}

func NewFrankfurterSource(ctx context.Context, c *http.Client) *FrankfurterSource {
	return &FrankfurterSource{
		name:   FRANKFURTER,
		client: c,
		ctx:    ctx,
	}
}

func (s *FrankfurterSource) ParseCurrency() *Message {
	r, d, err := s.getUSDRate()
	layout := "2006-01-02"
	date, err := time.Parse(layout, d)

	return &Message{
		Data:     r,
		Currency: "RUB",
		Date:     date,
		Source:   s.name,
		err:      err,
	}
}

const apiURL = "https://api.frankfurter.dev/v2/rate/usd/rub"

const REQUEST_TIMEOUT = 2 * time.Second

type RateResponse struct {
	Date  string  `json:"date"`
	Base  string  `json:"base"`
	Quote string  `json:"quote"`
	Rate  float64 `json:"rate"`
}

func (s *FrankfurterSource) getUSDRate() (float64, string, error) {
	ctx, stop := context.WithTimeout(s.ctx, REQUEST_TIMEOUT)
	defer stop()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return 0, "", fmt.Errorf("ошибка HTTP-запроса: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, "", fmt.Errorf("ошибка HTTP-запроса: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", fmt.Errorf("ошибка API: HTTP %d", resp.StatusCode)
	}

	var data RateResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, "", fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	return data.Rate, data.Date, nil
}
