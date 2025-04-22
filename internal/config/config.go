package config

import "time"

type Contract struct {
}

type Config struct {
	ApiKey            string        `yaml:"ApiKey" env:"API_KEY"`               //交易所API Key
	ApiSecret         string        `yaml:"ApiSecret" env:"API_SECRET"`         //交易所API密钥
	ApiPassphrase     string        `yaml:"ApiPassphrase" env:"API_PASSPHRASE"` //交易所API密码
	CexName           string        `yaml:"CexName"`                            //交易所名称(okex/binance...)
	Symbol            string        `yaml:"Symbol"`                             //下单交易对
	Leverage          int32         `yaml:"Leverage"`                           //杠杆倍数
	TickerDur         time.Duration `yaml:"TickerDur"`                          //市场价更新时间
	CheckDur          time.Duration `yaml:"CheckDur"`                           //市价检查间隔
	OrderDur          time.Duration `yaml:"OrderDur"`                           //下单时间间隔
	FastRise          float64       `yaml:"FastRise"`                           //暴涨或暴跌
	RiseThreshold     float64       `yaml:"RiseThreshold"`                      //涨幅触发阈值
	StopLossPercent   float64       `yaml:"StopLossPercent"`                    //止损百分比
	TakeProfitPercent float64       `yaml:"TakeProfitPercent"`                  //止盈百分比
	PullBackRate      float64       `yaml:"PullBackRate"`                       //价格回调比例
	TradeAmountRate   float64       `yaml:"TradeAmountRate"`                    //交易金额比例（余额百分比）
}
