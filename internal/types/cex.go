package types

type CexName string

const (
	CexNameOkex CexName = "okex"
)

func (c CexName) String() string {
	return string(c)
}

// TradeMode define trade mode
type TradeMode string

// SideType define side type of orders
type SideType string

// PositionSideType define position side type of order
type PositionSideType string

// OrderType define order type
type OrderType string

// MarginMode define margin mode
type MarginMode string

// Global enums
const (
	TradeModeIsolated TradeMode = "isolated" //逐仓
	TradeModeCross    TradeMode = "cross"    //全仓
	TradeModeCash     TradeMode = "cash"     //现金

	MarginModeIsolated MarginMode = "isolated" //保证金模式（逐仓）
	MarginModeCross    MarginMode = "cross"    //保证模式（全仓）

	SideTypeBuy  SideType = "buy"  //买
	SideTypeSell SideType = "sell" //卖

	PositionSideTypeNet   PositionSideType = "net"
	PositionSideTypeLong  PositionSideType = "long"  //买多
	PositionSideTypeShort PositionSideType = "short" //买空

	// Order type
	OrderTypeLimit           OrderType = "limit"
	OrderTypeMarket          OrderType = "market"
	OrderTypePostOnly        OrderType = "post_only"
	OrderTypeFOK             OrderType = "fok"
	OrderTypeIOC             OrderType = "ioc"
	OrderTypeOptimalLimitIOC OrderType = "optimal_limit_ioc"

	// Algo order type
	OrderTypeConditional OrderType = "conditional"
	OrderTypeOCO         OrderType = "oco"
	OrderTypeTrigger     OrderType = "trigger"
	OrderTypeIceberg     OrderType = "iceberg"
	OrderTypeTwap        OrderType = "twap"
)

func (m TradeMode) String() string {
	return string(m)
}

func (m SideType) String() string {
	return string(m)
}

func (m OrderType) String() string {
	return string(m)
}

func (m MarginMode) String() string {
	return string(m)
}

func (m PositionSideType) String() string {
	return string(m)
}
