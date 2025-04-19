package logic

import (
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/sqlca/v2"
	"log"
	"time"
)

type ContractStrategy struct {
	ticker        Ticker
	basePrice     sqlca.Decimal
	position      bool
	entryPrice    sqlca.Decimal
	highestPrice  sqlca.Decimal
	leverage      int           // 杠杆倍数(5-10倍)
	riseThreshold sqlca.Decimal // 上涨触发阈值
	stopLossPct   sqlca.Decimal // 止损百分比
	takeProfitPct sqlca.Decimal // 止盈百分比
}

func NewContractStrategy(cfg *config.Config, ticker Ticker, leverage int) *ContractStrategy {
	return &ContractStrategy{
		ticker:        ticker,
		leverage:      leverage,
		riseThreshold: sqlca.NewDecimal(cfg.RiseThreshold),     // 涨幅触发
		stopLossPct:   sqlca.NewDecimal(cfg.StopLossPercent),   // 止损百分比
		takeProfitPct: sqlca.NewDecimal(cfg.TakeProfitPercent), // 止盈百分比
	}
}

func (cs *ContractStrategy) Start() {
	// 初始化基准价格
	cs.basePrice = cs.ticker.GetCurrentPrice()

	ticker := time.NewTicker(1 * time.Minute) // 每分钟检查一次
	defer ticker.Stop()

	for range ticker.C {
		currentPrice := cs.ticker.GetCurrentPrice()
		cs.monitorPrice(currentPrice)
	}
}

func (cs *ContractStrategy) monitorPrice(currentPrice sqlca.Decimal) {
	if !cs.position {
		cs.checkEntryCondition(currentPrice)
	} else {
		cs.checkExitCondition(currentPrice)
	}
}

func (cs *ContractStrategy) checkEntryCondition(currentPrice sqlca.Decimal) {
	// 计算从基准价的涨幅
	risePct := (currentPrice.Float64() - cs.basePrice.Float64()) / cs.basePrice.Float64()

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
	// 更新最高价
	if currentPrice.Float64() > cs.highestPrice.Float64() {
		cs.highestPrice = currentPrice
	}

	// 计算盈亏比例
	profitPct := (currentPrice.Float64() - cs.entryPrice.Float64()) / cs.entryPrice.Float64() * float64(cs.leverage)

	// 止盈或止损检查
	if profitPct >= cs.takeProfitPct.Float64() || profitPct <= -cs.stopLossPct.Float64() {
		cs.closePosition(currentPrice)
	}
}

func (cs *ContractStrategy) openPosition(price sqlca.Decimal) {
	log.Printf("开仓信号 价格: %.4f 杠杆: %d倍", price, cs.leverage)
	cs.position = true
	cs.entryPrice = price
	cs.highestPrice = price

	// TODO: 实现实际合约开仓
	// 这里应包含杠杆设置和风险控制
}

func (cs *ContractStrategy) closePosition(price sqlca.Decimal) {
	profit := (price.Float64() - cs.entryPrice.Float64()) / cs.entryPrice.Float64() * float64(cs.leverage)
	log.Printf("平仓信号 价格: %.4f 收益率: %.2f%%", price, profit*100)
	cs.position = false

	// TODO: 实现实际合约平仓
	// 重置状态
	cs.basePrice = price // 平仓后重置基准价
}
