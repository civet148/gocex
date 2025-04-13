package logic

import (
	"context"
	"github.com/civet148/log"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexLogic) GetOrders(instId string) ([]*okex.OrderListDetail, error) {
	svc := m.client.NewGetOrderListService()
	if instId != "" {
		svc.InstrumentId(instId)
	}
	res, err := svc.Do(context.Background())
	if err != nil {
		return nil, log.Errorf(err)
	}
	return res.Data, nil
}
