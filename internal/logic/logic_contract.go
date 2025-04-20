package logic

import (
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/sqlca/v2"
)

type Action int32

const (
	ActionWait  Action = 0 //等待
	ActionPos   Action = 1 //持仓
	ActionClose Action = 2 //平仓
)

type ContractLogic struct {
	cex          api.CexApi     //交易所对象
	cfg          *config.Config //配置
	comparePrice sqlca.Decimal  //对比价格
	compareTime  int64          //对比时间
}

func NewContractLogic(cfg *config.Config, cex api.CexApi) *ContractLogic {
	return &ContractLogic{
		cfg: cfg,
		cex: cex,
	}
}

func (l *ContractLogic) Exec() error {
	cfg := l.cfg
	var ticker = NewTickerLogic(l.cex, cfg.Symbol, cfg.TickerDur) //市价行情
	strategy := NewContractStrategy(cfg, ticker)                  // 使用8倍杠杆
	return strategy.Start()
}
