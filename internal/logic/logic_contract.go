package logic

import (
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"time"
)

type ContractLogic struct {
	cex    api.CexApi
	cfg    *config.Config
	ticker *TickerLogic
}

func NewContractLogic(cfg *config.Config, cex api.CexApi) *ContractLogic {
	return &ContractLogic{
		cfg:    cfg,
		cex:    cex,
		ticker: NewTickerLogic(cex),
	}
}

func (l *ContractLogic) Exec() error {
	c := l.cfg.Contract
	l.ticker.Start(c.Symbol, c.TickerDur)
	for {

		time.Sleep(l.cfg.Contract.OrderDur)
	}
	return nil
}
