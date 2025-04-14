package logic

import (
	"github.com/civet148/gocex/internal/api"
	_ "github.com/civet148/gocex/internal/cexs/okx"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
)

type CexLogic struct {
	cfg *config.Config
	cex api.CexApi
}

func NewCexLogic(c *config.Config) *CexLogic {

	return &CexLogic{
		cfg: c,
		cex: api.NewCex(types.CexTypeOkex, c),
	}
}

func (m *CexLogic) Run() error {
	log.Infof("running...")
	bs, err := m.cex.GetBalance(types.USDT)
	if err != nil {
		return log.Errorf(err)
	}
	log.Json("balances", bs)
	ts, err := m.cex.GetTickerPrice(types.BTCUSDT)
	if err != nil {
		return log.Errorf(err)
	}
	log.Json("tickers", ts)
	odrs, err := m.cex.GetOrder()
	if err != nil {
		return log.Errorf(err)
	}
	log.Json("orders", odrs)
	return nil
}
