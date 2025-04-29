package okx

import (
	"context"
	"fmt"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"github.com/jinzhu/copier"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexOkex) GetOrder(ctx context.Context, instIds ...string) (orders []*types.OrderListDetail, err error) {
	svc := m.client.NewGetOrderListService()
	if len(instIds) != 0 {
		svc.InstrumentId(instIds[0])
	}
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}

// PlaceOrder 下单接口：sz表示USDT的数量
func (m *CexOkex) PlaceOrder(ctx context.Context, sideType types.SideType, instId string, sz sqlca.Decimal, opts ...options.TradeOption) (orders []*types.OrderDetail, err error) {
	var tradeOpts = options.GetTradeConfig(opts...)
	svc := m.client.NewPlaceOrderService()

	if len(instId) == 0 {
		return nil, fmt.Errorf("instId requires")
	}
	if sz.LessThanOrEqual(0) {
		return nil, fmt.Errorf("sz invalid")
	}
	if tradeOpts.Swap != nil && *tradeOpts.Swap {
		instId = types.ToSwapInstId(m.Name(), instId) //构造合约交易对
	}
	if tradeOpts.Px != nil {
		svc.OrderPrice(*tradeOpts.Px)
	}
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
	svc.InstrumentId(instId).Side(okex.SideType(sideType)).Size(sz.String())
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}

func (m *CexOkex) GetPosition(ctx context.Context, instIds ...string) (orders []*types.OrderListDetail, err error) {
	svc := m.client.NewGetPositionsService()
	if len(instIds) != 0 {
		svc.InstrumentId(instIds[0])
	}
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}

func (m *CexOkex) OpenPosition(ctx context.Context, instId string, sideType types.SideType, sz sqlca.Decimal, opts ...options.TradeOption) (orders []*types.OrderDetail, err error) { //开仓
	var tradeOpts = options.GetTradeConfig(opts...)

	if tradeOpts.OrderType == nil {
		opts = append(opts, options.WithOrderType(types.OrderTypeMarket))
	}
	if tradeOpts.TradeMode == nil {
		opts = append(opts, options.WithTradeMode(types.TradeModeIsolated))
	}
	if tradeOpts.MgnMode == nil {
		opts = append(opts, options.WithMarginMode(types.MarginModeIsolated))
	}
	opts = append(opts, options.WithSwap())
	tradeOpts = options.GetTradeConfig(opts...)

	if tradeOpts.Leverage != nil {
		//TODO: compare and update leverage
	}

	return m.PlaceOrder(ctx, sideType, instId, sz, opts...)
}

func (m *CexOkex) ClosePosition(ctx context.Context, instId string, opts ...options.TradeOption) (orders []*types.ClosePositionDetail, err error) { //平仓
	var tradeOpts = options.GetTradeConfig(opts...)
	instId = types.ToSwapInstId(m.Name(), instId) //构造合约交易对
	svc := m.client.NewClosePositionService()
	if tradeOpts.MgnMode == nil {
		svc.MarginMode(string(types.MarginModeIsolated))
	} else {
		svc.MarginMode(string(*tradeOpts.MgnMode))
	}
	if tradeOpts.CliOrdId != nil {
		svc.CliOrderId(*tradeOpts.CliOrdId)
	}
	svc.InstrumentId(instId)
	res, err := svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}
