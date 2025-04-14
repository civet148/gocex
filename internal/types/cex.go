package types

type CexType string

const (
	CexTypeOkex    CexType = "okex"
	CexTypeBinance CexType = "binance"
)

func (c CexType) String() string {
	return string(c)
}
