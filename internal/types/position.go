package types

import (
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"strings"
)

func ToSwapInstId(cexName string, symbol string) string {
	symbol = strings.ToUpper(symbol)
	switch CexName(cexName) {
	case CexNameOkex:
		if !strings.Contains("SWAP", symbol) {
			return fmt.Sprintf("%s-SWAP", symbol)
		}
	default:
		log.Panic("cex %s not support yet", cexName)
	}
	return symbol
}

// Position detail inside the positions
type PositionDetail struct {
	Adl         string        `json:"adl"`
	AvailPos    string        `json:"availPos"`
	AvgPx       sqlca.Decimal `json:"avgPx"`
	CTime       string        `json:"cTime"`
	Ccy         string        `json:"ccy"` //
	DeltaBS     string        `json:"deltaBS"`
	DeltaPA     string        `json:"deltaPA"`
	GammaBS     string        `json:"gammaBS"`
	GammaPA     string        `json:"gammaPA"`
	Imr         string        `json:"imr"`
	InstId      string        `json:"instId"`   //例如：PEPE-USDT-SWAP
	InstType    string        `json:"instType"` //例如: SWAP
	Interest    string        `json:"interest"`
	Last        sqlca.Decimal `json:"last"`
	Lever       sqlca.Decimal `json:"lever"`
	Liab        string        `json:"liab"`
	LiabCcy     string        `json:"liabCcy"`
	LiqPx       sqlca.Decimal `json:"liqPx"`
	Margin      string        `json:"margin"`
	MgnMode     string        `json:"mgnMode"`
	MgnRatio    sqlca.Decimal `json:"mgnRatio"`
	Mmr         string        `json:"mmr"`
	NotionalCcy string        `json:"notionalCcy"` // this is for AccountAndPositionRisk
	NotionalUsd sqlca.Decimal `json:"notionalUsd"`
	OptVal      string        `json:"optVal"`
	PTime       string        `json:"pTime"`
	Pos         sqlca.Decimal `json:"pos"`
	PosCcy      string        `json:"posCcy"`
	PosId       string        `json:"posId"`
	PosSide     string        `json:"posSide"`
	ThetaBS     string        `json:"thetaBS"`
	TradeId     string        `json:"tradeId"`
	UTime       string        `json:"uTime"`
	Upl         sqlca.Decimal `json:"upl"`
	UplRatio    sqlca.Decimal `json:"uplRatio"`
	VegaBS      string        `json:"vegaBS"`
	VegaPA      string        `json:"vegaPA"`
}
