package okx

import (
	"context"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/jinzhu/copier"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexOkex) GetBalance(ctx context.Context, ccy string) (balance *types.Balance, err error) {
	bs, err := m.GetBalances(ctx, ccy)
	if err != nil {
		return nil, err
	}
	return bs[0], nil
}

func (m *CexOkex) GetBalances(ctx context.Context, ccys ...string) (balances []*types.Balance, err error) {
	svc := m.client.NewGetBalanceService()
	var res *okex.GetBalanceServiceResponse
	if len(ccys) > 0 {
		svc.Currencies(ccys[0])
	}
	res, err = svc.Do(context.Background())
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&balances, res.Data)
	return balances, nil
}

func (m *CexOkex) GetLeverage(ctx context.Context, instId string, mgnMode types.MarginMode, opts ...options.TradeOption) (leverages []*types.LeverageDetail, err error) {
	var tradeOpts = options.GetTradeConfig(opts...)
	if tradeOpts.Swap != nil && *tradeOpts.Swap {
		instId = types.ToSwapInstId(m.Name(), instId) //构造合约交易对
	}

	svc := m.client.NewGetLeverageService()
	var res *okex.GetLeverageServiceResponse
	res, err = svc.InstrumentId(instId).MarginMode(string(mgnMode)).Do(ctx)
	if err != nil {
		return nil, log.Errorf("get leverage error: %s", err)
	}
	_ = copier.Copy(&leverages, res.Data)
	return leverages, nil
}

func (m *CexOkex) SetLeverage(ctx context.Context, instId string, mgnMode types.MarginMode, opts ...options.TradeOption) (leverages []*types.LeverageDetail, err error) {
	var tradeOpts = options.GetTradeConfig(opts...)
	if tradeOpts.Swap != nil && *tradeOpts.Swap {
		instId = types.ToSwapInstId(m.Name(), instId) //构造合约交易对
	}
	svc := m.client.NewSetLeverageService()
	svc.InstrumentId(instId)
	if tradeOpts.Leverage != nil {
		svc.Leverage(*tradeOpts.Leverage)
	}
	svc.MarginMode(mgnMode.String())
	var res *okex.SetLeverageServiceResponse
	res, err = svc.Do(ctx)
	if err != nil {
		return nil, log.Errorf(err.Error())
	}
	_ = copier.Copy(&leverages, res.Data)
	return leverages, nil
}
