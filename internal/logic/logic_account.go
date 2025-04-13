package logic

import (
	"context"
	"github.com/civet148/log"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexLogic) GetBalance(ccys ...string) (balance []*okex.Balance, err error) {
	svc := m.client.NewGetBalanceService()
	var res *okex.GetBalanceServiceResponse
	if len(ccys) > 0 {
		svc.Currencies(ccys[0])
	}
	res, err = svc.Do(context.Background())
	if err != nil {
		return nil, log.Errorf(err)
	}
	return res.Data, nil
}
