package logic

import (
	"context"
	"github.com/civet148/log"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexLogic) GetSymbolTickerPrice(symbol string) ([]*okex.TickerDetail, error) {
	ticker, err := m.client.NewGetTickerService().InstrumentId(symbol).Do(context.Background())
	if err != nil {
		return nil, log.Errorf("get ticker error: %s", err)
	}
	log.Json("ticker", ticker)
	return ticker.Data, nil
}
