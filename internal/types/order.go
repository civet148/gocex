package types

import "github.com/civet148/sqlca/v2"

type OrderListDetail struct {
	Adl                    string        `json:"adl"`
	AvailPos               string        `json:"availPos"`
	AvgPx                  sqlca.Decimal `json:"avgPx"` //开仓均价
	BaseBal                sqlca.Decimal `json:"baseBal"`
	BaseBorrowed           sqlca.Decimal `json:"baseBorrowed"`
	BaseInterest           sqlca.Decimal `json:"baseInterest"`
	BePx                   sqlca.Decimal `json:"bePx"`
	BizRefId               string        `json:"bizRefId"`
	BizRefType             string        `json:"bizRefType"`
	CTime                  string        `json:"cTime"`
	Ccy                    string        `json:"ccy"` //基础代币（例如：USDT）
	ClSpotInUseAmt         sqlca.Decimal `json:"clSpotInUseAmt"`
	DeltaBS                string        `json:"deltaBS"`
	DeltaPA                string        `json:"deltaPA"`
	Fee                    sqlca.Decimal `json:"fee"`
	FundingFee             sqlca.Decimal `json:"fundingFee"`
	GammaBS                string        `json:"gammaBS"`
	GammaPA                string        `json:"gammaPA"`
	IdxPx                  sqlca.Decimal `json:"idxPx"`
	Imr                    string        `json:"imr"`
	InstId                 string        `json:"instId"`   //产品ID（例如：PEPE-USDT-SWAP)
	InstType               string        `json:"instType"` //产品类型（例如：SWAP)
	Interest               string        `json:"interest"`
	Last                   sqlca.Decimal `json:"last"`  //最新价格
	Lever                  sqlca.Decimal `json:"lever"` //杠杆倍数
	Liab                   string        `json:"liab"`
	LiabCcy                string        `json:"liabCcy"`
	LiqPenalty             sqlca.Decimal `json:"liqPenalty"`
	LiqPx                  sqlca.Decimal `json:"liqPx"`
	Margin                 string        `json:"margin"` //保证金(单位：USDT)
	MarkPx                 sqlca.Decimal `json:"markPx"`
	MaxSpotInUseAmt        string        `json:"maxSpotInUseAmt"`
	MgnMode                string        `json:"mgnMode"`  //保证金模式（逐仓/全仓）
	MgnRatio               sqlca.Decimal `json:"mgnRatio"` //保证金比例
	Mmr                    string        `json:"mmr"`
	NonSettleAvgPx         sqlca.Decimal `json:"nonSettleAvgPx"`
	NotionalUsd            sqlca.Decimal `json:"notionalUsd"` //持仓量（单位：USDT）
	OptVal                 string        `json:"optVal"`
	PendingCloseOrdLiabVal string        `json:"pendingCloseOrdLiabVal"`
	Pnl                    string        `json:"pnl"`
	Pos                    string        `json:"pos"`
	PosCcy                 string        `json:"posCcy"`
	PosId                  string        `json:"posId"`
	PosSide                string        `json:"posSide"`
	QuoteBal               sqlca.Decimal `json:"quoteBal"`
	QuoteBorrowed          sqlca.Decimal `json:"quoteBorrowed"`
	QuoteInterest          sqlca.Decimal `json:"quoteInterest"`
	RealizedPnl            string        `json:"realizedPnl"`
	SettledPnl             string        `json:"settledPnl"`
	SpotInUseAmt           sqlca.Decimal `json:"spotInUseAmt"`
	SpotInUseCcy           string        `json:"spotInUseCcy"`
	ThetaBS                string        `json:"thetaBS"`
	ThetaPA                string        `json:"thetaPA"`
	TradeId                string        `json:"tradeId"`
	UTime                  string        `json:"uTime"`
	Upl                    sqlca.Decimal `json:"upl"`            //收益额
	UplLastPx              sqlca.Decimal `json:"uplLastPx"`      //空
	UplRatio               sqlca.Decimal `json:"uplRatio"`       //收益率
	UplRatioLastPx         sqlca.Decimal `json:"uplRatioLastPx"` //空
	UsdPx                  sqlca.Decimal `json:"usdPx"`          //空
	//VegaBS                 string        `json:"vegaBS"`
	//VegaPA                 string        `json:"vegaPA"`
	//CloseOrderAlgo         []interface{} `json:"closeOrderAlgo"`
}

type OrderDetail struct {
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
	ReqId   string `json:"reqId"`
}

type ClosePositionDetail struct {
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
