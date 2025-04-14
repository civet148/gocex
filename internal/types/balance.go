package types

import "github.com/civet148/sqlca/v2"

// Balance define account info
type BalanceDetail struct {
	AvailBal      sqlca.Decimal `json:"availBal"`
	AvailEq       sqlca.Decimal `json:"availEq"`
	CashBal       sqlca.Decimal `json:"cashBal"`
	Ccy           string        `json:"ccy"`
	CrossLiab     string        `json:"crossLiab"`
	DisEq         sqlca.Decimal `json:"disEq"`
	Eq            sqlca.Decimal `json:"eq"`
	EqUsd         sqlca.Decimal `json:"eqUsd"`
	FrozenBal     sqlca.Decimal `json:"frozenBal"`
	Interest      sqlca.Decimal `json:"interest"`
	IsoEq         sqlca.Decimal `json:"isoEq"`
	IsoLiab       string        `json:"isoLiab"`
	Liab          string        `json:"liab"`
	MaxLoan       sqlca.Decimal `json:"maxLoan"`
	MgnRatio      sqlca.Decimal `json:"mgnRatio"`
	NotionalLever string        `json:"notionalLever"`
	OrdFrozen     sqlca.Decimal `json:"ordFrozen"`
	Twap          string        `json:"twap"`
	UTime         string        `json:"uTime"`
	Upl           string        `json:"upl"`
	PplLiab       string        `json:"uplLiab"`
	StgyEq        sqlca.Decimal `json:"stgyEq"`
}

type Balance struct {
	AdjEq       sqlca.Decimal    `json:"adjEq"`
	Details     []*BalanceDetail `json:"details"`
	Imr         string           `json:"imr"`
	IsoEq       sqlca.Decimal    `json:"isoEq"`
	MgnRatio    sqlca.Decimal    `json:"mgnRatio"`
	Mnr         string           `json:"mnr"`
	NotionalUsd string           `json:"notionalUsd"`
	OrdFroz     sqlca.Decimal    `json:"ordFroz"`
	TotalEq     sqlca.Decimal    `json:"totalEq"`
	UTime       string           `json:"uTime"`
}
