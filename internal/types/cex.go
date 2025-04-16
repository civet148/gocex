package types

type CexName string

const (
	CexNameOkex CexName = "okex"
)

func (c CexName) String() string {
	return string(c)
}
