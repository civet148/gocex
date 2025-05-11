package logic

import (
	"context"
	"time"

	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/options"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/gocex/internal/utils"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type ContractLogic struct {
	*config.Config
	cex          api.CexApi             // 交易所对象
	ticker       Ticker                 // 市价ticker对象
	instId       string                 // 交易代币对
	basePrice    sqlca.Decimal          // 基础价格
	lastPrice    sqlca.Decimal          // 上次检查的价格
	position     bool                   // 是否已持仓
	entryPrice   sqlca.Decimal          // 开仓价
	highestPrice sqlca.Decimal          // 最高价
	riseCount    int32                  // 价格持续上涨次数
	pullCount    int32                  // 价格持续暴跌次数
	cliOrderId   string                 // 客户订单ID
	posOrder     *types.OrderListDetail // 持仓订单
}

func NewContractLogic(cfg *config.Config, cex api.CexApi) *ContractLogic {
	var ticker = NewTickerLogic(cex) //市价行情
	return &ContractLogic{
		Config: cfg,
		cex:    cex,
		ticker: ticker,
		instId: cfg.Symbol,
	}
}

func (l *ContractLogic) Exec() (err error) {
	err = l.initContractInstrument()
	if err != nil {
		return err
	}
	// 初始化基准价格
	l.basePrice, err = l.ticker.GetCurrentPrice(l.Symbol)
	if err != nil {
		return log.Errorf(err)
	}

	priceTicker := time.NewTicker(l.CheckDur) // n分钟检查一次价格涨跌幅
	defer priceTicker.Stop()

	// 添加止盈止损检查定时器
	winLossTicker := time.NewTicker(5 * time.Minute)
	defer winLossTicker.Stop()

	// 添加回调检查定时器
	var pullbackTicker *time.Ticker
	pullbackTicker = time.NewTicker(15 * time.Minute)
	defer pullbackTicker.Stop()

	for {
		select {
		case <-priceTicker.C:
			l.monitorPrice() //价格监视
		case <-winLossTicker.C:
			l.checkStopWinOrLoss() //止盈止损
		case <-pullbackTicker.C:
			l.checkPullback() //价格回调
		}
	}
	return nil
}

// 新增止损检查方法
func (l *ContractLogic) checkStopWinOrLoss() {
	if !l.position {
		return
	}
	if l.posOrder == nil {
		return
	}
	currentPrice, err := l.ticker.GetCurrentPrice(l.Symbol)
	if err != nil {
		return
	}

	// 当前盈亏比
	pct := l.posOrder.UplRatio
	if pct.LessThan(0) { //亏损
		if pct.Abs().GreaterThan(l.StopLossPct) {
			log.Warnf("[%v] 触发止损 开仓价: %v 当前价: %v 亏损比例: %v％  亏损金额：%vUSD",
				l.Symbol,
				utils.FormatDecimal(l.entryPrice, 9),
				utils.FormatDecimal(currentPrice, 9),
				l.formatRisePercent(pct),
				l.posOrder.Upl.Round(2),
			)
			_ = l.closePosition(currentPrice)
		}
	} else { //盈利
		if pct.GreaterThan(l.StopWinPct) {
			log.Warnf("[%v] 触发止赢 开仓价: %v 当前价: %v 盈利比例: %v％ 盈利金额：%vUSD",
				l.Symbol,
				utils.FormatDecimal(l.entryPrice, 9),
				utils.FormatDecimal(currentPrice, 9),
				l.formatRisePercent(pct),
				l.posOrder.Upl.Round(2),
			)
		}
	}
}

// 新增回调检查方法
func (l *ContractLogic) checkPullback() {
	if !l.position {
		return
	}

	currentPrice, err := l.ticker.GetCurrentPrice(l.Symbol)
	if err != nil {
		return
	}

	// 计算从最高价的回调幅度
	risePct := currentPrice.Sub(l.highestPrice).Div(l.highestPrice)

	if risePct.LessThan(0) && risePct.Abs().GreaterThan(l.PullBackRate) {
		log.Warnf("[%v] 触发回调平仓 最高价: %v 当前价: %v 回调幅度: %v％ (配置回调幅度：%v％)",
			l.Symbol,
			utils.FormatDecimal(l.highestPrice, 9),
			utils.FormatDecimal(currentPrice, 9),
			l.formatRisePercent(risePct), l.PullBackRate*100,
		)
		_ = l.closePosition(currentPrice)
	}
}

// 初始化合约参数
func (l *ContractLogic) initContractInstrument() (err error) {
	var ctx = context.Background()
	//检查杠杆倍数
	var instruments []*types.InstrumentDetail
	instruments, err = l.cex.GetInstrument(ctx, l.Symbol, types.InstType_SWAP)
	if err != nil {
		return log.Errorf("查询基础信息失败: %s", err.Error())
	}
	for _, instr := range instruments {
		//log.Infof("%v最大合约杠杆倍数为%v", l.Symbol, instr.Lever.String())
		if instr.Lever.LessThan(l.Leverage) {
			return log.Errorf("杠杆倍数%v不能超过最大杠杆倍数%v", l.Leverage, instr.Lever.String())
		}
		//设置杠杆倍数(逐仓)
		lever := sqlca.NewDecimal(l.Leverage)
		var opts []options.TradeOption
		opts = append(opts, options.WithLeverage(lever.String()))
		opts = append(opts, options.WithSwap())
		_, err = l.cex.SetLeverage(ctx, l.Symbol, types.MarginModeIsolated, opts...)
		if err != nil {
			return log.Errorf("设置杠杆失败: %s", err.Error())
		}
	}
	return nil
}

func (l *ContractLogic) loadContractPosition() (err error) {
	//加载已持仓合约
	var positions []*types.OrderListDetail
	positions, err = l.cex.GetPosition(context.Background(), l.Symbol)
	if err != nil {
		return log.Errorf("加载合约失败: %s", err.Error())
	}
	if len(positions) == 0 && l.position {
		l.resetPosition(l.lastPrice) //已手动结束合约
		return nil
	}
	for _, pos := range positions {
		l.position = true
		l.entryPrice = pos.AvgPx
		l.posOrder = pos
		if l.entryPrice.LessThan(pos.Last) {
			l.highestPrice = pos.Last
		} else {
			l.highestPrice = l.entryPrice
		}
		log.Infof("[%v合约] 倍数: %v 开仓均价: %v 标记价格: %v 收益：%v %vUSD",
			pos.InstId, pos.Lever,
			utils.FormatDecimal(pos.AvgPx, 9),
			utils.FormatDecimal(pos.Last, 9),
			l.formatRisePercent(pos.UplRatio),
			pos.Upl.Round(2))
	}
	return nil
}

func (l *ContractLogic) getActivePosition() (pos *types.OrderListDetail, err error) {
	var positions []*types.OrderListDetail
	positions, err = l.cex.GetPosition(context.Background(), l.Symbol)
	if err != nil {
		return nil, log.Errorf("查询合约持仓信息失败: %s", err.Error())
	}
	for _, p := range positions {
		pos = p
		break
	}
	return pos, nil
}

func (l *ContractLogic) monitorPrice() {
	currentPrice, err := l.ticker.GetCurrentPrice(l.Symbol)
	if err != nil {
		return
	}
	//加载已持仓合约
	err = l.loadContractPosition()
	if err != nil {
		return
	}

	if currentPrice.IsZero() {
		log.Errorf("current price is 0")
		return
	}
	if !l.position {
		l.checkEntryCondition(currentPrice)
	} else {
		// 更新最高价
		l.checkExitCondition(currentPrice)
	}
}

func (l *ContractLogic) checkEntryCondition(currentPrice sqlca.Decimal) {
	if l.basePrice.IsZero() {
		l.basePrice = currentPrice
	}
	if l.lastPrice.IsZero() {
		l.lastPrice = currentPrice
	}
	// 计算从基准价的涨幅
	riseBase := currentPrice.Sub(l.basePrice).Div(l.basePrice)

	// 计算上次检查价格的涨幅
	riseLast := currentPrice.Sub(l.lastPrice).Div(l.lastPrice)

	if riseLast.GreaterThanOrEqual(l.FastRise) {
		l.riseCount++
	} else {
		l.riseCount = 0
	}
	log.Infof("[%v] 基础价: %v 市场价: %v 总涨幅[%v％] 单次涨幅 [%v％] 持续次数 [%v]",
		l.Symbol, utils.FormatDecimal(l.basePrice, 9),
		utils.FormatDecimal(currentPrice, 9),
		l.formatRisePercent(riseBase),
		l.formatRisePercent(riseLast),
		l.riseCount)

	// 满足上涨阈值且未持仓
	if /*riseBase.GreaterThanOrEqual(l.RiseThreshold) ||*/ l.riseCount >= l.Continuous {
		err := l.openPosition(currentPrice)
		if err != nil {
			return
		}
		l.riseCount = 0
	}

	// 更新基准价(动态调整)
	if currentPrice.Float64() < l.basePrice.Float64() {
		l.basePrice = currentPrice
	}
	// 更新上次价格
	l.lastPrice = currentPrice
}

func (l *ContractLogic) checkExitCondition(currentPrice sqlca.Decimal) {
	if l.basePrice.IsZero() {
		l.basePrice = currentPrice
	}
	if l.lastPrice.IsZero() {
		l.lastPrice = currentPrice
	}
	if currentPrice.Float64() > l.highestPrice.Float64() {
		l.highestPrice = currentPrice
	}
	//// 计算从基准价的涨幅
	//riseBase := currentPrice.Sub(l.basePrice).Div(l.basePrice)
	//
	// 计算上次检查价格的涨幅
	riseLast := currentPrice.Sub(l.lastPrice).Div(l.lastPrice)

	if riseLast.LessThan(0) && riseLast.Abs().GreaterThanOrEqual(l.FastRise) { //计算暴跌持续次数
		l.pullCount++
	} else {
		l.pullCount = 0
	}

	// 满足持续暴跌次数强制平仓
	if l.pullCount >= l.Continuous-1 {
		log.Warnf("[%v] 基础价: %v 市场价: %v 暴跌持续次数 [%v] 强制平仓",
			l.Symbol,
			utils.FormatDecimal(l.basePrice, 9),
			utils.FormatDecimal(currentPrice, 9),
			l.pullCount)
		err := l.closePosition(currentPrice)
		if err != nil {
			return
		}
		l.pullCount = 0
	}

	// 更新基准价(动态调整)
	if currentPrice.Float64() < l.basePrice.Float64() {
		l.basePrice = currentPrice
	}
	// 更新上次价格
	l.lastPrice = currentPrice
}

func (l *ContractLogic) openPosition(price sqlca.Decimal) error {
	l.position = true
	l.entryPrice = price
	l.highestPrice = price

	sz, err := l.calcContractSz(price)
	if err != nil {
		return log.Errorf("计算购买合约张数失败：%s", err.Error())
	}
	// 实际合约开仓(这里应包含杠杆设置和风险控制)
	if !l.Simulate {
		l.cliOrderId, err = l.createPosition(sz, types.SideTypeBuy)
		if err != nil {
			return log.Errorf("合约建仓失败：%s", err.Error())
		}
		log.Warnf("[%v] 开仓信号 价格: %v 杠杆: %v倍", l.Symbol, utils.FormatDecimal(price, 9), l.Leverage)
	} else {
		log.Infof("[%v] 开仓信号 价格: %v 杠杆: %v倍 (模拟交易模式)", l.Symbol, utils.FormatDecimal(price, 9), l.Leverage)
	}
	return nil
}

func (l *ContractLogic) closePosition(price sqlca.Decimal) (err error) {
	if l.Simulate {
		profit := (price.Float64() - l.entryPrice.Float64()) / l.entryPrice.Float64() * float64(l.Leverage)
		log.Infof("[%v] 平仓信号 价格: %v 收益率: %.2f％ (模拟交易模式)", l.Symbol, utils.FormatDecimal(price, 9), profit*100)
	} else {
		var pos *types.OrderListDetail
		pos, err = l.getActivePosition()
		if err != nil {
			return err
		}
		if pos == nil {
			l.resetPosition(price)
			return log.Errorf("未查询到持仓信息(可能已手动平仓)，重置持仓状态")
		}
		log.Warnf("[%v] 平仓信号 价格: %v 收益率: %v％ 总收益: %vUSD",
			l.Symbol, utils.FormatDecimal(pos.Last, 9), pos.UplRatio.Round(2), pos.Upl.Round(2))
	}

	// 实际合约平仓
	if !l.Simulate {
		err = l.closePositionByInstId()
		if err != nil {
			return err
		}
	}
	l.resetPosition(price)
	return nil
}

func (l *ContractLogic) resetPosition(price sqlca.Decimal) {
	l.position = false  //重置持仓状态
	l.basePrice = price //平仓后重置基准价
	l.cliOrderId = ""   //重置客户订单ID
}

// 通过instId关闭仓位
func (l *ContractLogic) closePositionByInstId() (err error) {
	var opts []options.TradeOption
	opts = append(opts, options.WithMarginMode(types.MarginModeIsolated))
	if l.cliOrderId != "" {
		opts = append(opts, options.WithCliOrdId(l.cliOrderId))
	}
	var orders []*types.ClosePositionDetail
	orders, err = l.cex.ClosePosition(context.Background(), l.Symbol, opts...)
	if err != nil {
		return log.Errorf("[%s] 平仓失败：%s", l.Symbol, err.Error())
	}
	_ = orders
	return nil
}

// 开始建仓(buy=多 sell=空)
func (l *ContractLogic) createPosition(sz sqlca.Decimal, sideType types.SideType) (cliOrdId string, err error) {
	var opts []options.TradeOption
	cliOrdId = utils.GenClientOrderId()
	opts = append(opts, options.WithOrderType(types.OrderTypeMarket))
	opts = append(opts, options.WithTradeMode(types.TradeModeIsolated))
	opts = append(opts, options.WithCliOrdId(cliOrdId))

	var orders []*types.OrderDetail
	orders, err = l.cex.OpenPosition(context.Background(), l.Symbol, sideType, sz, opts...)
	if err != nil {
		return cliOrdId, log.Errorf(err.Error())
	}
	_ = orders
	log.Infof("[%s] 合约数量: %v 建仓成功，客户订单ID: %s", l.Symbol, sz, cliOrdId)
	return cliOrdId, nil
}

// 计算合约张数
func (l *ContractLogic) calcContractSz(price sqlca.Decimal) (sz sqlca.Decimal, err error) {
	var usdt sqlca.Decimal
	usdt, err = l.ticker.GetAvailableUSDT()
	if err != nil {
		return sz, log.Errorf("查询可用USDT余额失败 error: %s", err.Error())
	}
	usdt = usdt.Mul(l.TradeAmountRate) //实际交易的USDT数量

	var insts []*types.InstrumentDetail
	insts, err = l.cex.GetInstrument(context.Background(), l.Symbol, types.InstType_SWAP)
	if err != nil {
		return sz, log.Errorf("查询合约信息失败 error: %s", err.Error())
	}
	for _, inst := range insts {
		if inst.Uly == l.Symbol {
			ctValue := inst.CtVal.Mul(price).Round(2)       //单张合约USD价值
			sz = usdt.Div(ctValue).Mul(l.Leverage).Round(1) //根据杠杆倍数计算实际张数
			log.Infof("[%s] 市价: %v 合约单张价值：%vUSD 实际购买张数：%v 总费用：%vUSD",
				l.Symbol, utils.FormatDecimal(price, 9), ctValue, sz, sz.Mul(ctValue))
			break
		}
	}
	return sz, nil
}

func (l *ContractLogic) formatRisePercent(rise sqlca.Decimal) string {
	strPercent := rise.Mul(100).Round(2).String()
	if rise.LessThan(0) {
		return utils.Red(strPercent)
	}
	if rise.Float64() < l.FastRise {
		return utils.White(strPercent)
	}
	return utils.Green(strPercent)
}
