package okx

import (
	"context"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/jinzhu/copier"
)

func (m *CexOkex) GetTickerPrice(symbol string) (tickers []*types.TickerDetail, err error) {
	svc := m.client.NewGetTickerService()
	if symbol == "" {
		return nil, log.Errorf("symbol is empty")
	}
	res, err := svc.InstrumentId(symbol).Do(context.Background())
	if err != nil {
		return nil, log.Errorf("get ticker error: %s", err)
	}
	_ = copier.Copy(&tickers, &res.Data)
	return tickers, nil
}
