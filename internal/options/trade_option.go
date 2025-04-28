package options

import "github.com/civet148/gocex/internal/types"

type TradeConfig struct {
	Leverage     *string                 //杠杆倍数
	Swap         *bool                   //是否为合约
	Px           *string                 //价格
	OrderType    *types.OrderType        //订单类型(限价/市价/FOK...)
	PositionSide *types.PositionSideType //持仓方向(永续/交割...)
	TradeMode    *types.TradeMode        //交易模式(逐仓/全仓/现金)
	MgnMode      *types.MarginMode       //保证金模式(逐仓/全仓)
	CliOrdId     *string                 //客户自定义订单ID
}

type TradeOption func(o *TradeConfig)

func GetTradeConfig(options ...TradeOption) *TradeConfig {
	var tradeOption TradeConfig
	for _, o := range options {
		o(&tradeOption)
	}
	return &tradeOption
}

func WithPrice(px string) TradeOption {
	return func(c *TradeConfig) {
		c.Px = &px
	}
}

func WithLever(leverage string) TradeOption {
	return func(c *TradeConfig) {
		c.Leverage = &leverage
	}
}

func WithOrderType(orderType types.OrderType) TradeOption {
	return func(c *TradeConfig) {
		c.OrderType = &orderType
	}
}

func WithPositionSide(posType types.PositionSideType) TradeOption {
	return func(c *TradeConfig) {
		c.PositionSide = &posType
	}
}

func WithTradeMode(tradeMode types.TradeMode) TradeOption {
	return func(c *TradeConfig) {
		c.TradeMode = &tradeMode
	}
}

func WithSwap() TradeOption {
	return func(c *TradeConfig) {
		var yes = true
		c.Swap = &yes
	}
}

func WithMarginMode(mgnMode types.MarginMode) TradeOption {
	return func(c *TradeConfig) {
		c.MgnMode = &mgnMode
	}
}

func WithCliOrdId(cliOrdId string) TradeOption {
	return func(c *TradeConfig) {
		c.CliOrdId = &cliOrdId
	}
}
