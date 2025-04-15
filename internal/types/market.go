package types

import "github.com/civet148/sqlca/v2"

type TickerDetail struct {
	InstType  string        `json:"instType"`
	InstId    string        `json:"instId"`
	Last      sqlca.Decimal `json:"last"`
	LastSz    sqlca.Decimal `json:"lastSz"`
	AskPx     sqlca.Decimal `json:"askPx"`
	AskSz     sqlca.Decimal `json:"askSz"`
	BidPx     sqlca.Decimal `json:"bidPx"`
	BidSz     sqlca.Decimal `json:"bidSz"`
	Open24h   sqlca.Decimal `json:"open24h"`
	High24h   sqlca.Decimal `json:"high24h"`
	Low24h    sqlca.Decimal `json:"low24h"`
	VolCcy24h sqlca.Decimal `json:"volCcy24h"`
	Vol24h    sqlca.Decimal `json:"vol24h"`
	SodUtc0   string        `json:"sodUtc0"`
	SodUtc8   string        `json:"sodUtc8"`
	Ts        int64         `json:"ts"`
}
