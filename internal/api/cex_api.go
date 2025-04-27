package api

import (
	"context"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type CexApi interface {
	CommonApi
	AccountApi
	SpotApi
	ContractApi
	MarketApi
	OrderApi
}

type CexNew func(config *config.Config) CexApi

var cexInsts = make(map[types.CexName]CexNew)

func RegisterCex(ct types.CexName, inst CexNew) {
	cexInsts[ct] = inst
}

func NewCex(ct types.CexName, c *config.Config) CexApi {
	inst, ok := cexInsts[ct]
	if !ok {
		log.Panic("cex type %v not registered", ct)
	}
	return inst(c)
}

type CommonApi interface {
	Name() string
}

type AccountApi interface {
	GetBalance(ctx context.Context, ccy string) (balance *types.Balance, err error)
	GetBalances(ctx context.Context, ccys ...string) (balances []*types.Balance, err error)
}

type SpotApi interface {
}

type ContractApi interface {
}

type MarketApi interface {
	GetTickerPrice(ctx context.Context, symbol string) ([]*types.TickerDetail, error)
}

type OrderApi interface {
	GetOrder(ctx context.Context, symbols ...string) (orders []*types.OrderListDetail, err error)                                                                      //订单列表
	PlaceOrder(ctx context.Context, side types.SideType, symbol string, px, sz sqlca.Decimal, options ...options.TradeOption) (orders []*types.OrderDetail, err error) //订单下单
	GetPosition(ctx context.Context, symbols ...string) (orders []*types.OrderListDetail, err error)                                                                   //仓位
	OpenPosition(ctx context.Context, symbol string, px, sz sqlca.Decimal, options ...options.TradeOption) (orders []*types.OrderDetail, err error)                    //开仓
	ClosePosition(ctx context.Context, symbol string, opts ...options.TradeOption) (orders []*types.ClosePositionDetail, err error)                                    //平仓
}
