package api

import "github.com/civet148/gocex/internal/types"

type CexUnimplInterface struct {
}

func (c *CexUnimplInterface) Name() string {
	return "CexUnimplInterface"
}

func (c *CexUnimplInterface) GetBalance(ccy string) (balance *types.Balance, err error) {
	return &types.Balance{}, nil
}

func (c *CexUnimplInterface) GetBalances(ccys ...string) (balances []*types.Balance, err error) {
	return balances, nil
}

func (c *CexUnimplInterface) GetTickerPrice(symbol string) (tickers []*types.TickerDetail, err error) {
	return tickers, nil
}

func (c *CexUnimplInterface) GetOrder(symbols ...string) (orders []*types.OrderListDetail, err error) {
	return orders, nil
}
