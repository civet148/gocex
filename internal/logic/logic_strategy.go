package logic

import (
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/utils"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"time"
)

type ContractStrategy struct {
	cfg           *config.Config
	ticker        Ticker
	basePrice     sqlca.Decimal
	position      bool
	entryPrice    sqlca.Decimal
	highestPrice  sqlca.Decimal
	leverage      int32         // 杠杆倍数(5-10倍)
	riseThreshold sqlca.Decimal // 上涨触发阈值
	stopLossPct   sqlca.Decimal // 止损百分比
	takeProfitPct sqlca.Decimal // 止盈百分比
}

func NewContractStrategy(cfg *config.Config, ticker Ticker) *ContractStrategy {
	return &ContractStrategy{
		cfg:           cfg,
		ticker:        ticker,
		leverage:      cfg.Leverage,
		riseThreshold: sqlca.NewDecimal(cfg.RiseThreshold),     // 涨幅触发
		stopLossPct:   sqlca.NewDecimal(cfg.StopLossPercent),   // 止损百分比
		takeProfitPct: sqlca.NewDecimal(cfg.TakeProfitPercent), // 止盈百分比
	}
}

func (cs *ContractStrategy) Start() (err error) {
	// 初始化基准价格
	cs.basePrice = cs.ticker.GetCurrentPrice()
	ticker := time.NewTicker(cs.cfg.CheckDur) // n分钟检查一次价格涨跌幅
	defer ticker.Stop()

	for range ticker.C {
		cs.monitorPrice()
	}
	return nil
}

func (cs *ContractStrategy) monitorPrice() {
	currentPrice := cs.ticker.GetCurrentPrice()
	if currentPrice.IsZero() {
		log.Errorf("current price is 0")
		return
	}
	if !cs.position {
		cs.checkEntryCondition(currentPrice)
	} else {
		cs.checkExitCondition(currentPrice)
	}
}

func (cs *ContractStrategy) checkEntryCondition(currentPrice sqlca.Decimal) {
	if cs.basePrice.IsZero() {
		cs.basePrice = currentPrice
	}
	// 计算从基准价的涨幅
	risePct := (currentPrice.Float64() - cs.basePrice.Float64()) / cs.basePrice.Float64()
	if risePct == 0 {
		return
	}
	rise := sqlca.NewDecimal(risePct)
	log.Printf("[%v] base: %v current: %v rise: [%v％]",
		cs.cfg.Symbol, utils.FormatDecimal(cs.basePrice, 9),
		utils.FormatDecimal(currentPrice, 9), cs.getPercentRise(rise))

	// 满足上涨阈值且未持仓
	if risePct >= cs.riseThreshold.Float64() {
		cs.openPosition(currentPrice)
	}

	// 更新基准价(动态调整)
	if currentPrice.Float64() < cs.basePrice.Float64() {
		cs.basePrice = currentPrice
	}
}

func (cs *ContractStrategy) checkExitCondition(currentPrice sqlca.Decimal) {
	if cs.entryPrice.IsZero() {
		log.Errorf("entry price is 0")
		return
	}
	// 更新最高价
	if currentPrice.Float64() > cs.highestPrice.Float64() {
		cs.highestPrice = currentPrice
	}

	// 计算盈亏比例
	profitPct := (currentPrice.Float64() - cs.entryPrice.Float64()) / cs.entryPrice.Float64() * float64(cs.leverage)
	log.Printf("[%v] entry: %v current: %v profit: [%v％]",
		cs.cfg.Symbol, utils.FormatDecimal(cs.entryPrice, 9),
		utils.FormatDecimal(currentPrice, 9), cs.getPercentRise(sqlca.NewDecimal(profitPct)))
	// 止盈或止损检查
	if profitPct >= cs.takeProfitPct.Float64() || profitPct <= -cs.stopLossPct.Float64() {
		cs.closePosition(currentPrice)
	}
}

func (cs *ContractStrategy) openPosition(price sqlca.Decimal) {
	log.Printf("开仓信号 价格: %v 杠杆: %d倍", utils.FormatDecimal(price, 9), cs.leverage)
	cs.position = true
	cs.entryPrice = price
	cs.highestPrice = price

	// TODO: 实现实际合约开仓
	// 这里应包含杠杆设置和风险控制
}

func (cs *ContractStrategy) closePosition(price sqlca.Decimal) {
	profit := (price.Float64() - cs.entryPrice.Float64()) / cs.entryPrice.Float64() * float64(cs.leverage)
	log.Printf("平仓信号 价格: %v 收益率: %.2f%%", utils.FormatDecimal(price, 9), profit*100)
	cs.position = false

	// TODO: 实现实际合约平仓
	// 重置状态
	cs.basePrice = price // 平仓后重置基准价
}

func (cs *ContractStrategy) getPercentRise(rise sqlca.Decimal) string {
	strPercent := rise.Mul(100).Round(2).String()
	if rise.LessThan(0) {
		return utils.Red(strPercent)
	}
	if rise.Float64() < cs.cfg.FastRise {
		return utils.White(strPercent)
	}
	return utils.Green(strPercent)
}
