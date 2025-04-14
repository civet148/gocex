package okx

import (
	"context"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/jinzhu/copier"
)

func (m *CexOkex) GetOrder(symbols ...string) (orders []*types.OrderListDetail, err error) {
	svc := m.client.NewGetOrderListService()
	if len(symbols) != 0 {
		svc.InstrumentId(symbols[0])
	}
	res, err := svc.Do(context.Background())
	if err != nil {
		return nil, log.Errorf(err)
	}
	_ = copier.Copy(&orders, res.Data)
	return orders, nil
}
