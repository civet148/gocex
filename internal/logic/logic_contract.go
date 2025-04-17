package logic

import (
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/utils"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"time"
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
	c := l.cfg.Contract
	var ticker = NewTickerLogic(l.cex) //市价行情
	ticker.Start(c.Symbol, c.TickerDur)

	for {
		act := l.getAction(ticker)
		switch act {
		case ActionWait:
		case ActionPos:
		case ActionClose:

		}
		time.Sleep(l.cfg.Strategy.CheckDur)
	}
	return nil
}

func (l *ContractLogic) getAction(ticker *TickerLogic) Action {
	contract := l.cfg.Contract
	now64 := utils.NowUnix()
	curPrice := ticker.GetCurrentPrice()
	comparePrice := l.comparePrice
	compareTime := l.compareTime

	defer func() {
		l.compareTime = now64
		l.comparePrice = curPrice
	}()

	if compareTime == 0 || comparePrice.IsZero() {
		return ActionWait
	}
	diff := curPrice.Sub(comparePrice)
	rise := diff.Div(comparePrice)
	log.Infof("[%v] last: %v current: %v rise: %v％", contract.Symbol, comparePrice, curPrice, rise.Mul(100).Round(2))
	return ActionWait
}
