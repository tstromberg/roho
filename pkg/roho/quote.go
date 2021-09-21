package roho

import (
	"context"
	"strings"

	"github.com/tstromberg/roho/pkg/times"
)

// A Quote is a representation of the data returned by the Robinhood API for
// current stock quotes.
type Quote struct {
	AdjustedPreviousClose       float64 `json:"adjusted_previous_close,string"`
	AskPrice                    float64 `json:"ask_price,string"`
	AskSize                     int     `json:"ask_size"`
	BidPrice                    float64 `json:"bid_price,string"`
	BidSize                     int     `json:"bid_size"`
	LastExtendedHoursTradePrice float64 `json:"last_extended_hours_trade_price,string"`
	LastTradePrice              float64 `json:"last_trade_price,string"`
	PreviousClose               float64 `json:"previous_close,string"`
	PreviousCloseDate           string  `json:"previous_close_date"`
	Symbol                      string  `json:"symbol"`
	TradingHalted               bool    `json:"trading_halted"`
	UpdatedAt                   string  `json:"updated_at"`
}

// Quote returns the latest stock quotes for the list of stocks provided.
func (c *Client) Quote(ctx context.Context, symbols ...string) ([]Quote, error) {
	url := baseURL("quotes") + "?symbols=" + strings.Join(symbols, ",")
	var r struct{ Results []Quote }
	err := c.get(ctx, url, &r)
	return r.Results, err
}

// Price returns the proper stock price even after hours.
func (q Quote) Price() float64 {
	if times.IsRegularTradingTime() {
		return q.LastTradePrice
	}
	return q.LastExtendedHoursTradePrice
}
