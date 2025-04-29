package okx

import (
	"context"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/jinzhu/copier"
)

func (m *CexOkex) GetTickerPrice(ctx context.Context, instId string) (tickers []*types.TickerDetail, err error) {
	svc := m.client.NewGetTickerService()
	if instId == "" {
		return nil, log.Errorf("instId is empty")
	}
	res, err := svc.InstrumentId(instId).Do(ctx)
	if err != nil {
		return nil, log.Errorf("get ticker error: %s", err)
	}
	_ = copier.Copy(&tickers, &res.Data)
	return tickers, nil
}
