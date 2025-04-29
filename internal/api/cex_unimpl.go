package api

import (
	"context"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/sqlca/v2"
)

const ()

type CexUnimplement struct {
}

func (m *CexUnimplement) Name() string {
	return "CexUnimplement"
}

func (m *CexUnimplement) GetBalance(ctx context.Context, ccy string) (balance *types.Balance, err error) {
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) GetBalances(ctx context.Context, ccys ...string) (balances []*types.Balance, err error) {
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) GetTickerPrice(ctx context.Context, instId string) (tickers []*types.TickerDetail, err error) {
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) GetOrder(ctx context.Context, instIds ...string) (orders []*types.OrderListDetail, err error) {
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) PlaceOrder(ctx context.Context, sideType types.SideType, instId string, sz sqlca.Decimal, options ...options.TradeOption) (orders []*types.OrderDetail, err error) {
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) GetPosition(ctx context.Context, instIds ...string) (orders []*types.OrderListDetail, err error) {
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) OpenPosition(ctx context.Context, instId string, sideType types.SideType, sz sqlca.Decimal, options ...options.TradeOption) (orders []*types.OrderDetail, err error) { //开仓
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) ClosePosition(ctx context.Context, instId string, opts ...options.TradeOption) (orders []*types.ClosePositionDetail, err error) { //平仓
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) GetLeverage(ctx context.Context, instId string, mgnMode types.MarginMode, opts ...options.TradeOption) (leverages []*types.LeverageDetail, err error) {
	return nil, types.ErrorNotSupport
}

func (m *CexUnimplement) SetLeverage(ctx context.Context, instId string, mgnMode types.MarginMode, opts ...options.TradeOption) (leverages []*types.LeverageDetail, err error) {
	return nil, types.ErrorNotSupport
}
