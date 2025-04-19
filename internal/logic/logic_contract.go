package logic

import (
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/utils"
	"github.com/civet148/log"
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
	//var ticker = NewTickerLogic(l.cex, c.Symbol, c.TickerDur) //市价行情
	//strategy := NewContractStrategy(ticker, 8) // 使用8倍杠杆
	//for {
	//	act := l.getAction(ticker)
	//	switch act {
	//	case ActionWait:
	//	case ActionPos:
	//	case ActionClose:
	//	}
	//	time.Sleep(l.cfg.CheckDur)
	//}
	return nil
}

func (l *ContractLogic) getAction(ticker *TickerLogic) Action {
	contract := l.cfg
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
	log.Infof("[%v] last: %v current: %v rise: [%v％]", contract.Symbol, utils.FormatDecimal(comparePrice, 9), utils.FormatDecimal(curPrice, 9), l.GetPercentRise(rise))
	return ActionWait
}

func (l *ContractLogic) GetPercentRise(rise sqlca.Decimal) string {
	c := l.cfg
	strPercent := rise.Mul(100).Round(2).String()
	if rise.LessThan(0) {
		return utils.Red(strPercent)
	}
	if rise.Float64() < c.FastRise {
		return utils.White(strPercent)
	}
	return utils.Green(strPercent)
}
