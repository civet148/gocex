package okx

import (
	"context"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/jinzhu/copier"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexOkex) GetBalance(ccy string) (balance *types.Balance, err error) {
	bs, err := m.GetBalances(ccy)
	if err != nil {
		return nil, err
	}
	return bs[0], nil
}

func (m *CexOkex) GetBalances(ccys ...string) (balances []*types.Balance, err error) {
	svc := m.client.NewGetBalanceService()
	var res *okex.GetBalanceServiceResponse
	if len(ccys) > 0 {
		svc.Currencies(ccys[0])
	}
	res, err = svc.Do(context.Background())
	if err != nil {
		return nil, log.Errorf(err)
	}
	for _, b := range res.Data {
		var balance = &types.Balance{}
		_ = copier.Copy(balance, b)
		balances = append(balances, balance)
	}
	return balances, nil
}
