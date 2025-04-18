package api

import "github.com/civet148/gocex/internal/types"

type CexUnimplement struct {
}

func (c *CexUnimplement) Name() string {
	return "CexUnimplement"
}

func (c *CexUnimplement) GetBalance(ccy string) (balance *types.Balance, err error) {
	return &types.Balance{}, nil
}

func (c *CexUnimplement) GetBalances(ccys ...string) (balances []*types.Balance, err error) {
	return balances, nil
}

func (c *CexUnimplement) GetTickerPrice(symbol string) (tickers []*types.TickerDetail, err error) {
	return tickers, nil
}

func (c *CexUnimplement) GetOrder(symbols ...string) (orders []*types.OrderListDetail, err error) {
	return orders, nil
}
