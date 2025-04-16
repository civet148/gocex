package api

import (
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
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
	GetBalance(ccy string) (balance *types.Balance, err error)
	GetBalances(ccys ...string) (balances []*types.Balance, err error)
}

type SpotApi interface {
}

type ContractApi interface {
}

type MarketApi interface {
	GetTickerPrice(symbol string) ([]*types.TickerDetail, error)
}

type OrderApi interface {
	GetOrder(symbols ...string) (orders []*types.OrderListDetail, err error)
}
