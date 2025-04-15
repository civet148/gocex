package types

import "github.com/civet148/sqlca/v2"

type OrderListDetail struct {
	AccFillSz   sqlca.Decimal `json:"accFillSz"`
	AvgPx       sqlca.Decimal `json:"avgPx"`
	CTime       string        `json:"cTime"`
	Category    string        `json:"category"`
	Ccy         string        `json:"ccy"`
	ClOrdId     string        `json:"clOrdId"`
	Fee         sqlca.Decimal `json:"fee"`
	FeeCcy      string        `json:"feeCcy"`
	FillPx      sqlca.Decimal `json:"fillPx"`
	FillSz      sqlca.Decimal `json:"fillSz"`
	FillTime    string        `json:"fillTime"`
	InstId      string        `json:"instId"`
	InstType    string        `json:"instType"`
	Lever       sqlca.Decimal `json:"lever"`
	OrdId       string        `json:"ordId"`
	OrdType     string        `json:"ordType"`
	Pnl         string        `json:"pnl"`
	PosSide     string        `json:"posSide"`
	Px          sqlca.Decimal `json:"px"`
	Rebate      string        `json:"rebate"`
	RebateCcy   string        `json:"rebateCcy"`
	Side        string        `json:"side"`
	SlOrdPx     sqlca.Decimal `json:"slOrdPx"`
	SlTriggerPx sqlca.Decimal `json:"slTriggerPx"`
	State       string        `json:"state"`
	Sz          sqlca.Decimal `json:"sz"`
	Tag         string        `json:"tag"`
	TgtCcy      string        `json:"tgtCcy"`
	TdMode      string        `json:"tdMode"`
	TpOrdPx     sqlca.Decimal `json:"tpOrdPx"`
	TpTriggerPx sqlca.Decimal `json:"tpTriggerPx"`
	TradeId     string        `json:"tradeId"`
	UTime       int64         `json:"uTime"`
}
