package logic

import (
	"context"
	"github.com/civet148/log"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexLogic) GetSymbolTickerPrice(symbol string) ([]*okex.TickerDetail, error) {
	svc := m.client.NewGetTickerService()
	if symbol == "" {
		return nil, log.Errorf("symbol is empty")
	}
	ticker, err := svc.InstrumentId(symbol).Do(context.Background())
	if err != nil {
		return nil, log.Errorf("get ticker error: %s", err)
	}
	return ticker.Data, nil
}
