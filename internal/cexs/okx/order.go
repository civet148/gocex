package okx

import (
	"context"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"github.com/jinzhu/copier"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexOkex) GetOrder(ctx context.Context, symbols ...string) (orders []*types.OrderListDetail, err error) {
	svc := m.client.NewGetOrderListService()
	if len(symbols) != 0 {
		svc.InstrumentId(symbols[0])
	}
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}

func (m *CexOkex) PlaceOrder(ctx context.Context, side types.SideType, symbol string, px, sz sqlca.Decimal, opts ...options.TradeOption) (orders []*types.OrderDetail, err error) {
	var tradeOpts = options.GetTradeConfig(opts...)
	svc := m.client.NewPlaceOrderService()

	if tradeOpts.Swap != nil && *tradeOpts.Swap {
		symbol = types.ToSwapInstId(m.Name(), symbol) //构造合约交易对
	}

	svc.InstrumentId(symbol).Side(okex.SideType(side)).OrderPrice(px.String()).Size(sz.String())
	if tradeOpts.OrderType != nil {
		svc.OrderType(okex.OrderType(*tradeOpts.OrderType))
	}
	if tradeOpts.Leverage != nil {
		svc.Leverage(*tradeOpts.Leverage)
	}
	if tradeOpts.TradeMode != nil {
		svc.TradeMode(okex.TradeMode(*tradeOpts.TradeMode))
	}
	if tradeOpts.PositionSide != nil {
		svc.PositionSide(okex.PositionSideType(*tradeOpts.PositionSide))
	}
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}

func (m *CexOkex) GetPosition(ctx context.Context, symbols ...string) (orders []*types.OrderListDetail, err error) {
	svc := m.client.NewGetPositionsService()
	if len(symbols) != 0 {
		svc.InstrumentId(symbols[0])
	}
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}

func (m *CexOkex) OpenPosition(ctx context.Context, symbol string, px, sz sqlca.Decimal, opts ...options.TradeOption) (orders []*types.OrderDetail, err error) { //开仓
	var tradeOpts = options.GetTradeConfig(opts...)
	symbol = types.ToSwapInstId(m.Name(), symbol) //构造合约交易对

	if tradeOpts.OrderType == nil {
		opts = append(opts, options.WithOrderType(types.OrderTypeMarket))
	}
	if px.GreaterThan(0) {
		opts = append(opts, options.WithOrderType(types.OrderTypeLimit))
	}
	if tradeOpts.Leverage == nil {
		opts = append(opts, options.WithLever("1"))
	}
	if tradeOpts.TradeMode == nil {
		opts = append(opts, options.WithTradeMode(types.TradeModeIsolated))
	}
	//if tradeOpts.PositionSide == nil {//不能设置默认（会提示参数错误）
	//	opts = append(opts, options.WithPositionSide(types.PositionSideTypeLong))
	//}
	log.Json("trade options", tradeOpts)
	return m.PlaceOrder(ctx, types.SideTypeBuy, symbol, px, sz, opts...)
}

func (m *CexOkex) ClosePosition(ctx context.Context, symbol string, opts ...options.TradeOption) (orders []*types.ClosePositionDetail, err error) { //平仓
	var tradeOpts = options.GetTradeConfig(opts...)
	symbol = types.ToSwapInstId(m.Name(), symbol) //构造合约交易对
	svc := m.client.NewClosePositionService()
	if tradeOpts.MgnMode == nil {
		svc.MarginMode(string(types.MarginModeIsolated))
	} else {
		svc.MarginMode(string(*tradeOpts.MgnMode))
	}
	if tradeOpts.CliOrdId != nil {
		svc.CliOrderId(*tradeOpts.CliOrdId)
	}
	svc.InstrumentId(symbol)
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}
