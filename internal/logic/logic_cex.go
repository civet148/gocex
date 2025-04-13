package logic

import (
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/tbtc-bot/go-okex"
)

type CexLogic struct {
	cfg    *config.Config
	client *okex.Client
}

func NewCexLogic(c *config.Config) *CexLogic {

	return &CexLogic{
		cfg:    c,
		client: okex.NewClient(c.ApiKey, c.ApiSecret, c.ApiPassphrase),
	}
}

func (m *CexLogic) Run() error {
	log.Infof("running...")
	bs, err := m.GetBalance(types.USDT)
	if err != nil {
		return log.Errorf(err)
	}
	log.Json("balances", bs)
	ts, err := m.GetSymbolTickerPrice(types.BTCUSDT)
	if err != nil {
		return log.Errorf(err)
	}
	log.Json("tickers", ts)
	return nil
}
